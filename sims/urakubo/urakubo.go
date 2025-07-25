// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/lab/plot"
	"cogentcore.org/lab/tensor"
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/axon/v2/axon"
	"github.com/emer/emergent/v2/egui"
)

// Urakubo has all of the state for the Urakubo model.
type Urakubo struct {
	// Stim determines the stimulation protocol to drive the Neuron with.
	Stim Stims

	// Net is the axon network for basic activity dynamics
	Net *axon.Network `display:"no-inline"`

	// the spine state with Urakubo intracellular model
	Spine Spine

	// extra neuron state for additional channels: Vgcc, AK
	NeuronEx NeuronEx `display:"no-inline"`

	// whether to initialize values to known baseline values at start
	InitBaseline bool

	// inter-stimulus-interval in seconds -- between reps
	ISISec float64

	// number of repetitions -- takes 100 to produce classic STDP
	NReps int `default:"100"`

	// number of seconds to run after the manipulation -- results are strongest after 100,
	// decaying somewhat after that point -- 20 shows similar qualitative results but weaker,
	// 50 is pretty close to 100 -- less than 20 not recommended.
	FinalSecs float64 `default:"20,50,100"`

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
	NMDAGbar float32 `default:"0,0.15"`

	// strength of GABAB current -- 0.2 default for posterior cortex
	GABABGbar float32 `default:"0,0.2"`

	// strength of Vgcc current -- 1.46 for distal per Migliore, but 0.12 reproduces net Ca current..
	VgccGbar float32 `default:"0,0.12"`

	// strength of Mahp current
	MahpGbar float32 `default:"0"`

	// target calcium level for CaTarg stim
	CaTarg CaState

	// initial weight value: Trp_AMPA value at baseline
	InitWt float64 `edit:"-"`

	// current cycle of updating
	Msec int `edit:"-"`

	Stats     *tensorfs.Node   `display:"-"`
	GUI       *egui.GUI        `display:"-"`
	StimFuncs map[Stims]func() `display:"-"`
}

func NewUrakubo() *Urakubo {
	uk := &Urakubo{}
	uk.Config()
	return uk
}

func (uk *Urakubo) Config() {
	uk.InitBaseline = true
	uk.Spine.Defaults()
	uk.Spine.Init(uk)
	uk.InitWt = uk.Spine.States.AMPAR.Trp.Tot
	uk.Net = axon.NewNetwork("Urakubo")
	uk.Stim = ThetaErrComp
	uk.ISISec = 0.8
	uk.NReps = 100
	uk.FinalSecs = 50
	uk.DurMsec = 200
	uk.SendHz = 50
	uk.RecvHz = 25
	uk.DeltaT = 16
	uk.DeltaTRange = 50
	uk.DeltaTInc = 5
	uk.RGClamp = true
	uk.Defaults()
	uk.ConfigNet(uk.Net)
	uk.ConfigStimFuncs()
}

// Defaults sets default params
func (uk *Urakubo) Defaults() {
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
	net.Context().ThetaCycles = 200
	// net.SetRandSeed(ss.RandSeeds[0]) // init new separate random seed, using run = 0

	net.AddLayer2D("Neuron", axon.SuperLayer, 1, 1)

	net.Build()
	net.Defaults()
	uk.ApplyParams()
	net.InitWeights()
}

func (uk *Urakubo) ApplyParams() {
	lsheet, err := LayerParams.SheetByName("Base")
	if err != nil {
		errors.Log(err)
		return
	}
	lsheet.SelMatchReset()

	axon.ApplyLayerSheet(uk.Net, lsheet)
}

// Init restarts the run, and initializes everything, including network weights
// and resets the epoch log table
func (uk *Urakubo) Init() {
	uk.Spine.Init(uk)
	uk.InitWt = uk.Spine.States.AMPAR.Trp.Tot
	uk.NeuronEx.Init()
	uk.Msec = 0
	// uk.SetParams("", false) // all sheets
	ly := uk.Net.LayerByName("Neuron")
	if uk.NMDAAxon {
		ly.Params.Acts.NMDA.Ge = uk.NMDAGbar
	} else {
		ly.Params.Acts.NMDA.Ge = 0
	}
	ly.Params.Acts.GabaB.Gk = uk.GABABGbar
	ly.Params.Acts.VGCC.Ge = uk.VgccGbar
}

// Counters returns a string of the current counter state
// use tabs to achieve a reasonable formatting overall
// and add a few tabs at the end to allow for expansion..
func (uk *Urakubo) Counters() string {
	return fmt.Sprintf("Msec:\t%d\t\t\t", uk.Msec)
}

func (uk *Urakubo) Stop() {
	if uk.GUI == nil {
		return
	}
	uk.GUI.SetStopNow()
}

func (uk *Urakubo) StopNow() bool {
	if uk.GUI == nil {
		return false
	}
	return uk.GUI.StopNow()
}

