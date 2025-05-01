// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// PP1Vars are intracellular Ca-driven signaling variables for the
// PP1 - I-1 system
type PP1Vars struct {

	// dephosphorylated I-1 = I1_inactive
	I1 float64

	// phosphorylated (active) I-1 = I1_active
	I1P float64

	// PP1 = protein phosphatase 1 bound with I-1P
	PP1_I1P float64

	// activated PP1
	PP1act float64

	// PKA+I1 complex for PKAI1 enzyme reaction
	PKAI1C float64

	// CaN+I1P complex for CaNI1P enzyme reaction
	CaNI1PC float64

	// PP2A+I1P complex for PP2AI1P enzyme reaction
	PP2AI1PC float64
}

func (ps *PP1Vars) Init(uk *Urakubo, vol float64) {
	ps.I1 = chem.CoToN(2, vol)
	ps.I1P = 0
	ps.PP1_I1P = chem.CoToN(2, vol)
	ps.PP1act = 0
	ps.PKAI1C = 0
	ps.CaNI1PC = 0
	ps.PP2AI1PC = 0

	if uk.InitBaseline {
		// All vals below from 500 sec baseline
		ps.I1 = chem.CoToN(0.8722, vol)
		ps.I1P = chem.CoToN(1.131, vol)
		ps.CaNI1PC = chem.CoToN(0.002632, vol)
	}
}

