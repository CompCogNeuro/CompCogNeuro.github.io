// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: code converted directly from Urakubo et al (2008)
// MODEL/genesis_customizing/NMDAR.c

package urakubo

import (
	"math"

	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// NMDARState holds NMDA receptor states, with allosteric dynamics
// from Urakubo et al, (2008)
// The [3] arrays correspond to Nt0, Nt1, Nt2: plain NMDA, 2CaM phos, 3CaM phos
type NMDARState struct {

	// level of Mg block as a function of membrane potential: 1/(1 + (1.5/3.57)exp(-0.062*Vm)
	Mg float64

	// voltage-dependent calcium flux driver: determines Jca as function of V, includes Mg factor
	Vca float64

	// overall calcium current = Vca * Pca * Nopen
	Jca float64

	// ionic conductance through the NMDA channel for driving Vm changes = Mg * GMax * Nopen
	G float64

	// Number in state 1
	N1 [3]float64

	// Number in state 2
	N2 [3]float64

	// Number in state 3
	N3 [3]float64

	// Number in state 4
	N4 [3]float64

	// Number in Open state -- actually open to allow Ca to flow
	No [3]float64

	// Total N of NMDAR plain = sum of 0 index in N* states
	Nt0 float64 `edit:"-"`

	// Total N of NMDAR_2Ca2+CaM = sum of 1 index in N* states
	Nt1 float64 `edit:"-"`

	// Total N of NMDAR_3Ca2+CaM = sum of 2 index in N* states
	Nt2 float64 `edit:"-"`

	// Total N in open state = sum(No[0..2])
	Nopen float64 `edit:"-"`

	// overall total -- should be conserved
	Ntotal float64 `edit:"-"`

	// number of available non-bound GluN2B S3103 binding sites.
	// CaMKII and DAPK1 compete to bind here -- Ca/ÇaM induce CaMKII binding
	// and Thr286 auto-phosphorylation maintains it,  while *de* phosphorylated DAPK1
	// at Ser308 preferentially binds.
	// The number of each bound is tracked in the CaMKII and DAPK1 PSD states.
	GluN2B float64
}

func (cs *NMDARState) Init(uk *Urakubo) {
	cs.Zero()
	cs.N2[0] = 1
	cs.Total()
	cs.GluN2B = cs.Ntotal // will be multiplied later

	// todo InitBaseline for basline binding
}

func (cs *NMDARState) Zero() {
	cs.Mg = 0
	cs.Vca = 0
	cs.Jca = 0
	for k := 0; k < 3; k++ {
		cs.N1[k] = 0
		cs.N2[k] = 0
		cs.N3[k] = 0
		cs.N4[k] = 0
		cs.No[k] = 0
	}
	cs.GluN2B = 0
}

func (cs *NMDARState) Integrate(d *NMDARState) {
	for k := 0; k < 3; k++ {
		chem.Integrate(&cs.N1[k], d.N1[k])
		chem.Integrate(&cs.N2[k], d.N2[k])
		chem.Integrate(&cs.N3[k], d.N3[k])
		chem.Integrate(&cs.N4[k], d.N4[k])
		chem.Integrate(&cs.No[k], d.No[k])
	}
	cs.Total()
	chem.Integrate(&cs.GluN2B, d.GluN2B)
}

func (cs *NMDARState) Total() {
	cs.Nt0 = cs.N1[0] + cs.N2[0] + cs.N3[0] + cs.N4[0] + cs.No[0]
	cs.Nt1 = cs.N1[1] + cs.N2[1] + cs.N3[1] + cs.N4[1] + cs.No[1]
	cs.Nt2 = cs.N1[2] + cs.N2[2] + cs.N3[2] + cs.N4[2] + cs.No[2]
	cs.Nopen = (cs.No[0] + cs.No[1] + cs.No[2])
	cs.Ntotal = cs.Nt0 + cs.Nt1 + cs.Nt2
}

func (cs *NMDARState) Stats(dir *tensorfs.Node) {
	pre := "NMDA_"
	dir.Float64(pre + "Mg").AppendRowFloat(cs.Mg)
	dir.Float64(pre + "Nopen").AppendRowFloat(cs.Nopen)
	dir.Float64(pre + "Jca").AppendRowFloat(cs.Jca)
	dir.Float64(pre + "G").AppendRowFloat(cs.G)
	dir.Float64(pre + "Nt0").AppendRowFloat(cs.Nt0)
	dir.Float64(pre + "Nt1").AppendRowFloat(cs.Nt1)
	dir.Float64(pre + "Nt2").AppendRowFloat(cs.Nt2)
	// dir.Float64(pre+"N1[0]").AppendRowFloat(cs.N1[0])
	// dir.Float64(pre+"N1[1]").AppendRowFloat(cs.N1[1])
	// dir.Float64(pre+"N1[2]").AppendRowFloat(cs.N1[2])
	dir.Float64("GluN2B").AppendRowFloat(cs.GluN2B)
}

// NMDARParams holds parameters for NMDA receptor with allosteric dynamics
// from Urakubo et al, (2008)
// The [3] arrays correspond to Nt0, Nt1, Nt2: plain NMDA, 2CaM, 3CaM
type NMDARParams struct {

	// reversal potential for NMDARs
	Erev float64 `default:"0"`

	// Normalization for Ca flux (pmol sec-1 mV-1)
	Pca float64 `default:"89635"`

	// maximum conductance (nS)
	Gmax float64 `default:"10"`

	// For N1 state, CaM + 2CaCaM reaction
	CaM1 chem.React

	// For N2 state, CaM + 2CaCaM reaction
	CaM2 chem.React

	// For N3 state, CaM + 2CaCaM reaction
	CaM3 chem.React

	// For N4 state, CaM + 2CaCaM reaction
	CaM4 chem.React

	// For No state, CaM + 2CaCaM reaction
	CaMo chem.React

	// for all states, rate of Ca binding for 2Ca-CaM <-> 3Ca-CaM -- same rates as used in CaM code (25.6, 400)
	CaCaM23 chem.React

	// Glu binding driving N1 <-> N2 transitions
	GluN12 chem.React

	// Glu binding driving N2 <-> N3 transitions
	GluN23 chem.React

	// Glu quantity (uM sec)
	Glu float64 `default:"0.4"`

	// total amount of GluN2B -- NMDAR is all done in normalized units,
	// but this needs to be in real units to interact with CaMKII and DAPK1 binding.
	GluN2BN float64
}

func (nr *NMDARParams) Defaults() {
	nr.Erev = 0
	nr.Pca = 1.7927e5 * 0.5 // SVR_PSD
	nr.Gmax = 10

	nr.CaM1.Set(400, 34.8)
	nr.CaM2.Set(400, 34.8)
	nr.CaM3.Set(4, 0.348)
	nr.CaM4.Set(3.458, 0.891)
	nr.CaMo.Set(1.994, 2.355)
	nr.CaCaM23.Set(25.6, 400) // Same as CaCaM23 from CaMParams

	nr.GluN12.Set(10, 25)
	nr.GluN23.Set(5, 50)

	nr.Glu = 0.4 // was 0.12

	//	nr.GluN2BN = chem.CoToN(10, PSDVol) // 120 works for N2B only (no DAPK1)
	nr.GluN2BN = chem.CoToN(2, PSDVol) // 120 works for N2B only (no DAPK1)
}

// Special init function for state
func (nr *NMDARParams) Init(cs *NMDARState) {
	cs.GluN2B = nr.GluN2BN
}

// Step increments NMDAR state in response to Ca/CaM binding
// ca = Ca2+ Co, c2 = 2Ca2+CaM Co, c3 = 3Ca2+CaM Co
func (nr *NMDARParams) StepCaCaM(c, d *NMDARState, vm, ca, c2, c3 float64, spike bool, dca *float64) {
	dt := chem.IntegrationDt

	if spike {
		T := nr.Glu
		for k := 0; k < 3; k++ {
			d.N1[k] = c.N1[k] * math.Exp(-nr.GluN12.Kf*T)
			d.N2[k] = (c.N1[k] * nr.GluN12.Kf) / (nr.GluN23.Kf - nr.GluN12.Kf)
			d.N2[k] = (math.Exp(-nr.GluN12.Kf*T) - math.Exp(-nr.GluN23.Kf*T)) * d.N2[k]
			d.N2[k] = c.N2[k]*math.Exp(-nr.GluN23.Kf*T) + d.N2[k]
			d.N3[k] = c.N3[k] - (d.N2[k] - c.N2[k]) - (d.N1[k] - c.N1[k])

			c.N1[k] = d.N1[k] // immediate reset
			c.N2[k] = d.N2[k]
			c.N3[k] = d.N3[k]
		}
	}

	j := int(math.Ceil(dt / 0.00003))

	for i := 0; i < j; i++ {
		// backward and forward reactions
		d.N1[0] = c.N2[0] * nr.GluN12.Kb
		d.N2[0] = c.N3[0]*nr.GluN23.Kb - c.N2[0]*nr.GluN12.Kb
		d.N3[0] = c.N4[0]*1.8 + c.No[0]*275 - c.N3[0]*(nr.GluN23.Kb+8+280)
		d.N4[0] = c.N3[0]*8 - c.N4[0]*1.8
		d.No[0] = c.N3[0]*280 - c.No[0]*275

		for k := 1; k < 3; k++ {
			d.N1[k] = c.N2[k] * nr.GluN12.Kb
			d.N2[k] = c.N3[k]*nr.GluN23.Kb - c.N2[k]*nr.GluN12.Kb
			d.N3[k] = c.N4[k]*2 + c.No[k]*2000 - c.N3[k]*(nr.GluN23.Kb+3+150)
			d.N4[k] = c.N3[k]*3 - c.N4[k]*2
			d.No[k] = c.N3[k]*150 - c.No[k]*2000
		}
		//
		// NMDAR binding to 2Ca2+CaM
		// ca, cb, cab, da, dab

		nr.CaM1.StepCB(c.N1[0], c2, c.N1[1], &d.N1[0], &d.N1[1])
		nr.CaM2.StepCB(c.N2[0], c2, c.N2[1], &d.N2[0], &d.N2[1])
		nr.CaM3.StepCB(c.N3[0], c2, c.N3[1], &d.N3[0], &d.N3[1])
		nr.CaM4.StepCB(c.N4[0], c2, c.N4[1], &d.N4[0], &d.N4[1])
		nr.CaMo.StepCB(c.No[0], c2, c.No[1], &d.No[0], &d.No[1])
		//
		// NMDAR binding to 3Ca2+CaM -- note: original code has clear
		// typo bug using c2 instead of c3 here!

		oldBug := false
		if oldBug {
			nr.CaM1.StepCB(c.N1[0], c2, c.N1[2], &d.N1[0], &d.N1[2])
			nr.CaM2.StepCB(c.N2[0], c2, c.N2[2], &d.N2[0], &d.N2[2])
			nr.CaM3.StepCB(c.N3[0], c2, c.N3[2], &d.N3[0], &d.N3[2])
			nr.CaM4.StepCB(c.N4[0], c2, c.N4[2], &d.N4[0], &d.N4[2])
			nr.CaMo.StepCB(c.No[0], c2, c.No[2], &d.No[0], &d.No[2])
		} else {
			nr.CaM1.StepCB(c.N1[0], c3, c.N1[2], &d.N1[0], &d.N1[2])
			nr.CaM2.StepCB(c.N2[0], c3, c.N2[2], &d.N2[0], &d.N2[2])
			nr.CaM3.StepCB(c.N3[0], c3, c.N3[2], &d.N3[0], &d.N3[2])
			nr.CaM4.StepCB(c.N4[0], c3, c.N4[2], &d.N4[0], &d.N4[2])
			nr.CaMo.StepCB(c.No[0], c3, c.No[2], &d.No[0], &d.No[2])
		}

		//
		// NMDAR-2Ca2+CaM binding Ca to/from NMDAR-3Ca2+CaM
		//
		nr.CaCaM23.StepCB(c.N1[1], ca, c.N1[2], &d.N1[1], &d.N1[2])
		nr.CaCaM23.StepCB(c.N2[1], ca, c.N2[2], &d.N2[1], &d.N2[2])
		nr.CaCaM23.StepCB(c.N3[1], ca, c.N3[2], &d.N3[1], &d.N3[2])
		nr.CaCaM23.StepCB(c.N4[1], ca, c.N4[2], &d.N4[1], &d.N4[2])
		nr.CaCaM23.StepCB(c.No[1], ca, c.No[2], &d.No[1], &d.No[2])
	}
}

// Step increments NMDAR state
// ca = Ca2+ Co, c2 = 2Ca2+CaM Co, c3 = 3Ca2+CaM Co
func (nr *NMDARParams) Step(c, d *NMDARState, vm, ca, c2, c3 float64, spike bool, dca *float64) {
	nr.StepCaCaM(c, d, vm, ca, c2, c3, spike, dca)

	c.Mg = 1 / (1 + 0.4202*math.Exp(-0.062*vm)) // Mg(1.5)/3.57
	if vm > -0.1 && vm < 0.1 {
		c.Vca = (1.0 / (0.0756 + 0.5*vm)) * c.Mg
	} else {
		c.Vca = -vm / (1 - math.Exp(0.0756*vm)) * c.Mg
	}
	c.Jca = c.Vca * nr.Pca * c.Nopen
	c.G = c.Mg * nr.Gmax * c.Nopen

	*dca += c.Jca * PSDVol
}
