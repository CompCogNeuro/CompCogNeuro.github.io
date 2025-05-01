// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/axon/v2/axon"
	"github.com/emer/axon/v2/chans"
	"github.com/emer/emergent/v2/egui"
)

// Urakubo has all of the state for the Urakubo model.
type Urakubo struct {
	// Net is the axon network for basic activity dynamics
	Net *axon.Network `view:"no-inline"`

	// the spine state with Urakubo intracellular model
	Spine Spine

	// extra neuron state for additional channels: Vgcc, AK
	NeuronEx NeuronEx `view:"no-inline"`

	// what stimulation to drive with
	Stim Stims

	// whether to initialize values to known baseline values at start
	InitBaseline bool

	// inter-stimulus-interval in seconds -- between reps
	ISISec float64

	// number of repetitions -- takes 100 to produce classic STDP
	NReps int

	// number of seconds to run after the manipulation -- results are strongest after 100,
	// decaying somewhat after that point -- 20 shows similar qualitative results but weaker,
	// 50 is pretty close to 100 -- less than 20 not recommended.
	FinalSecs float64 `def:"20,50,100"`

	// duration for activity window
	DurMsec int

	// sending firing frequency (used as minus phase for ThetaErr)
	SendHz float32

	// receiving firing frequency (used as plus phase for ThetaErr)
	RecvHz float32

	// stimulating current injection
	GeStim float32

	// in msec, difference of Tpost - Tpre == pos = LTP, neg = LTD STDP
	DeltaT int

	// range for sweep of DeltaT -- actual range is - to +
	DeltaTRange int

	// increment for sweep of DeltaT
	DeltaTInc int

	// use Ge current clamping instead of distrete pulsing for firing rate-based manips, e.g., ThetaErr
	RGClamp bool

	// use dendritic Vm signal for driving spine channels
	VmDend bool

	// use the Axon NMDA channel instead of the allosteric Urakubo one
	NMDAAxon bool

	// strength of NMDA current -- 0.15 default for posterior cortex
	NMDAGbar float32 `def:"0,0.15"`

	// strength of GABAB current -- 0.2 default for posterior cortex
	GABABGbar float32 `def:"0,0.2"`

	// strength of Vgcc current -- 1.46 for distal per Migliore, but 0.12 reproduces net Ca current..
	VgccGbar float32 `def:"0,0.12"`

	// target calcium level for CaTarg stim
	CaTarg CaState

	// initial weight value: Trp_AMPA value at baseline
	InitWt float64 `editable:"-"`

	// current cycle of updating
	Msec int `inactive:"+"`

	Stats     *tensorfs.Node   `display:"-"`
	GUI       *egui.GUI        `display:"-"`
	StimFuncs map[Stims]func() `display:"-"`
}

func (uk *Urakubo) Config() {
	uk.Spine.Defaults()
	uk.Spine.Init()
	uk.InitWt = uk.Spine.States.AMPAR.Trp.Tot
	uk.Net = axon.NewNetwork("Urakubo")
	uk.Stim = ThetaErrComp
	uk.ISISec = 0.8
	uk.NReps = 10
	uk.FinalSecs = 0
	uk.DurMsec = 200
	uk.SendHz = 50
	uk.RecvHz = 25
	uk.DeltaT = 16
	uk.DeltaTRange = 50
	uk.DeltaTInc = 5
	uk.RGClamp = true
	uk.Defaults()
	uk.ConfigNet(uk.Net)
	uk.ConfigStats()
	uk.ConfigStimFuncs()
}

// Defaults sets default params
func (uk *Urakubo) Defaults() {
	uk.Opts.Defaults()
	uk.Spine.Defaults()
	uk.GeStim = 2
	uk.NMDAGbar = 0.15 // 0.1 to 0.15 matches pre-spike increase in vm -- note that nominal val is .5
	uk.GABABGbar = 0.0 // 0.2
	uk.VgccGbar = 0.12 // note: was .12 to match existing traces, but nominal val in paper is 1.46
	uk.CaTarg.Cyt = 10
	uk.CaTarg.PSD = 10
}

func (uk *Urakubo) ConfigNet(net *axon.Network) {
	net.SetMaxData(1)
	net.Context().ThetaCycles = int32(ss.Config.Run.Cycles)
	net.SetRandSeed(ss.RandSeeds[0]) // init new separate random seed, using run = 0

	ly := net.AddLayer2D("Neuron", axon.SuperLayer, 1, 1)

	net.Build()
	net.Defaults()
	uk.ApplyParams()
	net.InitWeights()
	uk.Neuron = &ly.Neurons[0]
}

