// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// stability: This simulation simply presents an input pattern and allows
// observation of the resulting stability of the network response, while
// varying the strength of the NMDA and GABA-B channel conductances, to
// test for their contributions in producing a stable activity pattern.
package stability

//go:generate core generate -add-types -add-funcs

import (
	"fmt"
	"os"
	"reflect"

	"cogentcore.org/core/base/metadata"
	"cogentcore.org/core/core"
	"cogentcore.org/core/enums"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/base/mpi"
	"cogentcore.org/lab/base/randx"
	"cogentcore.org/lab/patterns"
	"cogentcore.org/lab/plot"
	"cogentcore.org/lab/stats/stats"
	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/axon/v2/axon"
	"github.com/emer/axon/v2/fsfffb"
	"github.com/emer/emergent/v2/egui"
	"github.com/emer/emergent/v2/env"
	"github.com/emer/emergent/v2/looper"
	"github.com/emer/emergent/v2/paths"
)

// Modes are the looping modes (Stacks) for running and statistics.
type Modes int32 //enums:enum
const (
	Test  Modes = iota
	Train       // not used, but needed for some things
)

// Levels are the looping levels for running and statistics.
type Levels int32 //enums:enum
const (
	Cycle Levels = iota
	Trial
	Epoch
)

// StatsPhase is the phase of stats processing for given mode, level.
// Accumulated values are reset at Start, added each Step.
type StatsPhase int32 //enums:enum
const (
	Start StatsPhase = iota
	Step
)

// see params.go for params

// Sim encapsulates the entire simulation model, and we define all the
// functionality as methods on this struct.  This structure keeps all relevant
// state information organized and available without having to pass everything around
// as arguments to methods, and provides the core GUI interface (note the view tags
// for the fields which provide hints to how things should be displayed).
type Sim struct {

	// NmdaGe is the strength of contribution of the NMDA excitatory Ca++ current,
	// to the overall Ge(t) excitatory conductance value. This channel
	// has a long time constant and is essential for establishing
	// a more stable neural representation over time by keeping active neurons active.
	NmdaGe float32 `default:"0.006" step:"0.001"`

	// GababGk is the strength of contribution of the GABA-B inhibitory K current,
	// to the overall Gk(t) inhibitory potassium (K) conductance value. This channel
	// also has a long time constant like NMDA, and works in opposition to it,
	// by keeping inactive neurons inactive, synergistically helping to establish
	// stable neural representations.
	GababGk float32 `default:"0.015" step:"0.005"`

	// InhibGi is the strength of the automatically-computed inhibitory feedback function.
	InhibGi float32 `default:"0.85" step:"0.01" min:"0.5"`

	// simulation configuration parameters -- set by .toml config file and / or args
	Config *Config `new-window:"+"`

	// Net is the network: click to view / edit parameters for layers, paths, etc.
	Net *axon.Network `new-window:"+" display:"no-inline"`

	// Params manages network parameter setting.
	Params axon.Params `display:"inline"`

	// Loops are the control loops for running the sim, in different Modes
	// across stacks of Levels.
	Loops *looper.Stacks `new-window:"+" display:"no-inline"`

	// Envs provides mode-string based storage of environments.
	Envs env.Envs `new-window:"+" display:"no-inline"`

	// NetUpdate has netview update parameters.
	NetUpdate axon.NetViewUpdate `display:"inline"`

	// Root is the root tensorfs directory, where all stats and other misc sim data goes.
	Root *tensorfs.Node `display:"-"`

	// Stats has the stats directory within Root.
	Stats *tensorfs.Node `display:"-"`

	// Current has the current stats values within Stats.
	Current *tensorfs.Node `display:"-"`

	// StatFuncs are statistics functions called at given mode and level,
	// to perform all stats computations. phase = Start does init at start of given level,
	// and all intialization / configuration (called during Init too).
	StatFuncs []func(mode Modes, level Levels, phase StatsPhase) `display:"-"`

	// GUI manages all the GUI elements
	GUI egui.GUI `display:"-"`

	// RandSeeds is a list of random seeds to use for each run.
	RandSeeds randx.Seeds `display:"-"`
}

func Embed(b tree.Node)               { egui.Embed[Sim, Config](b) }
func (ss *Sim) SetConfig(cfg *Config) { ss.Config = cfg }
func (ss *Sim) Body() *core.Body      { return ss.GUI.Body }

