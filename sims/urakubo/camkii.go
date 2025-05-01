// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"fmt"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// CaCaMKIIVars are intracellular Ca-driven signaling variables for the
// CaMKII+CaM binding -- each can have different numbers of Ca bound
// Dupont = DupontHouartDekonnick03, has W* terms used in Genesis code
// stores N values -- Co = Concentration computed by volume as needed
type CaCaMKIIVars struct {

	// CaMKII-CaM bound together = WBn in Dupont
	CaM_CaMKII float64

	// CaMKIIP-CaM bound together, P = phosphorylated at Thr286 = WTn in Dupont
	CaM_CaMKIIP float64
}

func (cs *CaCaMKIIVars) Init(uk *Urakubo, vol float64) {
	cs.Zero()
}

func (cs *CaCaMKIIVars) Zero() {
	cs.CaM_CaMKII = 0
	cs.CaM_CaMKIIP = 0
}

func (cs *CaCaMKIIVars) Integrate(d *CaCaMKIIVars) {
	chem.Integrate(&cs.CaM_CaMKII, d.CaM_CaMKII)
	chem.Integrate(&cs.CaM_CaMKIIP, d.CaM_CaMKIIP)
}

// AutoPVars hold the auto-phosphorylation variables, for CaMKII
type AutoPVars struct {

	// total active CaMKII
	Act float64

	// total CaMKII across all states
	Total float64

	// rate constant for further autophosphorylation as function of current state
	K float64
}

func (av *AutoPVars) Zero() {
	av.Act = 0
	av.Total = 0
	av.K = 0
}

// CaMKIIVars are intracellular Ca-driven signaling states
// for CaMKII binding and phosphorylation with CaM + Ca
// Dupont = DupontHouartDekonnick03, has W* terms used in Genesis code
// stores N values -- Co = Concentration computed by volume as needed
type CaMKIIVars struct {

	// increasing levels of Ca binding, 0-3
	Ca [4]CaCaMKIIVars

	// unbound CaMKII = CaM kinase II -- WI in Dupont -- this is the inactive form for NMDA GluN2B binding
	CaMKII float64

	// unbound CaMKII P = phosphorylated at Thr286 -- shown with * in Figure S13 = WA in Dupont -- this is the active form for NMDA GluN2B binding
	CaMKIIP float64

	// PP1+CaMKIIP complex for PP1Thr286 enzyme reaction
	PP1Thr286C float64

	// PP2A+CaMKIIP complex for PP2AThr286 enzyme reaction
	PP2AThr286C float64

	// auto-phosphorylation state
	Auto AutoPVars `view:"inline" inactive:"+"`
}

func (cs *CaMKIIVars) Init(uk *Urakubo, vol float64) {
	for i := range cs.Ca {
		cs.Ca[i].Init(uk, vol)
	}
	cs.CaMKII = chem.CoToN(20, vol)
	cs.CaMKIIP = 0 // WA
	cs.PP1Thr286C = 0
	cs.PP2AThr286C = 0

	if uk.InitBaseline {
		cs.CaMKII = chem.CoToN(19.28, vol) // orig: 20
	}

	cs.ActiveK()
}

