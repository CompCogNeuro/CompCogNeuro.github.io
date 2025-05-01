// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// PKAVars are intracellular Ca-driven signaling states
// for PKA binding and phosphorylation with cAMP
// stores N values -- Co = Concentration computed by volume as needed
type PKAVars struct {

	// AC1
	AC1 float64

	// active AC1 = CaM-AC1
	AC1act float64

	// active PDE = cAMP-PDE -- buffered to 1
	PDEact float64

	// ATP -- buffered to 10000
	ATP float64

	// cAMP
	CAMP float64

	// AMP -- buffered to 1000
	AMP float64

	// R2C2
	R2C2 float64

	// R2C2-cAMP_B
	R2C2_B float64

	// R2C2-2cAMP-B-B
	R2C2_BB float64

	// R2C2-2cAMP-A-B
	R2C2_AB float64

	// R2C2-3cAMP-A-B-B
	R2C2_ABB float64

	// R2C2-4cAMP
	R2C2_4 float64

	// R2C-3cAMP -- note Fig SI4 mislabeled as R2-3
	R2C_3 float64

	// R2C-4cAMP
	R2C_4 float64

	// R2-3cAMP
	R2_3 float64

	// R2-4cAMP
	R2_4 float64

	// active PKA
	PKAact float64

	// AC1act+ATP complex for AC1ATP enzyme reaction -- reflects rate
	AC1ATPC float64

	// PDEact+cAMP complex for PDEcAMP enzyme reaction
	PDEcAMPC float64
}

func (ps *PKAVars) Init(uk *Urakubo, vol float64) {
	ps.AC1 = chem.CoToN(2, vol)
	ps.AC1act = 0
	ps.PDEact = chem.CoToN(1, vol)  // buffered!
	ps.ATP = chem.CoToN(10000, vol) // buffered! -- note: large #'s here contribute significantly to instability
	// todo: experiment with significantly smaller #'s
	ps.CAMP = 0
	ps.AMP = chem.CoToN(1000, vol) // buffered!
	ps.R2C2 = chem.CoToN(2, vol)
	ps.R2C2_B = 0
	ps.R2C2_BB = 0
	ps.R2C2_AB = 0
	ps.R2C2_ABB = 0
	ps.R2C2_4 = 0
	ps.R2C_3 = 0
	ps.R2C_4 = 0
	ps.R2_3 = 0
	ps.R2_4 = 0
	ps.PKAact = chem.CoToN(0.05, vol)
	ps.AC1ATPC = chem.CoToN(0.00025355, vol)
	ps.PDEcAMPC = 0

	if uk.InitBaseline {
		ps.AC1act = chem.CoToN(0.0002385, vol)
		ps.CAMP = chem.CoToN(0.004477, vol)
		ps.R2C2 = chem.CoToN(1.982, vol)
		ps.R2C2_B = chem.CoToN(0.01775, vol)
		ps.R2C2_BB = chem.CoToN(3.973e-05, vol)
		ps.R2C2_AB = chem.CoToN(3.179e-05, vol)
		ps.R2C2_ABB = chem.CoToN(1.423e-07, vol)
		ps.R2C2_4 = chem.CoToN(1.274e-10, vol)

		ps.R2C_3 = chem.CoToN(6.304e-07, vol)
		ps.R2C_4 = chem.CoToN(5.645e-08, vol)
		ps.R2_3 = chem.CoToN(6.981e-07, vol)
		ps.R2_4 = chem.CoToN(6.251e-06, vol)
		ps.PKAact = chem.CoToN(0.04515, vol)
		ps.AC1ATPC = chem.CoToN(2.335e+05, vol)
		ps.PDEcAMPC = chem.CoToN(0.0004477, vol)
	}
}