func (ss *Sim) ApplyParams() {
	ss.Params.Script = ss.Config.Params.Script
	ss.Params.ApplyAll(ss.Net)

	lys := ss.Net.LayersByType(axon.SuperLayer)
	for _, lnm := range lys {
		ly := ss.Net.LayerByName(lnm)
		lyp := ly.Params
		lyp.Acts.NMDA.Ge = ss.NmdaGe
		lyp.Acts.GabaB.Gk = ss.GababGk
		lyp.Inhib.Layer.Gi = ss.InhibGi
	}
}

func (uk *Urakubo) ConfigStimFuncs() {
	uk.StimFuncs = map[Stims]func(){
		Baseline:         uk.BaselineFun,
		CaTarg:           uk.CaTargFun,
		ClampCa1:         uk.ClampCa1Fun,
		STDP:             uk.STDPFun,
		STDPSweep:        uk.STDPSweepFun,
		STDPPacketSweep:  uk.STDPPacketSweepFun,
		Poisson:          uk.PoissonFun,
		SPoissonRGClamp:  uk.SPoissonRGClampFun,
		PoissonHzSweep:   uk.PoissonHzSweepFun,
		PoissonDurSweep:  uk.PoissonDurSweepFun,
		OpPhaseDurSweep:  uk.OpPhaseDurSweepFun,
		ThetaErr:         uk.ThetaErrFun,
		ThetaErrComp:     uk.ThetaErrCompFun,
		ThetaErrSweep:    uk.ThetaErrSweepFun,
		ThetaErrAllSweep: uk.ThetaErrAllSweepFun,
	}
}

// Init restarts the run, and initializes everything, including network weights
// and resets the epoch log table
func (uk *Urakubo) Init() {
	uk.Spine.Init()
	uk.InitWt = uk.Spine.States.AMPAR.Trp.Tot
	uk.NeuronEx.Init()
	uk.Msec = 0
	// uk.SetParams("", false) // all sheets
	ly := uk.Net.LayerByName("Neuron").(axon.AxonLayer).AsAxon()
	if uk.NMDAAxon {
		ly.Params.Act.NMDA.Gbar = uk.NMDAGbar
	} else {
		ly.Params.Act.NMDA.Gbar = 0
	}
	ly.Params.Act.GABAB.Gbar = uk.GABABGbar
	ly.Params.Act.VGCC.Gbar = uk.VgccGbar
}

// Counters returns a string of the current counter state
// use tabs to achieve a reasonable formatting overall
// and add a few tabs at the end to allow for expansion..
func (uk *Urakubo) Counters() string {
	return fmt.Sprintf("Msec:\t%d\t\t\t", uk.Msec)
}

func (uk *Urakubo) StopNow() bool {
	if uk.GUI == nil {
		return false
	}
	return uk.GUI.StopNow
}

func (uk *Urakubo) Stopped() {
	if uk.GUI == nil {
		return
	}
	uk.GUI.Stopped(Test, Cycle)
}

// RunStim runs current Stim selection
func (uk *Urakubo) RunStim() {
	fn, has := uk.StimFuncs[uk.Stim]
	if !has {
		fmt.Printf("Stim function: %s not found!\n", uk.Stim)
		return
	}
	go fn()
}

// NeuronUpdt updates the neuron and spine for given msec
func (uk *Urakubo) NeuronUpdt(msec int, ge, gi float32) {
	uk.Msec = msec
	ly := uk.Net.LayerByName("Neuron").(axon.AxonLayer).AsAxon()
	// nrn := uk.Neuron
	nex := &uk.NeuronEx

	vbio := chans.VToBio(nrn.Vm) // dend

	// note: Ge should only
	geExt := float32(0)
	nrn.GeRaw = ge
	nrn.GeSyn = ly.Params.Act.Dt.GeSynFmRaw(nrn.GeRaw, ly.Params.Act.Init.GeBase)
	nrn.Ge = nrn.GeSyn
	nrn.Gi = gi
	ly.Params.Act.NMDAFmRaw(nrn, geExt)
	ly.Params.Act.GvgccFmVm(nrn)
	ly.Params.Learn.LrnNMDAFmRaw(nrn, geExt)
	ly.Params.Act.GABAB.GABAB(nrn.GABAB, nrn.GABABx, nrn.Gi, &nrn.GABAB, &nrn.GABABx)
	nrn.GgabaB = ly.Params.Act.GABAB.GgabaB(nrn.GABAB, nrn.VmDend)
	nrn.Gak = ly.Params.Act.AK.Gak(nrn.VmDend)
	nrn.Gk += nrn.Gak
	nrn.Ge += nrn.Gvgcc + nrn.Gnmda
	if !uk.NMDAAxon {
		nrn.Ge += uk.NMDAGbar * float32(uk.Spine.States.NMDAR.G)
	}
	nrn.Gi += nrn.GgabaB

	// todo: Ca from NMDAAxon
	uk.Spine.Ca.SetInject(float64(nex.VgccJcaPSD), float64(nex.VgccJcaCyt))

	psd_pca := float32(1.7927e5 * 0.04) //  SVR_PSD
	cyt_pca := float32(1.0426e5 * 0.04) // SVR_CYT

	nex.VgccJcaPSD = -vbio * psd_pca * nrn.Gvgcc
	nex.VgccJcaCyt = -vbio * cyt_pca * nrn.Gvgcc

	uk.Spine.States.VmS = float64(vbio)

	ly.Params.Act.VmFmG(nrn)
	ly.Params.Act.SpikeFmVm(nrn)

	uk.Spine.StepTime(0.001)
}