func (uk *Urakubo) Stopped() {
	if uk.GUI == nil {
		return
	}
	uk.GUI.Stopped(Test, Cycle)
}

// RunStim runs current Stim selection.
func (uk *Urakubo) RunStim() {
	fn, has := uk.StimFuncs[uk.Stim]
	if !has {
		fmt.Printf("Stim function: %s not found!\n", uk.Stim)
		return
	}
	fn()
}

// NeuronUpdate updates the neuron and spine for given msec
func (uk *Urakubo) NeuronUpdate(msec int, ge, gi float32) {
	ctx := uk.Net.Context()
	uk.Msec = msec
	ly := uk.Net.LayerByName("Neuron").Params
	ni := 0
	di := 0
	niu := uint32(0)
	diu := uint32(0)
	nex := &uk.NeuronEx

	vm := axon.Neurons.Value(ni, di, int(axon.Vm))
	// vmDend := axon.Neurons.Value(ni, di, int(axon.VmDend)) // todo: explore

	gnmda := uk.NMDAGbar * float32(uk.Spine.States.NMDAR.G)

	geSyn := axon.Neurons.Value(ni, di, int(axon.GeSyn))
	geSyn = ly.Acts.Dt.GeSynFromRaw(geSyn, ge)
	giSyn := ly.Acts.Dt.GeSynFromRawSteady(gi)

	geSyn += gnmda

	axon.Neurons.Set(ge, ni, di, int(axon.GeRaw))
	axon.Neurons.Set(geSyn, ni, di, int(axon.GeSyn))

	// note: gi always = 0 in this model
	axon.Neurons.Set(gi, ni, di, int(axon.GiRaw))
	axon.Neurons.Set(giSyn, ni, di, int(axon.GiSyn))

	ly.CycleNeuron(ctx, niu, diu)

	gvgcc := axon.Neurons.Value(ni, di, int(axon.Gvgcc))

	// todo: Ca from NMDAAxon
	uk.Spine.Ca.SetInject(float64(nex.VgccJcaPSD), float64(nex.VgccJcaCyt))

	psd_pca := float32(1.7927e5 * 0.04) //  SVR_PSD
	cyt_pca := float32(1.0426e5 * 0.04) // SVR_CYT

	nex.VgccJcaPSD = -vm * psd_pca * gvgcc
	nex.VgccJcaCyt = -vm * cyt_pca * gvgcc

	uk.Spine.States.VmS = float64(vm)

	uk.Spine.StepTime(0.001)
}

//////// Stats

var TimeStatsNames = []string{"Msec", "Msec10", "Msec100"}

var DWtStatsNames = []string{"DWt", "DWtPhase"}

// ConfigStats
func (uk *Urakubo) ConfigStats(dir *tensorfs.Node) {
	uk.Stats = dir.Dir("Stats")
}

// StatsInit initializes all the stats by setting num rows to 0
func (uk *Urakubo) StatsInit() {
	// todo: iterations
	dir := uk.Stats
	idx := 0
	uk.StatsInitTime()
	uk.StatsInitDWt()
	uk.StatsInitDWtPhase()
	if uk.GUI == nil || uk.GUI.Tabs == nil {
		return
	}
	_, idx = uk.GUI.Tabs.AsLab().CurrentTab()
	nms := append(TimeStatsNames, DWtStatsNames...)
	for _, sn := range nms {
		sd := dir.Dir(sn)
		uk.GUI.Tabs.AsLab().PlotTensorFS(sd)
	}
	if idx >= 0 {
		uk.GUI.Tabs.AsLab().SelectTabIndex(idx)
	}
}

// StatsInitTime
func (uk *Urakubo) StatsInitTime() {
	for _, sn := range TimeStatsNames {
		sd := uk.Stats.Dir(sn)
		uk.StatsTime(sd)
		uk.StatsInitDir(uk.Stats, sn)
	}
}

// StatsInitDWt
func (uk *Urakubo) StatsInitDWt() {
	sd := uk.Stats.Dir("DWt")
	uk.StatsDWt(sd, 0, 0)
	uk.StatsInitDir(uk.Stats, "DWt")
}

// StatsInitDWtPhase
func (uk *Urakubo) StatsInitDWtPhase() {
	sd := uk.Stats.Dir("DWtPhase")
	uk.StatsDWtPhase(sd, []int{0, 0}, []int{0, 0})
	uk.StatsInitDir(uk.Stats, "DWtPhase")
}

func (uk *Urakubo) StatsPlotUpdate(sn string) {
	if uk.GUI == nil {
		return
	}
	nm := "Stats " + sn + " Plot"
	uk.GUI.Tabs.AsLab().GoUpdatePlot(nm)
}

// StatsInitDir initializes given stats directory
func (uk *Urakubo) StatsInitDir(dir *tensorfs.Node, sn string) {
	sd := dir.Dir(sn)
	tsrs, _ := sd.Values()
	for _, tsr := range tsrs {
		tsr.(tensor.Values).SetNumRows(0)
	}
}

