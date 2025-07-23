// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// CaMVars are intracellular Ca-driven signaling states for CaM binding with Ca
// From Urakubo: Ca2+ binding kinetics of CaM has extensively been analyzed
// (Linse et al., 1991; Holmes, 2000). CaM binds to four Ca2+ ions,
// but two or three Ca2+-binding is enough to activate CaM
// (James et al., 1995; Chin and Means, 2000).
// For simplicity, 3Ca2+⋅CaM is assumed to be an active form,
// and reactions for 4Ca2+⋅CaM are omitted.
type CaMVars struct {
	// increasing levels of Ca binding to CaM, 0-3, [3] is active form
	CaM [4]float64
}

func (cs *CaMVars) Init(uk *Urakubo, vol float64) {
	for i := range cs.CaM {
		cs.CaM[i] = 0
	}
	cs.CaM[0] = chem.CoToN(80, vol)

	if uk.InitBaseline {
		cs.CaM[0] = chem.CoToN(78.31, vol)
		cs.CaM[1] = chem.CoToN(1.002, vol)
		cs.CaM[2] = chem.CoToN(0.006681, vol)
		cs.CaM[3] = chem.CoToN(1.988e-05, vol)
	}
}

func (cs *CaMVars) Active() float64 {
	return cs.CaM[3]
}

// Generate Code for Initializing
func (cs *CaMVars) InitCode(vol float64, pre string) {
	for i := range cs.CaM {
		fmt.Printf("\tcs.%s.CaM[%d] = chem.CoToN(%.4g, vol)\n", pre, i, chem.CoFromN(cs.CaM[i], vol))
	}
}

func (cs *CaMVars) Zero() {
	for i := range cs.CaM {
		cs.CaM[i] = 0
	}
}

func (cs *CaMVars) Integrate(d *CaMVars) {
	for i := range cs.CaM {
		chem.Integrate(&cs.CaM[i], d.CaM[i])
	}
}

func (cs *CaMVars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	// dir.Float64(pre+"CaM").AppendRowFloat(chem.CoFromN(cs.CaM[0], vol))
	dir.Float64(pre + "CaMact").AppendRowFloat(chem.CoFromN(cs.Active(), vol))
	// dir.Float64(pre+"CaCaM").AppendRowFloat(chem.CoFromN(cs.Ca[1], vol))
	// dir.Float64(pre+"Ca2CaM").AppendRowFloat(chem.CoFromN(cs.Ca[2], vol))
}

// CaMState is overall intracellular Ca-driven signaling states
// for CaM in Cyt and PSD
// 32 state vars total
type CaMState struct {

	// in cytosol -- volume = 0.08 fl = 48
	Cyt CaMVars

	// in PSD -- volume = 0.02 fl = 12
	PSD CaMVars
}

func (cs *CaMState) Init(uk *Urakubo) {
	cs.Cyt.Init(uk, CytVol)
	cs.PSD.Init(uk, PSDVol)

	if uk.InitBaseline {
		vol := float64(PSDVol)
		cs.PSD.CaM[3] = chem.CoToN(2.738e-05, vol)
	}
}

func (cs *CaMState) InitCode() {
	fmt.Printf("\nCaMState:\n")
	cs.Cyt.InitCode(CytVol, "Cyt")
	cs.PSD.InitCode(PSDVol, "PSD")
}

func (cs *CaMState) Zero() {
	cs.Cyt.Zero()
	cs.PSD.Zero()
}

func (cs *CaMState) Integrate(d *CaMState) {
	cs.Cyt.Integrate(&d.Cyt)
	cs.PSD.Integrate(&d.PSD)
}

func (cs *CaMState) Stats(dir *tensorfs.Node) {
	cs.Cyt.Stats(dir, CytVol, "Cyt_")
	cs.PSD.Stats(dir, PSDVol, "PSD_")
}

// CaMParams are the parameters governing the Ca+CaM binding
type CaMParams struct {
	// 1: Ca+CaM -> 1CaCaM = CaM-bind-Ca
	CaCaM01 chem.React `desc:"1: Ca+CaM -> 1CaCaM = CaM-bind-Ca"`
	// 2: Ca+1CaM -> 2CaCaM = CaMCa-bind-Ca
	CaCaM12 chem.React `desc:"2: Ca+1CaM -> 2CaCaM = CaMCa-bind-Ca"`
	// 3: Ca+2CaM -> 3CaCaM = CaMCa2-bind-Ca
	CaCaM23 chem.React `desc:"3: Ca+2CaM -> 3CaCaM = CaMCa2-bind-Ca"`
	// CaM diffusion between Cyt and PSD
	CaMDiffuse chem.Diffuse `desc:"CaM diffusion between Cyt and PSD"`
}

func (cp *CaMParams) Defaults() {
	// note: following are all in Cyt -- PSD is 4x for first values
	// See React docs for more info
	cp.CaCaM01.SetVol(51.202, CytVol, 200) // 1: 51.202 μM-1 = 1.0667, PSD 4.2667 = CaM-bind-Ca
	cp.CaCaM12.SetVol(133.3, CytVol, 1000) // 2: 133.3 μM-1 = 2.7771, PSD 11.108 = CaMCa-bind-Ca
	cp.CaCaM23.SetVol(25.6, CytVol, 400)   // 3: 25.6 μM-1 = 0.53333, PSD 2.1333 = CaMCa2-bind-Ca
	cp.CaMDiffuse.SetSym(130.0 / 0.0225)
}

// StepCaM does the bulk of Ca + CaM + CaM binding reactions, in a given region
// cCa, nCa = current next Ca
func (cp *CaMParams) StepCaM(vol float64, c, d *CaMVars, cCa float64, dCa *float64) {
	kf := CytVol / vol
	cp.CaCaM01.StepK(kf, c.CaM[0], cCa, c.CaM[1], &d.CaM[0], dCa, &d.CaM[1]) // 1
	cp.CaCaM12.StepK(kf, c.CaM[1], cCa, c.CaM[2], &d.CaM[1], dCa, &d.CaM[2]) // 2
	cp.CaCaM23.StepK(kf, c.CaM[2], cCa, c.CaM[3], &d.CaM[2], dCa, &d.CaM[3]) // 3
}

// StepDiffuse does Cyt <-> PSD diffusion
func (cp *CaMParams) StepDiffuse(c, d *CaMState) {
	for i := 0; i < 4; i++ {
		cc := c.Cyt.CaM[i]
		cd := c.PSD.CaM[i]
		dc := &d.Cyt.CaM[i]
		dd := &d.PSD.CaM[i]
		cp.CaMDiffuse.Step(cc, cd, CytVol, PSDVol, dc, dd)
	}
}

// Step does one step of CaM updating, c=current, d=delta
// pp2a = current cyt pp2a
func (cp *CaMParams) Step(c, d *CaMState, cCa, dCa *CaState) {
	cp.StepCaM(CytVol, &c.Cyt, &d.Cyt, cCa.Cyt, &dCa.Cyt)
	cp.StepCaM(PSDVol, &c.PSD, &d.PSD, cCa.PSD, &dCa.PSD)
	cp.StepDiffuse(c, d)
}