//////// Stats

var StatsNames = []string{"Msec", "Msec10", "Msec100", "DWt"}

// ConfigStats
func (uk *Urakubo) ConfigStats(dir *tensorfs.Node) {
	uk.Stats = dir.Dir("Stats")
}

// StatsInit initializes all the stats by setting num rows to 0
func (uk *Urakubo) StatsInit() {
	// todo: iterations
	dir := uk.Stats
	idx := 0
	if uk.GUI.Tabs != nil {
		_, idx = uk.GUI.Tabs.AsLab().CurrentTab()
	}
	for _, sn := range StatsNames {
		uk.StatsInitDir(sn)
		if uk.GUI.Tabs != nil {
			sd := dir.Dir(sn)
			uk.GUI.Tabs.AsLab().PlotTensorFS(sd)
		}
	}
	if uk.GUI.Tabs != nil {
		uk.GUI.Tabs.AsLab().SelectTabIndex(idx)
	}
}

// StatsInitDir initializes given stats directory
func (uk *Urakubo) StatsInitDir(dir *tensorfs.Node, sn string) {
	sd := dir.Dir(sn)
	tsrs := sd.Values()
	for _, tsr := range tsrs {
		tsr.SetNumRows(0)
	}
}

// StatsTime adds data to given tensorfs directory
func (uk *Urakubo) StatsTime(dir *tensorfs.Node) {
	nex := &uk.NeuronEx
	ni := 0
	di := 0

	dir.Float64("Time").AppendRowFloat(float64(uk.Msec) * 0.001)
	dir.Float64("Ge").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Ge))))
	dir.Float64("Inet").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Inet))))
	dir.Float64("Vm").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Vm))))
	dir.Float64("Act").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Act))))
	dir.Float64("Spike").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Spike))))
	dir.Float64("Gk").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gk))))
	dir.Float64("ISI").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.ISI))))
	dir.Float64("AvgISI").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.ISIAvg))))
	dir.Float64("VmDend").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.VmDend))))
	dir.Float64("SnmdaO").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.SnmdaO))))
	dir.Float64("Gnmda").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gnmda))))
	dir.Float64("GABAB").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.GABAB))))
	dir.Float64("GgabaB").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.GgabaB))))
	dir.Float64("Gvgcc").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gvgcc))))
	dir.Float64("VgccM").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.VgccM))))
	dir.Float64("VgccH").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.VgccH))))
	dir.Float64("VgccJcaPSD").AppendRowFloat(float64(nex.VgccJcaPSD))
	dir.Float64("VgccJcaCyt").AppendRowFloat(float64(nex.VgccJcaCyt))
	dir.Float64("Gak").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gak))))
	dir.Float64("AKm").AppendRowFloat(float64(nex.AKm))
	dir.Float64("AKh").AppendRowFloat(float64(nex.AKh))

	uk.Spine.Stats(dir)
}

// StatsDefault does default logging for current Msec, for given iteration.
func (uk *Urakubo) StatsDefault(itr int) {
	// todo: deal with iteration as a subdir in stats
	dir := uk.Stats
	msec := uk.Msec
	uk.StatsTime(dir.Dir("Msec"))
	if uk.Msec%10 == 0 {
		uk.StatsInitDir(dir, "Msec")
		uk.StatsTime(dir.Dir("Msec10"))
		if uk.Msec%100 == 0 {
			uk.StatsInitDir(dir, "Msec10")
			uk.StatsTime(dir.Dir("Msec100"))
			// uk.UpdateTimePlots()
		}
	}
}

func (uk *Urakubo) GraphRun(secs float64, itr int) {
	nms := int(secs / 0.001)
	sms := uk.Msec
	for msec := 0; msec < nms; msec++ {
		uk.NeuronUpdt(sms+msec, 0, 0)
		uk.StatsDefault(itr)
		if uk.StopNow() {
			break
		}
	}
}

func (uk *Urakubo) RunQuiet(secs float64) {
	nms := int(secs / 0.001)
	sms := uk.Msec
	for msec := 0; msec < nms; msec++ {
		uk.NeuronUpdt(sms+msec, 0, 0)
		if uk.StopNow() {
			break
		}
	}
}
