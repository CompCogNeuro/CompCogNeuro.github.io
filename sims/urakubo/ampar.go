// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// AMPARVars have AMPAR Phosphorylation (Pd = phosphorylated, Dp = dephosphorylated) state.
// Two protein elements, separately Pd and Dp:
// AMPAR = AMPA receptor (GluR1), which can be Pd at Ser845 by PKA
// PDZs = PDZ domain binding proteins (e.g., SAP97 and stargazin), which bind to AMPAR
//
//	and are separately Pd by CaMKII -- denoted as StgP in Urakubo code
//
// Both can be Dp by PP1 and CaN (calcineurin)
// Variables named as P or D for the Pd or Dp state, 1st is AMPAR @ Ser845, 2nd is PDZs
type AMPARVars struct {

	// both dephosphorylated = Nophos
	DD float64

	// AMPA Ser845 phosphorylated by PKA = S845P
	PD float64

	// PDZs phosphorylated by CaMKII = StgP
	DP float64

	// both phosphorylated = S845PStgP
	PP float64

	// total of all phos levels
	Tot float64
}

func (as *AMPARVars) Init(uk *Urakubo) {
	as.Zero()
}

// Generate Code for Initializing
func (as *AMPARVars) InitCode(vol float64, pre string) {
	fmt.Printf("\tas.%s.DD = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(as.DD, vol))
	fmt.Printf("\tas.%s.PD = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(as.PD, vol))
	fmt.Printf("\tas.%s.DP = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(as.DP, vol))
	fmt.Printf("\tas.%s.PP = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(as.PP, vol))
}

func (as *AMPARVars) Zero() {
	as.DD = 0
	as.PD = 0
	as.DP = 0
	as.PP = 0
	as.Tot = 0
}

func (as *AMPARVars) Total() {
	as.Tot = as.DD + as.PD + as.DP + as.PP
}

func (as *AMPARVars) Integrate(d *AMPARVars) {
	chem.Integrate(&as.DD, d.DD)
	chem.Integrate(&as.PD, d.PD)
	chem.Integrate(&as.DP, d.DP)
	chem.Integrate(&as.PP, d.PP)
	as.Total()
}

func (as *AMPARVars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	dir.Float64(pre + "AMPAR").AppendRowFloat(chem.CoFromN(as.Tot, vol))
	// dir.Float64(pre+"DD").AppendRowFloat(chem.CoFromN(as.DD, vol))
	// dir.Float64(pre+"PD").AppendRowFloat(chem.CoFromN(as.PD, vol))
	// dir.Float64(pre+"DP").AppendRowFloat(chem.CoFromN(as.DP, vol))
	// dir.Float64(pre+"PP").AppendRowFloat(chem.CoFromN(as.PP, vol))
}

// AMPARState is AMPAR Phosphorylation and trafficking state.
// 4 Locations / states, which have their own time constants:
// Int = Cytosol, internal not integrated into membrane -- after endocyctosis
// Mbr = Cytosol, integrated into the membrane -- after exocytosis, still governed by Cyl rates
// PSD = In the postsynaptic density, but not trapped by scaffold
// Trp = Trapped by scaffolding in the PSD -- solidly fixed in place and active
// Trp.Tot is the net effective AMPA conductance
// 20 state vars total
type AMPARState struct {

	// cytosol internal
	Int AMPARVars `display:"inline"`

	// cytosol exocytosed into the membrane
	Mbr AMPARVars `display:"inline"`

	// in PSD but not trapped
	PSD AMPARVars `display:"inline"`

	// in PSD and trapped in place -- this is the effective weight
	Trp AMPARVars `display:"inline"`

	// amount of unbound scaffold used for trapping
	Scaffold float64
}

func (as *AMPARState) Init(uk *Urakubo) {
	as.Int.Init(uk)
	as.Mbr.Init(uk)
	as.PSD.Init(uk)
	as.Trp.Init(uk)

	as.Int.DD = chem.CoToN(3, CytVol) // Nophos_int
	as.Int.PD = chem.CoToN(3, CytVol) // S845P_int

	as.Trp.DD = chem.CoToN(1, PSDVol)
	as.Trp.PD = chem.CoToN(3, PSDVol)

	as.Scaffold = 0

	if uk.InitBaseline {

		as.Scaffold = chem.CoToN(2.234, PSDVol)
		vol := float64(CytVol)
		as.Int.DD = chem.CoToN(0.8335, vol)
		as.Int.PD = chem.CoToN(4.916, vol)
		as.Int.DP = chem.CoToN(0.05513, vol)
		as.Int.PP = chem.CoToN(0.3207, vol)
		as.Mbr.DD = chem.CoToN(0.002045, vol)
		as.Mbr.PD = chem.CoToN(0.3199, vol)
		as.Mbr.DP = chem.CoToN(0.0002025, vol)
		as.Mbr.PP = chem.CoToN(0.02452, vol)

		vol = PSDVol
		as.PSD.DD = chem.CoToN(0.0379, vol)
		as.PSD.PD = chem.CoToN(0.2744, vol)
		as.PSD.DP = chem.CoToN(0.004491, vol)
		as.PSD.PP = chem.CoToN(0.02951, vol)
		as.Trp.DD = chem.CoToN(0.2388, vol)
		as.Trp.PD = chem.CoToN(1.09, vol)
		as.Trp.DP = chem.CoToN(0.07655, vol)
		as.Trp.PP = chem.CoToN(0.305, vol)
	}

	as.Int.Total()
	as.Mbr.Total()
	as.PSD.Total()
	as.Trp.Total()
}

func (as *AMPARState) InitCode() {
	fmt.Printf("\nAMPARState:\n")
	as.Int.InitCode(CytVol, "Int")
	as.Mbr.InitCode(CytVol, "Mbr")
	as.PSD.InitCode(PSDVol, "PSD")
	as.Trp.InitCode(PSDVol, "Trp")
	fmt.Printf("\tas.Scaffold = chem.CoToN(%.4g, vol)\n", chem.CoFromN(as.Scaffold, PSDVol))
}

func (as *AMPARState) Zero() {
	as.Int.Zero()
	as.Mbr.Zero()
	as.PSD.Zero()
	as.Trp.Zero()
	as.Scaffold = 0
}

func (as *AMPARState) Integrate(d *AMPARState) {
	as.Int.Integrate(&d.Int)
	as.Mbr.Integrate(&d.Mbr)
	as.PSD.Integrate(&d.PSD)
	as.Trp.Integrate(&d.Trp)
	chem.Integrate(&as.Scaffold, d.Scaffold)
}

func (as *AMPARState) Stats(dir *tensorfs.Node) {
	as.Int.Stats(dir, CytVol, "Int_")
	as.Mbr.Stats(dir, CytVol, "Mbr_")
	as.PSD.Stats(dir, PSDVol, "PSD_")
	as.Trp.Stats(dir, PSDVol, "Trp_")
}

// AMPAR phosphorylation and trafficking parameters
// Original kinetic rate constants are in units of (μM-1s-1),
type AMPARPhosParams struct {

	// rate of phosphorylation of AMPA Ser845 by PKA
	PKA chem.SimpleEnz

	// rate of phosphorylation of PDZs by CaMKII
	CaMKII chem.SimpleEnz

	// rate of dephosphorylation of AMPA Ser845 by PP1
	PP_S845 chem.SimpleEnz

	// rate of dephosphorylation of PDZs by PP1
	PP_PDZs chem.SimpleEnz

	// rate of dephosphorylation of AMPA Ser845 by CaN
	CaN_S845 chem.SimpleEnz

	// rate of dephosphorylation of PDZs by CaN
	CaN_PDZs chem.SimpleEnz

	// rate of dephosphorylation of AMPA Ser845 by PP2A
	PP2A_S845 chem.SimpleEnz

	// rate of dephosphorylation of PDZs by PP2A
	PP2A_PDZs chem.SimpleEnz
}

func (ap *AMPARPhosParams) Defaults() {
	ap.PKA.Kf = 20
	ap.CaMKII.Kf = 1
	ap.PP_S845.Kf = 4
	ap.PP_PDZs.Kf = 100
	ap.CaN_S845.Kf = 1.5
	ap.CaN_PDZs.Kf = 1
	ap.PP2A_S845.Kf = 4
	ap.PP2A_PDZs.Kf = 100
}

// StepP updates the phosphorylation d=delta state from c=current
// based on current kinase / pp states
func (ap *AMPARPhosParams) StepP(c, d *AMPARVars, vol, camkii, can, pka, pp1 float64) {
	ap.PKA.StepCo(c.DD, pka, vol, &d.DD, &d.PD)
	ap.PKA.StepCo(c.DP, pka, vol, &d.DP, &d.PP)
	ap.CaMKII.StepCo(c.DD, camkii, vol, &d.DD, &d.DP)
	ap.CaMKII.StepCo(c.PD, camkii, vol, &d.PD, &d.PP)

	ap.PP_S845.StepCo(c.PD, pp1, vol, &d.PD, &d.DD)
	ap.PP_S845.StepCo(c.PP, pp1, vol, &d.PP, &d.DP)
	ap.PP_PDZs.StepCo(c.DP, pp1, vol, &d.DP, &d.DD)
	ap.PP_PDZs.StepCo(c.PP, pp1, vol, &d.PP, &d.PD)

	ap.CaN_S845.StepCo(c.PD, can, vol, &d.PD, &d.DD)
	ap.CaN_S845.StepCo(c.PP, can, vol, &d.PP, &d.DP)
	ap.CaN_PDZs.StepCo(c.DP, can, vol, &d.DP, &d.DD)
	ap.CaN_PDZs.StepCo(c.PP, can, vol, &d.PP, &d.PD)
}

// StepPP2A updates the phosphorylation n=next state from c=current
// based on current pp2a
func (ap *AMPARPhosParams) StepPP2A(c, d *AMPARVars, vol, pp2a float64) {
	ap.PP2A_S845.StepCo(c.PD, pp2a, vol, &d.PD, &d.DD)
	ap.PP2A_S845.StepCo(c.PP, pp2a, vol, &d.PP, &d.DP)
	ap.PP2A_PDZs.StepCo(c.DP, pp2a, vol, &d.DP, &d.DD)
	ap.PP2A_PDZs.StepCo(c.PP, pp2a, vol, &d.PP, &d.PD)
}

// AMPAR trafficking parameters
// Original kinetic rate constants are in units of (μM-1s-1),
type AMPARTrafParams struct {

	// Ser845P endocytosis, excocytosis rates -- Kf = 30min, Kb = 9min
	EndoExoP chem.React

	// Ser845D endocytosis rate -- Kf = 1sec, Kb = 0
	EndoD chem.React

	// PDZsP trapping in the PSD -- faster when P -- Kf is PSD + Scaffold -> Trp, Kb reverse
	TrapP chem.React

	// PDZsD trapping in the PSD -- slower when D -- Kf is PSD + Scaffold -> Trp, Kb reverse
	TrapD chem.React

	// diffusion for each category, all have the same constant
	Diffuse chem.Diffuse
}

func (ap *AMPARTrafParams) Defaults() {
	ap.EndoExoP.Set(1.0/(9*60), 1.0/(30*60))
	ap.EndoD.Set(1, 0)
	ap.TrapP.Set(0.041667, 0.033333)
	ap.TrapD.Set(0.0025, 0.033333)
	ap.Diffuse.SetSym(1.6)
}

// StepT computes trafficking deltas
func (ap *AMPARTrafParams) StepT(c, d *AMPARState) {

	var dummy float64
	// Exo = Int -> Mbr
	ap.EndoExoP.Step(c.Mbr.PD, 1, c.Int.PD, &d.Mbr.PD, &dummy, &d.Int.PD)
	ap.EndoExoP.Step(c.Mbr.PP, 1, c.Int.PP, &d.Mbr.PP, &dummy, &d.Int.PP)

	ap.EndoD.Step(c.Mbr.DD, 1, c.Int.DD, &d.Mbr.DD, &dummy, &d.Int.DD)
	ap.EndoD.Step(c.Mbr.DP, 1, c.Int.DP, &d.Mbr.DP, &dummy, &d.Int.DP)

	ap.TrapP.Step(c.PSD.DP, c.Scaffold, c.Trp.DP, &d.PSD.DP, &d.Scaffold, &d.Trp.DP)
	ap.TrapP.Step(c.PSD.PP, c.Scaffold, c.Trp.PP, &d.PSD.PP, &d.Scaffold, &d.Trp.PP)

	ap.TrapD.Step(c.PSD.DD, c.Scaffold, c.Trp.DD, &d.PSD.DD, &d.Scaffold, &d.Trp.DD)
	ap.TrapD.Step(c.PSD.PD, c.Scaffold, c.Trp.PD, &d.PSD.PD, &d.Scaffold, &d.Trp.PD)

	// Diffuse = Mbr -> PSD
	ap.Diffuse.Step(c.Mbr.DD, c.PSD.DD, CytVol, PSDVol, &d.Mbr.DD, &d.PSD.DD)
	ap.Diffuse.Step(c.Mbr.PD, c.PSD.PD, CytVol, PSDVol, &d.Mbr.PD, &d.PSD.PD)
	ap.Diffuse.Step(c.Mbr.DP, c.PSD.DP, CytVol, PSDVol, &d.Mbr.DP, &d.PSD.DP)
	ap.Diffuse.Step(c.Mbr.PP, c.PSD.PP, CytVol, PSDVol, &d.Mbr.PP, &d.PSD.PP)
}

// AMPAR phosphorylation and trafficking parameters
// Original kinetic rate constants are in units of (μM-1s-1),
// converted to msec instead of sec
type AMPARParams struct {

	// Phosphorylation parameters
	Phos AMPARPhosParams `display:"inline"`

	// Trafficking parameters
	Traffic AMPARTrafParams `display:"inline"`
}

func (ap *AMPARParams) Defaults() {
	ap.Phos.Defaults()
	ap.Traffic.Defaults()
}

// Step does full AMPAR updating, c=current, n=next
// based on current Ca signaling state
func (ap *AMPARParams) Step(c, d *AMPARState, cas *CaSigState, pp2a float64) {
	ap.Phos.StepP(&c.Int, &d.Int, CytVol, cas.CaMKII.Cyt.Auto.Act, cas.CaN.Cyt.CaNact, cas.PKA.Cyt.PKAact, cas.PP1.Cyt.PP1act)
	ap.Phos.StepP(&c.Mbr, &d.Mbr, CytVol, cas.CaMKII.Cyt.Auto.Act, cas.CaN.Cyt.CaNact, cas.PKA.Cyt.PKAact, cas.PP1.Cyt.PP1act)
	ap.Phos.StepP(&c.Trp, &d.Trp, PSDVol, cas.CaMKII.PSD.Auto.Act, cas.CaN.PSD.CaNact, cas.PKA.PSD.PKAact, cas.PP1.PSD.PP1act)
	ap.Phos.StepP(&c.PSD, &d.PSD, PSDVol, cas.CaMKII.PSD.Auto.Act, cas.CaN.PSD.CaNact, cas.PKA.PSD.PKAact, cas.PP1.PSD.PP1act)

	ap.Phos.StepPP2A(&c.Int, &d.Int, CytVol, pp2a) // Cyt only
	ap.Phos.StepPP2A(&c.Mbr, &d.Mbr, CytVol, pp2a) // Cyt only

	ap.Traffic.StepT(c, d)
}