func (ps *PKAVars) InitCode(vol float64, pre string) {
	fmt.Printf("\tps.%s.AC1 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.AC1, vol))
	fmt.Printf("\tps.%s.AC1act = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.AC1act, vol))
	fmt.Printf("\tps.%s.CAMP = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.CAMP, vol))
	fmt.Printf("\tps.%s.R2C2 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2, vol))
	fmt.Printf("\tps.%s.R2C2_B = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2_B, vol))
	fmt.Printf("\tps.%s.R2C2_BB = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2_BB, vol))
	fmt.Printf("\tps.%s.R2C2_AB = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2_AB, vol))
	fmt.Printf("\tps.%s.R2C2_ABB = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2_ABB, vol))
	fmt.Printf("\tps.%s.R2C2_4 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C2_4, vol))
	fmt.Printf("\tps.%s.R2C_3 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C_3, vol))
	fmt.Printf("\tps.%s.R2C_4 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2C_4, vol))
	fmt.Printf("\tps.%s.R2_3 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2_3, vol))
	fmt.Printf("\tps.%s.R2_4 = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.R2_4, vol))
	fmt.Printf("\tps.%s.PKAact = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.PKAact, vol))
	fmt.Printf("\tps.%s.AC1ATPC = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.AC1ATPC, vol))
	fmt.Printf("\tps.%s.PDEcAMPC = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(ps.PDEcAMPC, vol))
}

func (ps *PKAVars) Zero() {
	ps.AC1 = 0
	ps.AC1act = 0
	ps.PDEact = 0
	ps.ATP = 0
	ps.CAMP = 0
	ps.AMP = 0
	ps.R2C2 = 0
	ps.R2C2_B = 0
	ps.R2C2_BB = 0
	ps.R2C2_AB = 0
	ps.R2C2_ABB = 0
	ps.R2C2_4 = 0
	ps.R2C_3 = 0
	ps.R2C_4 = 0
	ps.R2_3 = 0
	ps.R2_4 = 0
	ps.PKAact = 0
	ps.AC1ATPC = 0
	ps.PDEcAMPC = 0
}

func (ps *PKAVars) Integrate(d *PKAVars) {
	chem.Integrate(&ps.AC1, d.AC1)
	chem.Integrate(&ps.AC1act, d.AC1act)
	// PDEact buffered
	// ATP buffered
	chem.Integrate(&ps.CAMP, d.CAMP)
	// AMP buffered
	chem.Integrate(&ps.R2C2, d.R2C2)
	chem.Integrate(&ps.R2C2_B, d.R2C2_B)
	chem.Integrate(&ps.R2C2_BB, d.R2C2_BB)
	chem.Integrate(&ps.R2C2_AB, d.R2C2_AB)
	chem.Integrate(&ps.R2C2_ABB, d.R2C2_ABB)
	chem.Integrate(&ps.R2C2_4, d.R2C2_4)
	chem.Integrate(&ps.R2C_3, d.R2C_3)
	chem.Integrate(&ps.R2C_4, d.R2C_4)
	chem.Integrate(&ps.R2_3, d.R2_3)
	chem.Integrate(&ps.R2_4, d.R2_4)
	chem.Integrate(&ps.PKAact, d.PKAact)
	// chem.Integrate(&ps.AC1ATPC, d.AC1ATPC) // set directly
	chem.Integrate(&ps.PDEcAMPC, d.PDEcAMPC)
}

func (ps *PKAVars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	dir.Float64(pre + "AC1act").AppendRowFloat(chem.CoFromN(ps.AC1act, vol))
	dir.Float64(pre + "cAMP").AppendRowFloat(chem.CoFromN(ps.CAMP, vol))
	dir.Float64(pre + "PKAact").AppendRowFloat(chem.CoFromN(ps.PKAact, vol))
	// dir.Float64(pre+"AC1").AppendRowFloat(chem.CoFromN(ps.AC1, vol))
	// dir.Float64(pre+"R2C2").AppendRowFloat(chem.CoFromN(ps.R2C2, vol))
	// dir.Float64(pre+"R2C2_B").AppendRowFloat(chem.CoFromN(ps.R2C2_B, vol))
	// dir.Float64(pre+"R2C2_ABB").AppendRowFloat(chem.CoFromN(ps.R2C2_ABB, vol))
}

// PKAState is overall intracellular Ca-driven signaling states
// for PKA binding and phosphorylation with cAMP
// 32 state vars total
type PKAState struct {

	// in cytosol -- volume = 0.08 fl = 48
	Cyt PKAVars

	// in PSD -- volume = 0.02 fl = 12
	PSD PKAVars
}

func (ps *PKAState) Init(uk *Urakubo) {
	ps.Cyt.Init(uk, CytVol)
	ps.PSD.Init(uk, PSDVol)

	vol := float64(PSDVol)
	ps.PSD.AC1act = chem.CoToN(0.0003286, vol)
	ps.PSD.AC1ATPC = chem.CoToN(3.214e+05, vol)
}

func (ps *PKAState) InitCode() {
	fmt.Printf("\nPKAState:\n")
	ps.Cyt.InitCode(CytVol, "Cyt")
	ps.PSD.InitCode(PSDVol, "PSD")
}

func (ps *PKAState) Zero() {
	ps.Cyt.Zero()
	ps.PSD.Zero()
}

func (ps *PKAState) Integrate(d *PKAState) {
	ps.Cyt.Integrate(&d.Cyt)
	ps.PSD.Integrate(&d.PSD)
}

func (ps *PKAState) Stats(dir *tensorfs.Node) {
	ps.Cyt.Stats(dir, CytVol, "Cyt_")
	ps.PSD.Stats(dir, PSDVol, "PSD_")
}

// PKAParams are the parameters governing the
// PKA binding and phosphorylation with cAMP
type PKAParams struct {

	// 1: 3Ca-CaM + AC1 -> AC1act
	CaMAC1 chem.React

	// 2: basal activity of ATP -> cAMP without AC1 enzyme
	ATPcAMP chem.React

	// 3: R2C2 + cAMP = cAMP-bind-site-B
	R2C2_B chem.React

	// 4: R2C2-cAMP B + cAMP -> BB = cAMP-bind-site-B[1]
	R2C2_B1 chem.React

	// 5: R2C2-cAMP B + cAMP -> AB = cAMP-bind-site-A[1]
	R2C2_A1 chem.React

	// 6: R2C2-cAMP BB + cAMP -> ABB = cAMP-bind-site-A[2]
	R2C2_A2 chem.React

	// 7: R2C2-cAMP AB + cAMP -> ABB = cAMP-bind-site-B[2]
	R2C2_B2 chem.React

	// 8: R2C2-cAMP ABB + cAMP -> 4 = cAMP-bind-site-A
	R2C2_A chem.React

	// 9: R2C-3cAMP -> R2C-4cAMP = cAMP-bind-site-A[3]
	R2C_A3 chem.React

	// 10: R2-3cAMP -> R2-4cAMP = cAMP-bind-site-A[4]
	R2_A4 chem.React

	// 11: R2C-3cAMP + PKAact -> R2C2-3cAMP ABB (backwards) = Release-C1[1] -- Fig SI4 R2-3 -> R2C-3
	R2C_3 chem.React

	// 12: R2C-4cAMP + PKAact -> R2C2-4cAMP (backwards) = Release-C1
	R2C_4 chem.React

	// 13: R2-3cAMP + PKAact -> R2C-3cAMP (backwards) = Release-C2[1]
	R2_3 chem.React

	// 14: R2-4cAMP + PKAact -> R2C-4cAMP (backwards) = Release-C2
	R2_4 chem.React

	// 15: AC1act catalyzing ATP -> cAMP -- table SIg numbered 9 -> 15 -- note: uses EnzRate not std Enz -- does not consume AC1act
	AC1ATP chem.EnzRate

	// 16: PDE1act catalyzing cAMP -> AMP -- table SIg numbered 10 -> 16
	PDEcAMP chem.Enz

	// PKA diffusion between Cyt and PSD
	PKADiffuse chem.Diffuse

	// cAMP diffusion between Cyt and PSD
	CAMPDiffuse chem.Diffuse
}

func (cp *PKAParams) Defaults() {
	// note: following are all in Cyt -- PSD is 4x for first values
	// See React docs for more info
	cp.CaMAC1.SetVol(6, CytVol, 1)      // 1: 6 μM-1 = 0.125 -- NOTE: error in table (5) vs model 0.10416
	cp.ATPcAMP.Set(4.0e-7, 0)           // 2: called "leak"
	cp.R2C2_B.SetVol(0.2, CytVol, 0.1)  // 3: 0.2 μM-1 = 0.0041667 = cAMP-bind-site-B
	cp.R2C2_B1.SetVol(0.1, CytVol, 0.2) // 4: 0.1 μM-1 = 0.002083, cAMP-bind-site-B[1]
	cp.R2C2_A1.SetVol(2, CytVol, 5)     // 5: 2 μM-1 = 0.041667 = cAMP-bind-site-A[1]
	cp.R2C2_A2.SetVol(4, CytVol, 5)     // 6: 4 μM-1 = 0.083333 = cAMP-bind-site-A[2]
	cp.R2C2_B2.SetVol(0.1, CytVol, 0.1) // 7: 0.1 μM-1 = 0.002083 = cAMP-bind-site-B[2]
	cp.R2C2_A.SetVol(2, CytVol, 10)     // 8: 2 μM-1 = 0.041667 = cAMP-bind-site-A
	cp.R2C_A3.SetVol(20, CytVol, 1)     // 9: 20 μM-1 = 0.41667 = cAMP-bind-site-A[3]
	cp.R2_A4.SetVol(200, CytVol, 0.1)   // 10: 200 μM-1 = 4.1667 = cAMP-bind-site-A[4]

	cp.R2C_3.SetVol(10, CytVol, 2) // 11: 10 μM-1 = 0.20833 = Release-C1[1]
	cp.R2C_4.SetVol(1, CytVol, 20) // 12: 1 μM-1 = 0.020833 = Release-C1
	cp.R2_3.SetVol(20, CytVol, 1)  // 13: 20 μM-1 = 0.41667 = Release-C2[1]
	cp.R2_4.SetVol(2, CytVol, 10)  // 14: 2 μM-1 = 0.041667 = Release-C2

	cp.AC1ATP.SetKmVol(40, CytVol, 40, 10)  // 15: Km = 40 * 48 (ac) = 0.026042
	cp.PDEcAMP.SetKmVol(10, CytVol, 80, 20) // 16: Km = 10 = 0.20834

	cp.PKADiffuse.SetSym(32.0 / 0.0225)
	cp.CAMPDiffuse.SetSym(500.0 / 0.0225)
}

// StepPKA does the PKA + cAMP reactions, in a given region
// cCaM, dCaM = current, delta 3CaCaM from CaMKIIVars
func (cp *PKAParams) StepPKA(vol float64, c, d *PKAVars, cCaM float64, dCaM *float64) {
	kf := CytVol / vol
	var dummy float64
	cp.CaMAC1.StepK(kf, c.AC1, cCaM, c.AC1act, &d.AC1, dCaM, &d.AC1act)                   // 1
	cp.ATPcAMP.StepK(kf, c.ATP, 1, c.CAMP, &d.ATP, &dummy, &d.CAMP)                       // 2
	cp.R2C2_B.StepK(kf, c.CAMP, c.R2C2, c.R2C2_B, &d.CAMP, &d.R2C2, &d.R2C2_B)            // 3
	cp.R2C2_B1.StepK(kf, c.CAMP, c.R2C2_B, c.R2C2_BB, &d.CAMP, &d.R2C2, &d.R2C2_BB)       // 4
	cp.R2C2_A1.StepK(kf, c.CAMP, c.R2C2_B, c.R2C2_AB, &d.CAMP, &d.R2C2, &d.R2C2_AB)       // 5
	cp.R2C2_A2.StepK(kf, c.CAMP, c.R2C2_BB, c.R2C2_ABB, &d.CAMP, &d.R2C2_BB, &d.R2C2_ABB) // 6
	cp.R2C2_B2.StepK(kf, c.CAMP, c.R2C2_AB, c.R2C2_ABB, &d.CAMP, &d.R2C2_AB, &d.R2C2_ABB) // 7
	cp.R2C2_A.StepK(kf, c.CAMP, c.R2C2_ABB, c.R2C2_4, &d.CAMP, &d.R2C2_ABB, &d.R2C2_4)    // 8
	cp.R2C_A3.StepK(kf, c.CAMP, c.R2C_3, c.R2C_4, &d.CAMP, &d.R2C_3, &d.R2C_4)            // 9
	cp.R2_A4.StepK(kf, c.CAMP, c.R2_3, c.R2_4, &d.CAMP, &d.R2_3, &d.R2_4)                 // 10

	cp.R2C_3.StepK(kf, c.PKAact, c.R2C_3, c.R2C2_ABB, &d.PKAact, &d.R2C_3, &d.R2C2_ABB) // 11
	cp.R2C_4.StepK(kf, c.PKAact, c.R2C_4, c.R2C2_4, &d.PKAact, &d.R2C_4, &d.R2C2_4)     // 12
	cp.R2_3.StepK(kf, c.PKAact, c.R2_3, c.R2C_3, &d.PKAact, &d.R2_3, &d.R2C_3)          // 13
	cp.R2_4.StepK(kf, c.PKAact, c.R2_4, c.R2C_4, &d.PKAact, &d.R2_4, &d.R2C_4)          // 14

	// cs, ce, ds, dp, cc
	cp.AC1ATP.StepK(kf, c.ATP, c.AC1act, &d.ATP, &d.CAMP, &c.AC1ATPC)
	// cs, ce, cc, cp -> ds, de, dc, dp
	cp.PDEcAMP.StepK(kf, c.CAMP, c.PDEact, c.PDEcAMPC, c.AMP, &d.CAMP, &d.PDEact, &d.PDEcAMPC, &d.AMP)
}

// StepDiffuse does diffusion update, c=current, d=delta
func (cp *PKAParams) StepDiffuse(c, d *PKAState) {
	cp.PKADiffuse.Step(c.Cyt.R2C2, c.PSD.R2C2, CytVol, PSDVol, &d.Cyt.R2C2, &d.PSD.R2C2)
	cp.PKADiffuse.Step(c.Cyt.R2C2_B, c.PSD.R2C2_B, CytVol, PSDVol, &d.Cyt.R2C2_B, &d.PSD.R2C2_B)
	cp.PKADiffuse.Step(c.Cyt.R2C2_BB, c.PSD.R2C2_BB, CytVol, PSDVol, &d.Cyt.R2C2_BB, &d.PSD.R2C2_BB)
	cp.PKADiffuse.Step(c.Cyt.R2C2_AB, c.PSD.R2C2_AB, CytVol, PSDVol, &d.Cyt.R2C2_AB, &d.PSD.R2C2_AB)
	cp.PKADiffuse.Step(c.Cyt.R2C2_ABB, c.PSD.R2C2_ABB, CytVol, PSDVol, &d.Cyt.R2C2_ABB, &d.PSD.R2C2_ABB)
	cp.PKADiffuse.Step(c.Cyt.R2C2_4, c.PSD.R2C2_4, CytVol, PSDVol, &d.Cyt.R2C2_4, &d.PSD.R2C2_4)
	cp.PKADiffuse.Step(c.Cyt.R2C_3, c.PSD.R2C_3, CytVol, PSDVol, &d.Cyt.R2C_3, &d.PSD.R2C_3)
	cp.PKADiffuse.Step(c.Cyt.R2C_4, c.PSD.R2C_4, CytVol, PSDVol, &d.Cyt.R2C_4, &d.PSD.R2C_4)
	cp.PKADiffuse.Step(c.Cyt.R2_3, c.PSD.R2_3, CytVol, PSDVol, &d.Cyt.R2_3, &d.PSD.R2_3)
	cp.PKADiffuse.Step(c.Cyt.R2_4, c.PSD.R2_4, CytVol, PSDVol, &d.Cyt.R2_4, &d.PSD.R2_4)

	cp.CAMPDiffuse.Step(c.Cyt.CAMP, c.PSD.CAMP, CytVol, PSDVol, &d.Cyt.CAMP, &d.PSD.CAMP)
	cp.CAMPDiffuse.Step(c.Cyt.PDEact, c.PSD.PDEact, CytVol, PSDVol, &d.Cyt.PDEact, &d.PSD.PDEact)
}

// Step does full PKA updating, c=current, d=delta
func (cp *PKAParams) Step(c, d *PKAState, cCaM, dCaM *CaMState) {
	cp.StepPKA(CytVol, &c.Cyt, &d.Cyt, cCaM.Cyt.CaM[3], &dCaM.Cyt.CaM[3])
	cp.StepPKA(PSDVol, &c.PSD, &d.PSD, cCaM.PSD.CaM[3], &dCaM.PSD.CaM[3])
	cp.StepDiffuse(c, d)
}