func (ps *PP1Vars) InitCode(vol float64, pre string) {
	fmt.Printf("\tps.%s.I1 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.I1, vol))
	fmt.Printf("\tps.%s.I1P = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.I1P, vol))
	fmt.Printf("\tps.%s.PP1_I1P = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.PP1_I1P, vol))
	fmt.Printf("\tps.%s.PP1act = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.PP1act, vol))
	fmt.Printf("\tps.%s.CaNI1PC = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.CaNI1PC, vol))
	fmt.Printf("\tps.%s.PP2AI1PC = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.PP2AI1PC, vol))
}

func (ps *PP1Vars) Zero() {
	ps.I1 = 0
	ps.I1P = 0
	ps.PP1_I1P = 0
	ps.PP1act = 0
	ps.PKAI1C = 0
	ps.CaNI1PC = 0
	ps.PP2AI1PC = 0
}

func (ps *PP1Vars) Integrate(d *PP1Vars) {
	chem.Integrate(&ps.I1, d.I1)
	chem.Integrate(&ps.I1P, d.I1P)
	chem.Integrate(&ps.PP1_I1P, d.PP1_I1P)
	chem.Integrate(&ps.PP1act, d.PP1act)
	chem.Integrate(&ps.PKAI1C, d.PKAI1C)
	chem.Integrate(&ps.CaNI1PC, d.CaNI1PC)
	chem.Integrate(&ps.PP2AI1PC, d.PP2AI1PC)
}

func (ps *PP1Vars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	// dir.Float64(pre+"I1").AppendRowFloat(chem.CoFromN(ps.I1, vol))
	dir.Float64(pre + "I1P").AppendRowFloat(chem.CoFromN(ps.I1P, vol))
	dir.Float64(pre + "PP1act").AppendRowFloat(chem.CoFromN(ps.PP1act, vol))
}

// PP1State is overall intracellular Ca-driven signaling states
// for PP1-I-1 in Cyt and PSD
// 14 state vars total
type PP1State struct {

	// in cytosol -- volume = 0.08 fl
	Cyt PP1Vars

	// in PSD -- volume = 0.02 fl
	PSD PP1Vars
}

func (ps *PP1State) Init(uk *Urakubo) {
	ps.Cyt.Init(uk, CytVol)
	ps.PSD.Init(uk, PSDVol)

	if uk.InitBaseline {
		vol := float64(CytVol)
		ps.Cyt.PP1_I1P = chem.CoToN(0.9909, vol)
		ps.Cyt.PP1act = chem.CoToN(0.008759, vol)
		ps.Cyt.PP2AI1PC = chem.CoToN(0.008215, vol)

		vol = PSDVol
		ps.PSD.PP1_I1P = chem.CoToN(5.949, vol)
		ps.PSD.PP1act = chem.CoToN(0.05258, vol)
		ps.PSD.PP2AI1PC = chem.CoToN(0, vol)
	}
}

func (ps *PP1State) InitCode() {
	fmt.Printf("\nPP1State:\n")
	ps.Cyt.InitCode(CytVol, "Cyt")
	ps.PSD.InitCode(PSDVol, "PSD")
}

func (ps *PP1State) Zero() {
	ps.Cyt.Zero()
	ps.PSD.Zero()
}

func (ps *PP1State) Integrate(d *PP1State) {
	ps.Cyt.Integrate(&d.Cyt)
	ps.PSD.Integrate(&d.PSD)
}

func (ps *PP1State) Stats(dir *tensorfs.Node) {
	ps.Cyt.Stats(dir, CytVol, "Cyt_")
	ps.PSD.Stats(dir, PSDVol, "PSD_")
}

// PP1Params are the parameters governing the PP1-I-1 binding
type PP1Params struct {

	// 1: I-1P + PP1act -> PP1-I1P -- Table SIi constants are backward = I1-PP1
	I1PP1 chem.React

	// 2: I-1P phosphorylated by PKA -- Table SIj numbers != Figure SI4
	PKAI1 chem.Enz

	// 3: I-1P dephosphorylated by CaN -- Table SIj number
	CaNI1P chem.Enz

	// 4: I-1P dephosphorylated by PP2A -- Table SIj number
	PP2aI1P chem.Enz

	// I1 diffusion between Cyt and PSD
	I1Diffuse chem.Diffuse

	// PP1 diffusion between Cyt and PSD
	PP1Diffuse chem.Diffuse
}

func (cp *PP1Params) Defaults() {
	// note: following are all in Cyt -- PSD is 4x for first values
	cp.I1PP1.SetVol(100, CytVol, 1)           // Kb = 100 μM-1 = 2.0834 -- reversed for product = PP1-I1P
	cp.PKAI1.SetKmVol(8.1, CytVol, 21.2, 5.3) // Km = 8.1 μM-1 k1 = 0.068157
	cp.CaNI1P.SetKmVol(3, CytVol, 11.2, 2.8)  // Km = 3 μM-1 = 0.097222
	cp.PP2aI1P.SetKmVol(3, CytVol, 11.2, 2.8) // Km = 3 μM-1 = 0.097222
	cp.I1Diffuse.SetSym(35.9 / 0.0225)
	cp.PP1Diffuse.Set(31.4/0.0225, 5.23/0.0225)
}

// StepPP1 does the bulk of Ca + PP1 + CaM binding reactions, in a given region
// cCaM, nCaM = current, new 3CaCaM from CaMKIIVars
// cCa, nCa = current new Ca
func (cp *PP1Params) StepPP1(vol float64, c, d *PP1Vars, pka, can, pp2a float64, dpka, dcan, dpp2a *float64) {
	kf := CytVol / vol
	cp.I1PP1.StepK(kf, c.I1P, c.PP1act, c.PP1_I1P, &d.I1P, &d.PP1act, &d.PP1_I1P) // 1

	// cs, ce, cc, cp -> ds, de, dc, dp
	cp.PKAI1.StepK(kf, c.I1, pka, c.PKAI1C, c.I1P, &d.I1, dpka, &d.PKAI1C, &d.I1P)    // 2
	cp.CaNI1P.StepK(kf, c.I1P, can, c.CaNI1PC, c.I1, &d.I1P, dcan, &d.CaNI1PC, &d.I1) // 3
	if dpp2a != nil {                                                                 // no PP2A in PSD
		cp.PP2aI1P.StepK(kf, c.I1P, pp2a, c.PP2AI1PC, c.I1, &d.I1P, dpp2a, &d.PP2AI1PC, &d.I1) // 3
	}
}

// StepDiffuse does diffusion update, c=current, d=delta
func (cp *PP1Params) StepDiffuse(c, d *PP1State) {
	cp.I1Diffuse.Step(c.Cyt.I1, c.PSD.I1, CytVol, PSDVol, &d.Cyt.I1, &d.PSD.I1)
	cp.I1Diffuse.Step(c.Cyt.I1P, c.PSD.I1P, CytVol, PSDVol, &d.Cyt.I1P, &d.PSD.I1P)
	cp.PP1Diffuse.Step(c.Cyt.PP1_I1P, c.PSD.PP1_I1P, CytVol, PSDVol, &d.Cyt.PP1_I1P, &d.PSD.PP1_I1P)
	cp.PP1Diffuse.Step(c.Cyt.PP1act, c.PSD.PP1act, CytVol, PSDVol, &d.Cyt.PP1act, &d.PSD.PP1act)
}

// Step does full CaN updating, c=current, d=delta
func (cp *PP1Params) Step(c, d *PP1State, pka, dpka *PKAState, can, dcan *CaNState, pp2a float64, dpp2a *float64) {
	cp.StepPP1(CytVol, &c.Cyt, &d.Cyt, pka.Cyt.PKAact, can.Cyt.CaNact, pp2a, &dpka.Cyt.PKAact, &dcan.Cyt.CaNact, dpp2a)
	cp.StepPP1(PSDVol, &c.PSD, &d.PSD, pka.PSD.PKAact, can.PSD.CaNact, 0, &dpka.PSD.PKAact, &dcan.PSD.CaNact, nil)
	cp.StepDiffuse(c, d)
}