func (ss *Sim) Defaults() {
	ss.Config.Defaults()
	ss.NmdaGe = 0.006
	ss.GababGk = 0.015
	ss.InhibGi = 0.85
}

func (ss *Sim) ConfigSim() {
	ss.Defaults()
	ss.Root, _ = tensorfs.NewDir("Root")
	tensorfs.CurRoot = ss.Root
	ss.Net = axon.NewNetwork(ss.Config.Name)
	ss.Params.Config(LayerParams, PathParams, ss.Config.Params.Sheet, ss.Config.Params.Tag, reflect.ValueOf(ss))
	ss.RandSeeds.Init(100) // max 100 runs
	ss.InitRandSeed(0)
	ss.ConfigInputs()
	ss.ConfigEnv()
	ss.ConfigNet(ss.Net)
	ss.ConfigLoops()
	ss.ConfigStats()
	if ss.Config.Params.SaveAll {
		ss.Config.Params.SaveAll = false
		ss.Net.SaveParamsSnapshot(&ss.Config, ss.Config.Params.Good)
		os.Exit(0)
	}
}

func (ss *Sim) ConfigEnv() {
	// Can be called multiple times -- don't re-create
	var tst *env.FixedTable
	if len(ss.Envs) == 0 {
		tst = &env.FixedTable{}
	} else {
		tst = ss.Envs.ByMode(Test).(*env.FixedTable)
	}

	inputs := tensorfs.DirTable(ss.Root.Dir("Inputs/Test"), nil)

	tst.Name = Test.String()
	tst.Config(table.NewView(inputs))
	tst.Sequential = true
	tst.Validate()

	tst.Init(0)

	// note: names must be in place when adding
	ss.Envs.Add(tst)
}

func (ss *Sim) ReConfigNet() {
	ss.Net.DeleteAll()
	ss.ConfigNet(ss.Net)
	// ss.GUI.NetView.Config()
}

func LayNm(n int) string {
	return fmt.Sprintf("Layer%d", n)
}

func LayByNm(net *axon.Network, n int) *axon.Layer {
	return net.LayerByName(LayNm(n))
}

