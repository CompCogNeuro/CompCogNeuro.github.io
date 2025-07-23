// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

const (
	CytVol = 48 // volume of cytosol, in essentially arbitrary units
	PSDVol = 12 // volume of PSD
)

func init() {
	chem.IntegrationDt = 5e-6
}

// CaSigState is entire intracellular Ca-driven signaling state
// Total state vars: 2 + 32 + 14 + 32 + 14 + 1 = 95
type CaSigState struct {
	// Ca state
	Ca CaState

	// CaM calmodulin state
	CaM CaMState

	// CaMKII state
	CaMKII CaMKIIState

	// CaN = calcineurin state
	CaN CaNState

	// PKA = protein kinase A
	PKA PKAState

	// PP1 = protein phosphatase 1
	PP1 PP1State

	// PP2A = protein phosphatase 2A, only in Cyt
	PP2A float64
}

func (cs *CaSigState) Init(uk *Urakubo) {
	cs.Ca.Init(uk)
	cs.CaM.Init(uk)
	cs.CaMKII.Init(uk)
	cs.CaN.Init(uk)
	cs.PKA.Init(uk)
	cs.PP1.Init(uk)

	cs.PP2A = chem.CoToN(0.03, CytVol)

	if uk.InitBaseline {
		cs.PP2A = chem.CoToN(0.02239, CytVol)
	}
}

func (cs *CaSigState) InitCode() {
	cs.CaM.InitCode()
	cs.CaMKII.InitCode()
	cs.CaN.InitCode()
	cs.PKA.InitCode()
	cs.PP1.InitCode()
	fmt.Printf("\nCaSigState:\n")
	fmt.Printf("\tcs.PP2A = chem.CoToN(%.4g, CytVol)\n", chem.CoFromN(cs.PP2A, CytVol))
}

func (cs *CaSigState) Zero() {
	cs.Ca.Zero()
	cs.CaM.Zero()
	cs.CaMKII.Zero()
	cs.CaN.Zero()
	cs.PKA.Zero()
	cs.PP1.Zero()
	cs.PP2A = 0
}

func (cs *CaSigState) Integrate(d *CaSigState) {
	cs.Ca.Integrate(&d.Ca)
	cs.CaM.Integrate(&d.CaM)
	cs.CaMKII.Integrate(&d.CaMKII)
	cs.CaN.Integrate(&d.CaN)
	cs.PKA.Integrate(&d.PKA)
	cs.PP1.Integrate(&d.PP1)
	chem.Integrate(&cs.PP2A, d.PP2A)
}

func (cs *CaSigState) Stats(dir *tensorfs.Node) {
	cs.Ca.Stats(dir)
	cs.CaM.Stats(dir)
	cs.CaMKII.Stats(dir)
	cs.CaN.Stats(dir)
	cs.PKA.Stats(dir)
	cs.PP1.Stats(dir)
}

// SpineState is entire state of spine including Ca signaling and AMPAR
// Total state vars: 95 + 20 = 115
type SpineState struct {

	// internal time counter, in seconds, incremented by Dt
	Time float64

	// NMDA receptor state
	NMDAR NMDARState

	// calcium signaling systems
	CaSig CaSigState

	// AMPA receptor state
	AMPAR AMPARState

	// Vm in spine
	VmS float64

	// discrete spike firing -- 0 = no spike, 1 = spike
	PreSpike float64

	// time of last spike firing -- needed to prevent repeated spiking from same singal
	PreSpikeT float64
}

func (ss *SpineState) Init(uk *Urakubo) {
	ss.Time = 0
	ss.NMDAR.Init(uk)
	ss.CaSig.Init(uk)
	ss.AMPAR.Init(uk)
	ss.VmS = -65
	ss.PreSpike = 0
	ss.PreSpikeT = -1
}

func (ss *SpineState) InitCode() {
	ss.CaSig.InitCode()
	ss.AMPAR.InitCode()
}

func (ss *SpineState) Zero() {
	ss.Time = 0
	ss.NMDAR.Zero()
	ss.CaSig.Zero()
	ss.AMPAR.Zero()
	ss.VmS = 0
	ss.PreSpike = 0
	ss.PreSpikeT = 0
}

func (ss *SpineState) Integrate(d *SpineState) {
	ss.Time += chem.IntegrationDt
	ss.NMDAR.Integrate(&d.NMDAR)
	ss.CaSig.Integrate(&d.CaSig)
	ss.AMPAR.Integrate(&d.AMPAR)
}

func (ss *SpineState) Stats(dir *tensorfs.Node) {
	dir.Float64("VmS").AppendRowFloat(ss.VmS)
	dir.Float64("PreSpike").AppendRowFloat(ss.PreSpike)
	ss.NMDAR.Stats(dir)
	ss.CaSig.Stats(dir)
	ss.AMPAR.Stats(dir)
}