// Generate Code for Initializing
func (cs *CaMKIIVars) InitCode(vol float64, pre string) {
	for i := range cs.Ca {
		fmt.Printf("\tcs.%s.Ca[%d].CaM_CaMKII = chem.CoToN(%.4g, vol)\n", pre, i, chem.CoFromN(cs.Ca[i].CaM_CaMKII, vol))
		fmt.Printf("\tcs.%s.Ca[%d].CaM_CaMKIIP = chem.CoToN(%.4g, vol)\n", pre, i, chem.CoFromN(cs.Ca[i].CaM_CaMKIIP, vol))
	}
	fmt.Printf("\tcs.%s.CaMKII = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(cs.CaMKII, vol))
	fmt.Printf("\tcs.%s.CaMKIIP = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(cs.CaMKIIP, vol))
	fmt.Printf("\tcs.%s.PP1Thr286C = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(cs.PP1Thr286C, vol))
	fmt.Printf("\tcs.%s.PP2AThr286C = chem.CoToN(%.4g, vol)\n", pre, chem.CoFromN(cs.PP2AThr286C, vol))
}

func (cs *CaMKIIVars) Zero() {
	for i := range cs.Ca {
		cs.Ca[i].Zero()
	}
	cs.CaMKII = 0
	cs.CaMKIIP = 0
	cs.PP1Thr286C = 0
	cs.PP2AThr286C = 0
	cs.Auto.Zero()
}

func (cs *CaMKIIVars) Integrate(d *CaMKIIVars) {
	for i := range cs.Ca {
		cs.Ca[i].Integrate(&d.Ca[i])
	}
	chem.Integrate(&cs.CaMKII, d.CaMKII)
	chem.Integrate(&cs.CaMKIIP, d.CaMKIIP)
	chem.Integrate(&cs.PP1Thr286C, d.PP1Thr286C)
	chem.Integrate(&cs.PP2AThr286C, d.PP2AThr286C)
	cs.ActiveK()
}

// ActiveK updates active, total, and the Auto.K auto-phosphorylation rate constant
// Code is from genesis_customizing/T286Phos/T286Phos.c and would be impossible to
// reconstruct without that source (my first guess was wildy off, based only on
// the supplement)
func (cs *CaMKIIVars) ActiveK() {
	WI := cs.CaMKII
	WA := cs.CaMKIIP

	var WB, WT float64

	for i := 0; i < 3; i++ {
		WB += cs.Ca[i].CaM_CaMKII
		WT += cs.Ca[i].CaM_CaMKIIP
	}
	WB += cs.Ca[3].CaM_CaMKII
	WP := cs.Ca[3].CaM_CaMKIIP

	TotalW := WI + WB + WP + WT + WA
	Wb := WB / TotalW
	Wp := WP / TotalW
	Wt := WT / TotalW
	Wa := WA / TotalW
	cb := 0.75
	ct := 0.8
	ca := 0.8

	T := Wb + Wp + Wt + Wa
	tmp := T * (-0.22 + 1.826*T + -0.8*T*T)
	tmp *= 0.75 * (cb*Wb + Wp + ct*Wt + ca*Wa)
	if tmp < 0 {
		tmp = 0
	}
	cs.Auto.K = 0.29 * tmp
	cs.Auto.Act = cb*WB + WP + ct*WT + ca*WA
	cs.Auto.Total = T
}

func (cs *CaMKIIVars) Stats(dir *tensorfs.Node, vol float64, pre string) {
	dir.Float64(pre + "CaMKIIact").AppendRowFloat(chem.CoFromN(cs.Auto.Act, vol))
	// dir.Float64(pre+"Ca0CaM_CaMKII").AppendRowFloat(chem.CoFromN(cs.Ca[0].CaM_CaMKII, vol))
	// dir.Float64(pre+"Ca1CaM_CaMKII").AppendRowFloat(chem.CoFromN(cs.Ca[1].CaM_CaMKII, vol))
	// dir.Float64(pre+"Ca0CaM_CaMKIIP").AppendRowFloat(chem.CoFromN(cs.Ca[0].CaM_CaMKIIP, vol))
	// dir.Float64(pre+"Ca1CaM_CaMKIIP").AppendRowFloat(chem.CoFromN(cs.Ca[1].CaM_CaMKIIP, vol))
	dir.Float64(pre + "CaMKII").AppendRowFloat(chem.CoFromN(cs.CaMKII, vol))
	dir.Float64(pre + "CaMKIIP").AppendRowFloat(chem.CoFromN(cs.CaMKIIP, vol))
	// dir.Float64(pre+"CaMKII_AutoK").AppendRowFloat(chem.CoFromN(cs.Auto.K, vol))
}

// CaMKIIState is overall intracellular Ca-driven signaling states
// for CaMKII in Cyt and PSD
// 28 state vars total
type CaMKIIState struct {

	// in cytosol -- volume = 0.08 fl = 48
	Cyt CaMKIIVars

	// in PSD -- volume = 0.02 fl = 12
	PSD CaMKIIVars
}

func (cs *CaMKIIState) Init(uk *Urakubo) {
	cs.Cyt.Init(uk, CytVol)
	cs.PSD.Init(uk, PSDVol)

	if uk.InitBaseline {
		vol := float64(CytVol)
		cs.Cyt.Ca[0].CaM_CaMKII = chem.CoToN(0.2612, vol)
		cs.Cyt.Ca[1].CaM_CaMKII = chem.CoToN(0.003344, vol)
		cs.Cyt.Ca[2].CaM_CaMKII = chem.CoToN(2.228e-05, vol)
		cs.Cyt.Ca[3].CaM_CaMKII = chem.CoToN(0.0014, vol)
		cs.Cyt.CaMKII = chem.CoToN(19.37, vol) // orig: 20

		vol = PSDVol
		cs.PSD.Ca[0].CaM_CaMKII = chem.CoToN(1.991, vol)
		cs.PSD.Ca[1].CaM_CaMKII = chem.CoToN(0.02548, vol)
		cs.PSD.Ca[2].CaM_CaMKII = chem.CoToN(0.0001698, vol)
		cs.PSD.Ca[3].CaM_CaMKII = chem.CoToN(0.01968, vol)
		cs.PSD.CaMKII = chem.CoToN(19.43, vol) // orig: 20
	}
	cs.Cyt.ActiveK()
	cs.PSD.ActiveK()
}

func (cs *CaMKIIState) InitCode() {
	fmt.Printf("\nCaMKIIState:\n")
	cs.Cyt.InitCode(CytVol, "Cyt")
	cs.PSD.InitCode(PSDVol, "PSD")
}

func (cs *CaMKIIState) Zero() {
	cs.Cyt.Zero()
	cs.PSD.Zero()
}

func (cs *CaMKIIState) Integrate(d *CaMKIIState) {
	cs.Cyt.Integrate(&d.Cyt)
	cs.PSD.Integrate(&d.PSD)
}

func (cs *CaMKIIState) Stats(dir *tensorfs.Node) {
	cs.Cyt.Stats(dir, CytVol, "Cyt_")
	cs.PSD.Stats(dir, PSDVol, "PSD_")
}

// CaMKIIParams are the parameters governing the Ca+CaM binding
type CaMKIIParams struct {

	// 1: Ca+CaM-CaMKII -> 1CaCaM-CaMKII = CaM-bind-Ca
	CaCaM01 chem.React

	// 2: Ca+1CaM-CaMKII -> 2CaCaM-CaMKII = CaMCa-bind-Ca
	CaCaM12 chem.React

	// 6: Ca+2CaCaM-CaMKII -> 3CaCaM-CaMKII = CaMCa2-bind-Ca
	CaCaM23 chem.React

	// 4: CaM+CaMKII -> CaM-CaMKII [0-2] -- kIB_kBI_[0-2] -- WI = plain CaMKII, WBn = CaM bound
	CaMCaMKII chem.React

	// 5: 3CaCaM+CaMKII -> 3CaCaM-CaMKII = kIB_kBI_3 -- active CaM binds strongly
	CaMCaMKII3 chem.React

	// 9: CaM+CaMKIIP -> CaM-CaMKIIP = kAT_kTA -- T286P causes strong CaM binding
	CaMCaMKIIP chem.React

	// 8: Ca+nCaCaM-CaMKIIP -> n+1CaCaM-CaMKIIP = kTP_PT_*
	CaCaM_CaMKIIP chem.React

	// 10: PP1 dephosphorylating CaMKIIP
	PP1Thr286 chem.Enz

	// 11: PP2A dephosphorylating CaMKIIP
	PP2AThr286 chem.Enz

	// CaMKII symmetric diffusion between Cyt and PSD -- only for WI
	CaMKIIDiffuse chem.Diffuse

	// CaMKIIP diffusion between Cyt and PSD -- asymmetric, everything else
	CaMKIIPDiffuse chem.Diffuse
}

func (cp *CaMKIIParams) Defaults() {
	// note: following are all in Cyt -- PSD is 4x for first values
	// See React docs for more info
	cp.CaCaM01.SetVol(51.202, CytVol, 200) // 1: 51.202 μM-1 = 1.0667, PSD 4.2667 = CaM-bind-Ca
	cp.CaCaM12.SetVol(133.3, CytVol, 1000) // 2: 133.3 μM-1 = 2.7771, PSD 11.108 = CaMCa-bind-Ca
	cp.CaCaM23.SetVol(25.6, CytVol, 400)   // 6: 25.6 μM-1 = 0.53333, PSD 2.1333 = CaMCa2-bind-Ca

	cp.CaMCaMKII.SetVol(0.0004, CytVol, 1) // 4: 0.0004 μM-1 = 8.3333e-6, PSD 3.3333e-5 = kIB_kBI_[0-2]
	cp.CaMCaMKII3.SetVol(8, CytVol, 1)     // 5: 8 μM-1 = 0.16667, PSD 3.3333e-5 = kIB_kBI_3 -- 3CaCaM is active
	cp.CaMCaMKIIP.SetVol(8, CytVol, 0.001) // 9: 8 μM-1 = 0.16667, PSD 0.66667 = kAT_kTA

	cp.CaCaM_CaMKIIP.SetVol(1, CytVol, 1) // 8: 1 μM-1 = 0.020834, PSD 0.0833335 = kTP_PT_*

	cp.PP1Thr286.SetKmVol(11, CytVol, 1.34, 0.335)  // 10: 11 μM Km = 0.0031724
	cp.PP2AThr286.SetKmVol(11, CytVol, 1.34, 0.335) // 11: 11 μM Km = 0.0031724

	cp.CaMKIIDiffuse.SetSym(6.0 / 0.0225)
	cp.CaMKIIPDiffuse.Set(6.0/0.0225, 0.6/0.0225)
}

//////// No N2B versions

// StepCaMKII does the bulk of Ca + CaM + CaMKII binding reactions, in a given region
// cCa, nCa = current next Ca
func (cp *CaMKIIParams) StepCaMKII(vol float64, c, d *CaMKIIVars, cm, dm *CaMVars, cCa, pp1, pp2a float64, dCa, dpp1, dpp2a *float64) {
	kf := CytVol / vol

	cp.CaCaM01.StepK(kf, c.Ca[0].CaM_CaMKII, cCa, c.Ca[1].CaM_CaMKII, &d.Ca[0].CaM_CaMKII, dCa, &d.Ca[1].CaM_CaMKII) // 1
	cp.CaCaM12.StepK(kf, c.Ca[1].CaM_CaMKII, cCa, c.Ca[2].CaM_CaMKII, &d.Ca[1].CaM_CaMKII, dCa, &d.Ca[2].CaM_CaMKII) // 2

	for i := 0; i < 3; i++ {
		cp.CaMCaMKII.StepK(kf, cm.CaM[i], c.CaMKII, c.Ca[i].CaM_CaMKII, &dm.CaM[i], &d.CaMKII, &d.Ca[i].CaM_CaMKII) // 4
	}
	cp.CaMCaMKII3.StepK(kf, cm.CaM[3], c.CaMKII, c.Ca[3].CaM_CaMKII, &dm.CaM[3], &d.CaMKII, &d.Ca[3].CaM_CaMKII) // 5

	cp.CaMCaMKIIP.StepK(kf, cm.CaM[0], c.CaMKIIP, c.Ca[0].CaM_CaMKIIP, &dm.CaM[0], &d.CaMKIIP, &d.Ca[0].CaM_CaMKIIP) // 9

	for i := 0; i < 3; i++ {
		cp.CaCaM_CaMKIIP.StepK(kf, c.Ca[i].CaM_CaMKIIP, cCa, c.Ca[i+1].CaM_CaMKIIP, &d.Ca[i].CaM_CaMKIIP, dCa, &d.Ca[i+1].CaM_CaMKIIP) // 8
	}

	// cs, ce, cc, cp -> ds, de, dc, dp
	cp.PP1Thr286.StepK(kf, c.CaMKIIP, pp1, c.PP1Thr286C, c.CaMKII, &d.CaMKIIP, dpp1, &d.PP1Thr286C, &d.CaMKII) // 10
	if dpp2a != nil {
		cp.PP2AThr286.StepK(kf, c.CaMKIIP, pp2a, c.PP2AThr286C, c.CaMKII, &d.CaMKIIP, dpp2a, &d.PP2AThr286C, &d.CaMKII) // 11
	}

	for i := 0; i < 4; i++ {
		cc := &c.Ca[i]
		dc := &d.Ca[i]
		dak := c.Auto.K * cc.CaM_CaMKII
		dc.CaM_CaMKIIP += dak
		dc.CaM_CaMKII -= dak
		// cs, ce, cc, cp -> ds, de, dc, dp
		cp.PP1Thr286.StepK(kf, cc.CaM_CaMKIIP, pp1, c.PP1Thr286C, cc.CaM_CaMKII, &dc.CaM_CaMKIIP, dpp1, &d.PP1Thr286C, &dc.CaM_CaMKII) // 10
		if dpp2a != nil {
			cp.PP2AThr286.StepK(kf, cc.CaM_CaMKIIP, pp2a, c.PP2AThr286C, cc.CaM_CaMKII, &dc.CaM_CaMKIIP, dpp2a, &d.PP2AThr286C, &dc.CaM_CaMKII) // 11
		}
	}
}

// StepDiffuse does Cyt <-> PSD diffusion
func (cp *CaMKIIParams) StepDiffuse(c, d *CaMKIIState) {
	for i := 0; i < 4; i++ {
		cc := &c.Cyt.Ca[i]
		cd := &c.PSD.Ca[i]
		dc := &d.Cyt.Ca[i]
		dd := &d.PSD.Ca[i]
		cp.CaMKIIPDiffuse.Step(cc.CaM_CaMKII, cd.CaM_CaMKII, CytVol, PSDVol, &dc.CaM_CaMKII, &dd.CaM_CaMKII)
		cp.CaMKIIPDiffuse.Step(cc.CaM_CaMKIIP, cd.CaM_CaMKIIP, CytVol, PSDVol, &dc.CaM_CaMKIIP, &dd.CaM_CaMKIIP)
	}
	cp.CaMKIIDiffuse.Step(c.Cyt.CaMKII, c.PSD.CaMKII, CytVol, PSDVol, &d.Cyt.CaMKII, &d.PSD.CaMKII)
	cp.CaMKIIPDiffuse.Step(c.Cyt.CaMKIIP, c.PSD.CaMKIIP, CytVol, PSDVol, &d.Cyt.CaMKIIP, &d.PSD.CaMKIIP) // P = N2B
}

// Step does one step of CaMKII updating, c=current, d=delta
// pp2a = current cyt pp2a
func (cp *CaMKIIParams) Step(c, d *CaMKIIState, cm, dm *CaMState, cCa, dCa *CaState, pp1, dpp1 *PP1State, pp2a float64, dpp2a *float64) {
	cp.StepCaMKII(CytVol, &c.Cyt, &d.Cyt, &cm.Cyt, &dm.Cyt, cCa.Cyt, pp1.Cyt.PP1act, pp2a, &dCa.Cyt, &dpp1.Cyt.PP1act, dpp2a)
	cp.StepCaMKII(PSDVol, &c.PSD, &d.PSD, &cm.PSD, &dm.PSD, cCa.PSD, pp1.PSD.PP1act, 0, &dCa.PSD, &dpp1.PSD.PP1act, nil)
	cp.StepDiffuse(c, d)
}
