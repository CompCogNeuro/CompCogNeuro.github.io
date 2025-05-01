// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// CaNVars are intracellular Ca-driven signaling variables for the
// CaN and CaM binding, at different levels of Ca binding
// stores N values -- Co = Concentration computed by volume as needed
type CaNVars struct {

	// Calcineurin
	CaN float64

	// CaN-CaM bound
	CaNCaM float64
}

func (cs *CaNVars) Init(uk *Urakubo, vol float64) {
	cs.CaN = 0
	cs.CaNCaM = 0
}

func (cs *CaNVars) Zero() {
	cs.CaN = 0
	cs.CaNCaM = 0
}

func (cs *CaNVars) Integrate(d *CaNVars) {
	chem.Integrate(&cs.CaN, d.CaN)
	chem.Integrate(&cs.CaNCaM, d.CaNCaM)
}

// CaNCaMVars are intracellular Ca-driven signaling states
// for CaN-CaM binding
// stores N values -- Co = Concentration computed by volume as needed
type CaNCaMVars struct {

	// increasing levels of Ca binding, 0-2
	Ca [3]CaNVars

	// active CaN = Ca[2].CaNCaM
	CaNact float64
}

func (cs *CaNCaMVars) Init(uk *Urakubo, vol float64) {
	for i := range cs.Ca {
		cs.Ca[i].Init(uk, vol)
	}
	// original
	cs.Ca[0].CaN = chem.CoToN(3, vol)
	cs.CaNact = 0

	if uk.InitBaseline {
		cs.Ca[0].CaN = chem.CoToN(1.305, vol)
		cs.Ca[0].CaNCaM = chem.CoToN(0.0279, vol)
		cs.Ca[1].CaN = chem.CoToN(1.305, vol)
		cs.Ca[1].CaNCaM = chem.CoToN(0.0279, vol)
		cs.Ca[2].CaN = chem.CoToN(0.3263, vol)
		cs.Ca[2].CaNCaM = chem.CoToN(0.006977, vol)
		cs.CaNact = chem.CoToN(0.006977, vol)
	}
}