// Spine represents all of the state and parameters of the Spine
// involved in LTP / LTD
type Spine struct {

	// NMDA receptors
	NMDAR NMDARParams

	// Ca buffering and diffusion parameters
	Ca CaParams

	// CaM calmodulin Ca binding parameters
	CaM CaMParams

	// CaMKII parameters
	CaMKII CaMKIIParams

	// CaN calcineurin parameters
	CaN CaNParams

	// PKA = protein kinase A parameters
	PKA PKAParams

	// PP1 = protein phosphatase 1 parameters
	PP1 PP1Params

	// AMPAR parameters
	AMPAR AMPARParams

	// the current spine states
	States SpineState

	// the derivative changes in spine states
	Deltas SpineState
}

func (sp *Spine) Defaults() {
	sp.NMDAR.Defaults()
	sp.Ca.Defaults()
	sp.CaM.Defaults()
	sp.CaMKII.Defaults()
	sp.CaN.Defaults()
	sp.PKA.Defaults()
	sp.PP1.Defaults()
	sp.AMPAR.Defaults()
	// fmt.Printf("Integration Dt = %g (%g steps per msec)\n", chem.IntegrationDt, 0.001/chem.IntegrationDt)
}

func (sp *Spine) Init(uk *Urakubo) {
	sp.States.Init(uk)
	sp.Deltas.Zero()
	sp.Ca.Init()                    // drivers
	sp.NMDAR.Init(&sp.States.NMDAR) // special init
}

func (sp *Spine) InitCode() {
	sp.States.InitCode()
}

// Step computes the new Delta values
func (sp *Spine) Step() {
	sp.Deltas.Zero()

	vms := sp.States.VmS
	preSpike := false
	if sp.States.PreSpike > 0 {
		if sp.States.Time-sp.States.PreSpikeT > 0.003 { // refractory period
			preSpike = true
			sp.States.PreSpikeT = sp.States.Time
		}
	}

	sp.NMDAR.Step(&sp.States.NMDAR, &sp.Deltas.NMDAR, vms, chem.CoFromN(sp.States.CaSig.Ca.PSD, PSDVol), chem.CoFromN(sp.States.CaSig.CaM.PSD.CaM[2], PSDVol), chem.CoFromN(sp.States.CaSig.CaM.PSD.CaM[3], PSDVol), preSpike, &sp.Deltas.CaSig.Ca.PSD)

	sp.CaM.Step(&sp.States.CaSig.CaM, &sp.Deltas.CaSig.CaM, &sp.States.CaSig.Ca, &sp.Deltas.CaSig.Ca)

	sp.CaMKII.Step(&sp.States.CaSig.CaMKII, &sp.Deltas.CaSig.CaMKII, &sp.States.CaSig.CaM, &sp.Deltas.CaSig.CaM, &sp.States.CaSig.Ca, &sp.Deltas.CaSig.Ca, &sp.States.CaSig.PP1, &sp.Deltas.CaSig.PP1, sp.States.CaSig.PP2A, &sp.Deltas.CaSig.PP2A)

	sp.CaN.Step(&sp.States.CaSig.CaN, &sp.Deltas.CaSig.CaN, &sp.States.CaSig.CaM, &sp.Deltas.CaSig.CaM, &sp.States.CaSig.Ca, &sp.Deltas.CaSig.Ca)

	sp.PKA.Step(&sp.States.CaSig.PKA, &sp.Deltas.CaSig.PKA, &sp.States.CaSig.CaM, &sp.Deltas.CaSig.CaM)

	sp.PP1.Step(&sp.States.CaSig.PP1, &sp.Deltas.CaSig.PP1, &sp.States.CaSig.PKA, &sp.Deltas.CaSig.PKA, &sp.States.CaSig.CaN, &sp.Deltas.CaSig.CaN, sp.States.CaSig.PP2A, &sp.Deltas.CaSig.PP2A)

	sp.Ca.Step(&sp.States.CaSig.Ca, &sp.Deltas.CaSig.Ca)
	sp.AMPAR.Step(&sp.States.AMPAR, &sp.Deltas.AMPAR, &sp.States.CaSig, sp.States.CaSig.PP2A)
}

// Integrate integrates the deltas
func (sp *Spine) Integrate() {
	sp.States.Integrate(&sp.Deltas)
}

// StepTime steps and integrates for given amount of time in secs
func (sp *Spine) StepTime(secs float64) {
	for t := 0.0; t < secs; t += chem.IntegrationDt {
		sp.Step()
		sp.Integrate()
	}
}

func (sp *Spine) Stats(dir *tensorfs.Node) {
	sp.States.Stats(dir)
}
