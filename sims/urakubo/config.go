// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/reflectx"
)

// ParamConfig has config parameters related to sim params.
type ParamConfig struct {

	// Script is an interpreted script that is run to set parameters in Layer and Path
	// sheets, by default using the "Script" set name.
	Script string `new-window:"+" width:"100"`

	// Sheet is the extra params sheet name(s) to use (space separated
	// if multiple). Must be valid name as listed in compiled-in params
	// or loaded params.
	Sheet string

	// Tag is an extra tag to add to file names and logs saved from this run.
	Tag string

	// Note is additional info to describe the run params etc,
	// like a git commit message for the run.
	Note string

	// SaveAll will save a snapshot of all current param and config settings
	// in a directory named params_<datestamp> (or _good if Good is true),
	// then quit. Useful for comparing to later changes and seeing multiple
	// views of current params.
	SaveAll bool `nest:"+"`

	// Good is for SaveAll, save to params_good for a known good params state.
	// This can be done prior to making a new release after all tests are passing.
	// Add results to git to provide a full diff record of all params over level.
	Good bool `nest:"+"`
}

// RunConfig has config parameters related to running the sim.
type RunConfig struct {

	// Use 500 sec pre-compiled baseline for initialization
	InitBaseline bool `def:"true"`

	// Trials is the total number of trials of different random patterns to generate.
	Trials int `default:"10"`

	// Cycles is the total number of cycles per trial: at least 200.
	Cycles int `default:"200"`

	// PlusCycles is the total number of plus-phase cycles per trial. For Cycles=300, use 100.
	PlusCycles int `default:"50"`
}

// LogConfig has config parameters related to logging data.
type LogConfig struct {

	// Save has the list of levels to save log files for.
	Save []string `default:"['Trial', 'Cycle']" nest:"+"`
}

// Config has the overall Sim configuration options.
type Config struct {

	// Name is the short name of the sim.
	Name string `display:"-" default:"Urakubo"`

	// Title is the longer title of the sim.
	Title string `display:"-" default:"Axon urakubo test"`

	// URL is a link to the online README or other documentation for this sim.
	URL string `display:"-" default:"https://github.com/emer/axon/blob/main/examples/inhib/README.md"`

	// Doc is brief documentation of the sim.
	Doc string `display:"-" default:"This simulation replicates the Urakubo et al, 2008 detailed model of spike-driven learning, including intracellular Ca-driven signaling, involving CaMKII, CaN, PKA, PP1."`

	// Includes has a list of additional config files to include.
	// After configuration, it contains list of include files added.
	Includes []string `display:"-"`

	// GUI means open the GUI. Otherwise it runs automatically and quits,
	// saving results to log files.
	GUI bool `default:"true"`

	// Debug reports debugging information.
	Debug bool

	// Params has parameter related configuration options.
	Params ParamConfig `display:"add-fields"`

	// Run has sim running related configuration options.
	Run RunConfig `display:"add-fields"`

	// Log has data logging related configuration options.
	Log LogConfig `display:"add-fields"`
}

func (cfg *Config) IncludesPtr() *[]string { return &cfg.Includes }

func (cfg *Config) Defaults() {
	errors.Log(reflectx.SetFromDefaultTags(cfg))
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.Defaults()
	return cfg
}
