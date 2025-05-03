// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// urakubo: This simulation replicates the Urakubo et al, 2008 detailed model of spike-driven
// learning, including intracellular Ca-driven signaling, involving CaMKII, CaN, PKA, PP1.
package urakubo

//go:generate core generate -add-types -add-funcs

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/base/mpi"
	"cogentcore.org/lab/base/randx"
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/egui"
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
	// Stim determines the stimulation protocol to drive the Neuron with.
	Stim Stims

	// Urakubo contains all the urakubo model state and parameters.
	Urakubo *Urakubo `new-window:"+"`

	// simulation configuration parameters -- set by .toml config file and / or args
	Config *Config `new-window:"+"`

	// Root is the root tensorfs directory, where all stats and other misc sim data goes.
	Root *tensorfs.Node `display:"-"`

	// Stats has the stats directory within Root.
	Stats *tensorfs.Node `display:"-"`

	// GUI manages all the GUI elements
	GUI egui.GUI `display:"-"`

	// RandSeeds is a list of random seeds to use for each run.
	RandSeeds randx.Seeds `display:"-"`
}

// RunSim runs the simulation as a standalone app
// with given configuration.
func RunSim(cfg *Config) error {
	ss := &Sim{Config: cfg}
	ss.ConfigSim()
	if ss.Config.GUI {
		ss.RunGUI()
	} else {
		ss.RunNoGUI()
	}
	return nil
}

// EmbedSim runs the simulation with default configuration
// embedded within given body element.
func EmbedSim(b tree.Node) *Sim {
	cfg := NewConfig()
	cfg.GUI = true
	ss := &Sim{Config: cfg}
	ss.ConfigSim()
	ss.Init()
	ss.ConfigGUI(b)
	return ss
}

func (ss *Sim) Defaults() {
	ss.Config.Defaults()
	ss.Stim = STDPSweep
}

func (ss *Sim) ConfigSim() {
	ss.Defaults()
	ss.Root, _ = tensorfs.NewDir("Root")
	tensorfs.CurRoot = ss.Root
	ss.Urakubo = NewUrakubo()
	ss.Urakubo.ConfigStats(ss.Root)
	ss.RandSeeds.Init(100) // max 100 runs
	ss.InitRandSeed(0)
}

////////  Init, utils

// Init restarts the run, and initializes everything, including network weights
// and resets the epoch log table
func (ss *Sim) Init() {
	ss.InitRandSeed(0)
	ss.Urakubo.Init()
}

// InitRandSeed initializes the random seed based on current training run number
func (ss *Sim) InitRandSeed(run int) {
	ss.RandSeeds.Set(run)
}

//////// GUI

// UpdateGUI updates the GUI window if GUI present
func (ss *Sim) UpdateGUI() {
	if !ss.GUI.Active {
		return
	}
	ss.GUI.UpdateWindow()
}

// ConfigGUI configures the Cogent Core GUI interface for this simulation.
func (ss *Sim) ConfigGUI(b tree.Node) {
	ss.GUI.MakeBody(b, ss, ss.Root, ss.Config.Name, ss.Config.Title, ss.Config.Doc)
	ss.GUI.CycleUpdateInterval = 10
	ss.Urakubo.GUI = &ss.GUI
	ss.Urakubo.StatsInit()
	ss.GUI.FinalizeGUI(false)
}

func (ss *Sim) MakeToolbar(p *tree.Plan) {
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "Init",
		Icon:    icons.Update,
		Tooltip: "Initialize sim and apply params.",
		Active:  egui.ActiveStopped,
		Func: func() {
			ss.Urakubo.Init()
			ss.Urakubo.StatsInit()
			ss.UpdateGUI()
		},
	})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "Stop",
		Icon:    icons.Stop,
		Tooltip: "Stops running.",
		Active:  egui.ActiveRunning,
		Func: func() {
			ss.Urakubo.Stop()
		},
	})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "Run",
		Icon:    icons.RunCircle,
		Tooltip: "Run current Stims.",
		Active:  egui.ActiveStopped,
		Func: func() {
			ss.Urakubo.Stim = ss.Stim
			ss.Urakubo.GUI.StartRun()
			go func() {
				ss.Urakubo.RunStim()
				ss.Urakubo.Stopped()
			}()
			ss.GUI.Toolbar.Restyle()
		},
	})

	tree.Add(p, func(w *core.Separator) {})

	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "Defaults",
		Icon:    icons.Update,
		Tooltip: "Reset parameters to defaults.",
		Active:  egui.ActiveStopped,
		Func: func() {
			ss.Urakubo.Defaults()
			ss.Urakubo.Init()
		},
	})
	/*
		tbar.AddSeparator("run-sep")

		tbar.AddAction(gi.ActOpts{Label: "Reset Plots", Icon: "update", Tooltip: "Reset Time Plots.", UpdateFunc: func(act *gi.Action) {
			act.SetActiveStateUpdt(!uk.IsRunning)
		}}, win.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			if !uk.IsRunning {
				uk.ResetTimePlots()
			}
		})

		tbar.AddAction(gi.ActOpts{Label: "Reset DWt Plot", Icon: "update", Tooltip: "Reset DWt Plot.", UpdateFunc: func(act *gi.Action) {
			act.SetActiveStateUpdt(!uk.IsRunning)
		}}, win.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			if !uk.IsRunning {
				uk.ResetDWtPlot()
			}
		})

		tbar.AddAction(gi.ActOpts{Label: "AutoK Plot", Icon: "update", Tooltip: "Plot AutoK function.", UpdateFunc: func(act *gi.Action) {
			act.SetActiveStateUpdt(!uk.IsRunning)
		}}, win.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			if !uk.IsRunning {
				uk.AutoK()
			}
		})

	*/

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

func (ss *Sim) RunGUI() {
	ss.Init()
	ss.ConfigGUI(nil)
	ss.GUI.Body.RunMainWindow()
}

func (ss *Sim) RunNoGUI() {
	ss.Init()
	if ss.Config.Params.Note != "" {
		mpi.Printf("Note: %s\n", ss.Config.Params.Note)
	}
	// todo: open files, run
}