func (ss *Sim) ConfigNet(net *axon.Network) {
	net.SetMaxData(1)
	net.Context().ThetaCycles = int32(ss.Config.Run.Cycles)
	net.SetRandSeed(ss.RandSeeds[0]) // init new separate random seed, using run = 0

	sz := ss.Config.Params.HiddenSize

	inlay := net.AddLayer2D(LayNm(0), axon.InputLayer, sz.Y, sz.X)
	_ = inlay

	for hi := 1; hi <= ss.Config.Params.NLayers; hi++ {
		net.AddLayer2D(LayNm(hi), axon.SuperLayer, sz.Y, sz.X)
	}

	full := paths.NewFull()
	rndcut := paths.NewUniformRand()
	rndcut.PCon = 0.1

	for hi := 1; hi <= ss.Config.Params.NLayers; hi++ {
		ll := LayByNm(net, hi-1)
		tl := LayByNm(net, hi)
		net.ConnectLayers(ll, tl, full, axon.ForwardPath).AddClass("Excite")
		tl.PlaceAbove(ll)
		if hi < ss.Config.Params.NLayers {
			nl := LayByNm(net, hi+1)
			net.ConnectLayers(tl, nl, full, axon.ForwardPath).AddClass("Excite")
			net.ConnectLayers(nl, tl, full, axon.BackPath).AddClass("Excite")
		}
	}

	net.Build()
	net.Defaults()
	ss.ApplyParams()
	net.InitWeights()
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

////////  Init, utils

// Init restarts the run, and initializes everything, including network weights
// and resets the epoch log table
func (ss *Sim) Init() {
	ss.Loops.ResetCounters()
	ss.SetRunName()
	ss.InitRandSeed(0)
	ss.ApplyParams()
	ss.StatsInit()
	ss.NewRun()
	ss.NetUpdate.RecordSyns()
	ss.NetUpdate.Update(Test, Trial)
}

// InitRandSeed initializes the random seed based on current training run number
func (ss *Sim) InitRandSeed(run int) {
	ss.RandSeeds.Set(run)
	ss.RandSeeds.Set(run, &ss.Net.Rand)
}

// NetViewUpdater returns the NetViewUpdate for given mode.
func (ss *Sim) NetViewUpdater(mode enums.Enum) *axon.NetViewUpdate {
	return &ss.NetUpdate
}

// ConfigLoops configures the control loops: Training, Testing
func (ss *Sim) ConfigLoops() {
	ls := looper.NewStacks()

	cycles := ss.Config.Run.Cycles
	plusPhase := ss.Config.Run.PlusCycles

	ls.AddStack(Test, Trial).
		AddLevel(Epoch, 1).
		AddLevelIncr(Trial, ss.Config.Run.Trials, 1).
		AddLevel(Cycle, cycles)

	axon.LooperStandard(ls, ss.Net, ss.NetViewUpdater, cycles-plusPhase, cycles-1, Cycle, Trial, Train)

	ls.Stacks[Test].OnInit.Add("Init", func() { ss.Init() })

	ls.AddOnStartToLoop(Trial, "ApplyInputs", func(mode enums.Enum) {
		ss.ApplyInputs(mode.(Modes))
	})

	ls.AddOnStartToAll("StatsStart", ss.StatsStart)
	ls.AddOnEndToAll("StatsStep", ss.StatsStep)

	if ss.Config.GUI {
		axon.LooperUpdateNetView(ls, Cycle, Trial, ss.NetViewUpdater)
		ls.Stacks[Test].OnInit.Add("GUI-Init", func() { ss.GUI.UpdateWindow() })
	}

	if ss.Config.Debug {
		mpi.Println(ls.DocString())
	}
	ss.Loops = ls
}

// ApplyInputs applies input patterns from given environment for given mode.
// Any other start-of-trial logic can also be put here.
func (ss *Sim) ApplyInputs(mode Modes) {
	net := ss.Net
	curModeDir := ss.Current.Dir(mode.String())
	ev := ss.Envs.ByMode(mode)
	lays := net.LayersByType(axon.InputLayer, axon.TargetLayer)
	net.InitExt()
	ev.Step()
	curModeDir.StringValue("TrialName", 1).SetString1D(ev.String(), 0)
	for _, lnm := range lays {
		ly := ss.Net.LayerByName(lnm)
		st := ev.State("Input")
		if st != nil {
			ly.ApplyExt(uint32(0), st)
		}
	}
	net.ApplyExts()
}

// NewRun intializes a new Run level of the model.
func (ss *Sim) NewRun() {
	ctx := ss.Net.Context()
	ss.InitRandSeed(0)
	ss.Envs.ByMode(Test).Init(0)
	ctx.Reset()
	ss.Net.InitWeights()
}

////////  Inputs

func (ss *Sim) ConfigInputs() {
	dt := table.New()
	metadata.SetName(dt, "Test")
	metadata.SetDoc(dt, "Testing inputs")
	dt.AddStringColumn("Name")
	dt.AddFloat32Column("Input", 10, 10)
	dt.SetNumRows(25)

	patterns.PermutedBinaryMinDiff(dt.Columns.Values[1], int(ss.Config.Params.InputPct), 1, 0, int(ss.Config.Params.InputPct)/2)

	tensorfs.DirFromTable(ss.Root.Dir("Inputs/Test"), dt)
}

//////// Stats

// AddStat adds a stat compute function.
func (ss *Sim) AddStat(f func(mode Modes, level Levels, phase StatsPhase)) {
	ss.StatFuncs = append(ss.StatFuncs, f)
}

// StatsStart is called by Looper at the start of given level, for each iteration.
// It needs to call RunStats Start at the next level down.
// e.g., each Epoch is the start of the full set of Trial Steps.
func (ss *Sim) StatsStart(lmd, ltm enums.Enum) {
	mode := lmd.(Modes)
	level := ltm.(Levels)
	if level < Trial {
		return
	}
	ss.RunStats(mode, level-1, Start)
}

// StatsStep is called by Looper at each step of iteration,
// where it accumulates the stat results.
func (ss *Sim) StatsStep(lmd, ltm enums.Enum) {
	mode := lmd.(Modes)
	level := ltm.(Levels)
	ss.RunStats(mode, level, Step)
	tensorfs.DirTable(axon.StatsNode(ss.Stats, mode, level), nil).WriteToLog()
}

// RunStats runs the StatFuncs for given mode, level and phase.
func (ss *Sim) RunStats(mode Modes, level Levels, phase StatsPhase) {
	for _, sf := range ss.StatFuncs {
		sf(mode, level, phase)
	}
	if phase == Step && ss.GUI.Tabs != nil {
		nm := mode.String() + " " + level.String() + " Plot"
		ss.GUI.Tabs.AsLab().GoUpdatePlot(nm)
	}
}

// SetRunName sets the overall run name, used for naming output logs and weight files
// based on params extra sheets and tag, and starting run number (for distributed runs).
func (ss *Sim) SetRunName() string {
	runName := ss.Params.RunName(0)
	ss.Current.StringValue("RunName", 1).SetString1D(runName, 0)
	return runName
}

// RunName returns the overall run name, used for naming output logs and weight files
// based on params extra sheets and tag, and starting run number (for distributed runs).
func (ss *Sim) RunName() string {
	return ss.Current.StringValue("RunName", 1).String1D(0)
}

// StatsInit initializes all the stats by calling Start across all modes and levels.
func (ss *Sim) StatsInit() {
	for md, st := range ss.Loops.Stacks {
		mode := md.(Modes)
		for _, lev := range st.Order {
			level := lev.(Levels)
			ss.RunStats(mode, level, Start)
		}
	}
	if ss.GUI.Tabs != nil {
		tbs := ss.GUI.Tabs.AsLab()
		_, idx := tbs.CurrentTab()
		tbs.PlotTensorFS(axon.StatsNode(ss.Stats, Test, Cycle))
		tbs.PlotTensorFS(axon.StatsNode(ss.Stats, Test, Trial))
		tbs.SelectTabIndex(idx)
	}
}

// ConfigStats handles configures functions to do all stats computation
// in the tensorfs system.
func (ss *Sim) ConfigStats() {
	net := ss.Net
	ss.Stats = ss.Root.Dir("Stats")
	ss.Current = ss.Stats.Dir("Current")

	ss.SetRunName()

	// last arg(s) are levels to exclude
	counterFunc := axon.StatLoopCounters(ss.Stats, ss.Current, ss.Loops, net, Trial)
	ss.AddStat(func(mode Modes, level Levels, phase StatsPhase) {
		counterFunc(mode, level, phase == Start)
	})
	runNameFunc := axon.StatRunName(ss.Stats, ss.Current, ss.Loops, net, Trial)
	ss.AddStat(func(mode Modes, level Levels, phase StatsPhase) {
		runNameFunc(mode, level, phase == Start)
	})
	trialNameFunc := axon.StatTrialName(ss.Stats, ss.Current, ss.Loops, net, Trial)
	ss.AddStat(func(mode Modes, level Levels, phase StatsPhase) {
		trialNameFunc(mode, level, phase == Start)
	})

	// up to a point, it is good to use loops over stats in one function,
	// to reduce repetition of boilerplate.
	layers := ss.Net.LayersByType(axon.SuperLayer) // axon.InputLayer,
	statNames := []string{"Spikes", "Ge", "Act", "Gi"}
	ss.AddStat(func(mode Modes, level Levels, phase StatsPhase) {
		for _, lnm := range layers {
			ly := ss.Net.LayerByName(lnm)
			pi := int(ly.Params.PoolIndex(0))
			di := 0
			for _, stnm := range statNames {
				name := lnm + "_" + stnm
				modeDir := ss.Stats.Dir(mode.String())
				curModeDir := ss.Current.Dir(mode.String())
				levelDir := modeDir.Dir(level.String())
				tsr := levelDir.Float64(name)
				ndata := 1
				if phase == Start {
					tsr.SetNumRows(0)
					plot.SetFirstStyler(tsr, func(s *plot.Style) {
						s.Range.SetMin(0).SetMax(1)
						s.On = false
						switch stnm {
						case "TotalGi":
							s.On = true
						}
					})
					continue
				}
				switch level {
				case Cycle:
					var stat float32
					switch stnm {
					case "Spikes":
						stat = ly.AvgMaxVarByPool("Spike", 0, di).Avg
					case "Ge":
						stat = axon.PoolAvgMax(axon.AMGeInt, axon.AMCycle, axon.Avg, uint32(pi), uint32(di))
					case "Act":
						stat = axon.PoolAvgMax(axon.AMAct, axon.AMCycle, axon.Avg, uint32(pi), uint32(di))
					case "Gi":
						stat = axon.Neurons.Value(int(ly.NeurStIndex), di, int(axon.Gi))
					default:
						var ivar fsfffb.InhibVars
						ivar.SetString(stnm)
						stat = axon.Pools.Value(pi, di, int(ivar))
					}
					curModeDir.Float64(name, ndata).SetFloat1D(float64(stat), di)
					tsr.AppendRowFloat(float64(stat))
				default:
					subDir := modeDir.Dir((level - 1).String())
					stat := stats.StatMean.Call(subDir.Value(name)).Float1D(0)
					tsr.AppendRowFloat(stat)
				}
			}
		}
	})

	perTrlFunc := axon.StatPerTrialMSec(ss.Stats, Test, Trial)
	ss.AddStat(func(mode Modes, level Levels, phase StatsPhase) {
		perTrlFunc(mode, level, phase == Start)
	})

}

// StatCounters returns counters string to show at bottom of netview.
func (ss *Sim) StatCounters(mode, level enums.Enum) string {
	counters := ss.Loops.Stacks[mode].CountersString()
	vu := ss.NetViewUpdater(mode)
	if vu == nil || vu.View == nil {
		return counters
	}
	di := vu.View.Di
	counters += fmt.Sprintf(" Di: %d", di)
	curModeDir := ss.Current.Dir(mode.String())
	if curModeDir.Node("TrialName") == nil {
		return counters
	}
	counters += fmt.Sprintf(" TrialName: %s", curModeDir.StringValue("TrialName").String1D(di))
	statNames := []string{"CorSim", "UnitErr", "Err"}
	if level == Cycle || curModeDir.Node(statNames[0]) == nil {
		return counters
	}
	for _, name := range statNames {
		counters += fmt.Sprintf(" %s: %.4g", name, curModeDir.Float64(name).Float1D(di))
	}
	return counters
}

//////// GUI

// ConfigGUI configures the Cogent Core GUI interface for this simulation.
func (ss *Sim) ConfigGUI(b tree.Node) {
	ss.GUI.MakeBody(b, ss, ss.Root, ss.Config.Name, ss.Config.Title, ss.Config.Doc)
	ss.GUI.CycleUpdateInterval = 10

	nv := ss.GUI.AddNetView("Network")
	nv.Options.MaxRecs = 2 * ss.Config.Run.Cycles
	nv.Options.Raster.Max = ss.Config.Run.Cycles
	nv.SetNet(ss.Net)
	ss.NetUpdate.Config(nv, axon.Theta, ss.StatCounters)
	ss.GUI.OnStop = func(mode, level enums.Enum) {
		vu := ss.NetViewUpdater(mode)
		vu.UpdateWhenStopped(mode, level)
	}

	nv.SceneXYZ().Camera.Pose.Pos.Set(0, 1.3, 2.6)
	nv.SceneXYZ().Camera.LookAt(math32.Vec3(0, 0, 0), math32.Vec3(0, 1, 0))

	ss.StatsInit()
	ss.GUI.FinalizeGUI(false)
}

func (ss *Sim) MakeToolbar(p *tree.Plan) {
	ss.GUI.AddLooperCtrl(p, ss.Loops)

	tree.Add(p, func(w *core.Separator) {})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "New Seed",
		Icon:    icons.Add,
		Tooltip: "Generate a new initial random seed to get different results.  By default, Init re-establishes the same initial seed every time.",
		Active:  egui.ActiveAlways,
		Func: func() {
			ss.RandSeeds.NewSeeds()
		},
	})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "README",
		Icon:    icons.FileMarkdown,
		Tooltip: "Opens your browser on the README file that contains instructions for how to run this model.",
		Active:  egui.ActiveAlways,
		Func: func() {
			core.TheApp.OpenURL(ss.Config.URL)
		},
	})
}

func (ss *Sim) RunNoGUI() {
	ss.Init()

	if ss.Config.Params.Note != "" {
		mpi.Printf("Note: %s\n", ss.Config.Params.Note)
	}

	runName := ss.SetRunName()
	netName := ss.Net.Name
	cfg := &ss.Config.Log
	axon.OpenLogFiles(ss.Loops, ss.Stats, netName, runName, [][]string{cfg.Save})

	ss.Loops.Run(Test)

	axon.CloseLogFiles(ss.Loops, ss.Stats)
	axon.GPURelease()
}