// StatsTime adds basic timestep state data to given tensorfs directory.
func (uk *Urakubo) StatsTime(dir *tensorfs.Node) {
	nex := &uk.NeuronEx
	ni := 0
	di := 0

	dir.Float64("Time").AppendRowFloat(float64(uk.Msec) * 0.001)
	dir.Float64("Ge").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Ge))))
	dir.Float64("Inet").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Inet))))
	tsr := dir.Float64("Vm")
	tsr.AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Vm))))
	plot.SetFirstStyler(tsr, func(s *plot.Style) {
		s.On = true
		s.RightY = true
	})
	dir.Float64("Act").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Act))))
	dir.Float64("Spike").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Spike))))
	dir.Float64("Gk").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gk))))
	dir.Float64("ISI").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.ISI))))
	dir.Float64("AvgISI").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.ISIAvg))))
	dir.Float64("VmDend").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.VmDend))))
	// dir.Float64("SnmdaO").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.SnmdaO))))
	dir.Float64("Gnmda").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.Gnmda))))
	dir.Float64("GababM").AppendRowFloat(float64(axon.Neurons.Value(ni, di, int(axon.GababM))))
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

// StatsTimeUpdate does plot update after
func (uk *Urakubo) StatsTimeUpdate(dir *tensorfs.Node) {
	uk.StatsTime(dir)
	uk.StatsPlotUpdate(dir.Name())
}

// StatsDWt adds data for current dwt value as function of x, y values
func (uk *Urakubo) StatsDWt(dir *tensorfs.Node, x, y float64) {
	tsr := dir.Float64("X")
	tsr.AppendRowFloat(x)
	plot.SetFirstStyler(tsr, func(s *plot.Style) {
		s.Role = plot.X
	})
	dir.Float64("Y").AppendRowFloat(y)

	wt := uk.Spine.States.AMPAR.Trp.Tot
	dwt := (wt / uk.InitWt) - 1

	tsr = dir.Float64("DWt")
	tsr.AppendRowFloat(float64(dwt))
	plot.SetFirstStyler(tsr, func(s *plot.Style) {
		s.On = true
	})
	uk.Spine.Stats(dir)
}

// StatsDWtUpdate does plot update after
func (uk *Urakubo) StatsDWtUpdate(dir *tensorfs.Node, x, y float64) {
	uk.StatsDWt(dir, x, y)
	uk.StatsPlotUpdate(dir.Name())
}

// StatsDWtPhase adds data for current dwt value as function of phase hz levels
func (uk *Urakubo) StatsDWtPhase(dir *tensorfs.Node, sphz, rphz []int) {
	chl := (float64(sphz[1])/100.0)*(float64(rphz[1])/100.0) - (float64(sphz[0])/100.0)*(float64(rphz[0])/100.0)

	dir.Float64("CHL").AppendRowFloat(chl)
	dir.Float64("SMhz").AppendRowFloat(float64(sphz[0]))
	dir.Float64("SPhz").AppendRowFloat(float64(sphz[1]))
	dir.Float64("RMhz").AppendRowFloat(float64(rphz[0]))
	dir.Float64("RPhz").AppendRowFloat(float64(rphz[1]))

	wt := uk.Spine.States.AMPAR.Trp.Tot
	dwt := (wt / uk.InitWt) - 1

	tsr := dir.Float64("DWt")
	tsr.AppendRowFloat(float64(dwt))
	plot.SetFirstStyler(tsr, func(s *plot.Style) {
		s.On = true
	})
	uk.Spine.Stats(dir)
}

// StatsDWtPhaseUpdate does plot update after.
func (uk *Urakubo) StatsDWtPhaseUpdate(dir *tensorfs.Node, sphz, rphz []int) {
	uk.StatsDWtPhase(dir, sphz, rphz)
	uk.StatsPlotUpdate(dir.Name())
}

// StatsDefault does default logging for current Msec, for given iteration.
func (uk *Urakubo) StatsDefault(itr int) {
	// todo: deal with iteration as a subdir in stats
	dir := uk.Stats
	msec := uk.Msec
	uk.StatsTimeUpdate(dir.Dir("Msec"))
	if msec%10 == 0 {
		uk.StatsTimeUpdate(dir.Dir("Msec10"))
		if msec%100 == 0 {
			uk.StatsInitDir(dir, "Msec")
			if msec%1000 == 0 {
				uk.StatsInitDir(dir, "Msec10")
			}
			uk.StatsTimeUpdate(dir.Dir("Msec100"))
		}
	}
}

func (uk *Urakubo) RunWithStats(secs float64, itr int) {
	nms := int(secs / 0.001)
	sms := uk.Msec
	for msec := 0; msec < nms; msec++ {
		uk.NeuronUpdate(sms+msec, 0, 0)
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
		uk.NeuronUpdate(sms+msec, 0, 0)
		if uk.StopNow() {
			break
		}
	}
}