// Generate Code for Initializing
func (cs *CaNCaMVars) InitCode(vol float64, pre string) {
	for i := range cs.Ca {
		fmt.Printf("\tcs.%s.Ca[%d].CaN = chem.CoToN(%.4g, vol)\n", pre, i, chem.CoFromN(cs.Ca[i].CaN, vol))
		fmt.Printf("\tcs.%s.Ca[%d].CaNCaM = chem.CoToN(%.4g, vol)\n", pre, i, chem.CoFromN(cs.Ca[i].CaNCaM, vol))
	}
	fmt.Printf("\tcs.%s.CaNact = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(cs.CaNact, vol))
}

func (cs *CaNCaMVars) Zero() {
	for i := range cs.Ca {
		cs.Ca[i].Zero()
	}
	cs.CaNact = 0
}

func (cs *CaNCaMVars) Integrate(d *CaNCaMVars) {
	for i := range cs.Ca {
		cs.Ca[i].Integrate(&d.Ca[i])
	}
	cs.CaNact = cs.Ca[2].CaNCaM
}

func (cs *CaNCaMVars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	dir.Float64(pre + "CaNact").AppendRowFloat(chem.CoFromN(cs.CaNact, vol))
}

// CaNState is overall intracellular Ca-driven signaling states
// for CaN-CaM binding in Cyt and PSD
// 14 state vars total
type CaNState struct {

	// in cytosol -- volume = 0.08 fl = 48
	Cyt CaNCaMVars

	// in PSD  -- volume = 0.02 fl = 12
	PSD CaNCaMVars
}

func (cs *CaNState) Init(uk *Urakubo) {
	cs.Cyt.Init(uk, CytVol)
	cs.PSD.Init(uk, PSDVol)
}

func (cs *CaNState) InitCode() {
	fmt.Printf("\nCaNState:\n")
	cs.Cyt.InitCode(CytVol, "Cyt")
	cs.PSD.InitCode(PSDVol, "PSD")
}

func (cs *CaNState) Zero() {
	cs.Cyt.Zero()
	cs.PSD.Zero()
}

func (cs *CaNState) Integrate(d *CaNState) {
	cs.Cyt.Integrate(&d.Cyt)
	cs.PSD.Integrate(&d.PSD)
}

func (cs *CaNState) Stats(dir *tensorfs.Node) {
	cs.Cyt.Stats(dir, CytVol, "Cyt_")
	cs.PSD.Stats(dir, PSDVol, "PSD_")
}

// CaNParams are the parameters governing the Ca+CaN-CaM binding
type CaNParams struct {

	// 1: CaN+CaM -> CaN-CaM
	CaNCaM chem.React

	// 2: Ca+CaM -> CaCaM
	CaCaN01 chem.React

	// 3: Ca+CaCaM -> 2CaCaM
	CaCaN12 chem.React

	// CaN diffusion between Cyt and PSD
	CaNDiffuse chem.Diffuse
}

func (cp *CaNParams) Defaults() {
	// note: following are all in Cyt -- PSD is 4x for first values
	// See React docs for more info
	cp.CaNCaM.SetVol(40, CytVol, 0.04) // 1: 40 μM-1 = 0.83333, PSD = 3.3333
	cp.CaCaN01.SetVol(20, CytVol, 1.0) // 2: 20 μM-1 = 0.41667, PSD = 1.6667
	cp.CaCaN12.SetVol(10, CytVol, 2.0) // 3: 10 μM-1 = 0.20833, PSD = 0.83333
	cp.CaNDiffuse.SetSym(20.0 / 0.0225)
}

// StepCaN does the bulk of Ca + CaN + CaM binding reactions, in a given region
// cCaM, dCaM = current, delta 3CaCaM from CaMKIIVars
// cCa, dCa = current, delta Ca
func (cp *CaNParams) StepCaN(vol float64, c, d *CaNCaMVars, cCa, cCaM float64, dCa, dCaM *float64) {
	kf := CytVol / vol
	for i := 0; i < 3; i++ {
		cp.CaNCaM.StepK(kf, c.Ca[i].CaN, cCaM, c.Ca[i].CaNCaM, &d.Ca[i].CaN, dCaM, &d.Ca[i].CaNCaM) // 1
	}
	cp.CaCaN01.StepK(kf, c.Ca[0].CaN, cCa, c.Ca[1].CaN, &d.Ca[0].CaN, dCa, &d.Ca[1].CaN)             // 2
	cp.CaCaN01.StepK(kf, c.Ca[0].CaNCaM, cCa, c.Ca[1].CaNCaM, &d.Ca[0].CaNCaM, dCa, &d.Ca[1].CaNCaM) // 2

	cp.CaCaN12.StepK(kf, c.Ca[1].CaN, cCa, c.Ca[2].CaN, &d.Ca[1].CaN, dCa, &d.Ca[2].CaN)             // 3
	cp.CaCaN12.StepK(kf, c.Ca[1].CaNCaM, cCa, c.Ca[2].CaNCaM, &d.Ca[1].CaNCaM, dCa, &d.Ca[2].CaNCaM) // 3
}

// StepDiffuse does diffusion update, c=current, d=delta
func (cp *CaNParams) StepDiffuse(c, d *CaNState) {
	for i := 0; i < 3; i++ {
		cp.CaNDiffuse.Step(c.Cyt.Ca[i].CaN, c.PSD.Ca[i].CaN, CytVol, PSDVol, &d.Cyt.Ca[i].CaN, &d.PSD.Ca[i].CaN)
		cp.CaNDiffuse.Step(c.Cyt.Ca[i].CaNCaM, c.PSD.Ca[i].CaNCaM, CytVol, PSDVol, &d.Cyt.Ca[i].CaNCaM, &d.PSD.Ca[i].CaNCaM)
	}
}

// Step does full CaN updating, c=current, d=delta
func (cp *CaNParams) Step(c, d *CaNState, cCaM, dCaM *CaMState, cCa, dCa *CaState) {
	cp.StepCaN(CytVol, &c.Cyt, &d.Cyt, cCa.Cyt, cCaM.Cyt.CaM[3], &dCa.Cyt, &dCaM.Cyt.CaM[3])
	cp.StepCaN(PSDVol, &c.PSD, &d.PSD, cCa.PSD, cCaM.PSD.CaM[3], &dCa.PSD, &dCaM.PSD.CaM[3])
	cp.StepDiffuse(c, d)
}
