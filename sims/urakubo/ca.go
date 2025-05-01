// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/emergent/v2/chem"
)

// CaState records the Ca levels
// 2 state vars total
type CaState struct {

	// in cytosol
	Cyt float64

	// in PSD
	PSD float64
}

func (cs *CaState) Init(uk *Urakubo) {
	cs.Cyt = chem.CoToN(0.05, CytVol)
	cs.PSD = chem.CoToN(0.05, PSDVol)
}

func (cs *CaState) Zero() {
	cs.Cyt = 0
	cs.PSD = 0
}

func (cs *CaState) Integrate(d *CaState) {
	chem.Integrate(&cs.Cyt, d.Cyt)
	chem.Integrate(&cs.PSD, d.PSD)
}

func (cs *CaState) Stats(dir *tensorfs.Node) {
	dir.Float64("Cyt_Ca").AppendRowFloat(chem.CoFromN(cs.Cyt, CytVol))
	dir.Float64("PSD_Ca").AppendRowFloat(chem.CoFromN(cs.PSD, PSDVol))
}

// CaParams manages Ca parameters including soft buffering dynamics of calcium
type CaParams struct {
	// Ca buffering in the cytosol
	CytBuffer chem.Buffer

	// Ca buffering in the PSD
	PSDBuffer chem.Buffer

	// Ca diffusion between Cyt and PSD
	Diffuse chem.Diffuse

	// extra Ca injection values in N terms -- see SetInject for concentration.
	// Be sure to zero or update as needed
	InjectCa CaState

	// clamp ca by fixed values
	Clamp bool

	// clamped Ca values -- in N terms -- see SetClamp for concentration
	ClampCa CaState
}

func (cp *CaParams) Defaults() {
	// note: verified constants from initial_routines/Ca2_efflux.g
	// using showmsg /efflux_PSD / cytosol and showfield /efflux_PSD *
	// and doing the math.. replicates corresponding behavior in model
	cp.CytBuffer.SetTargVol(0.05, CytVol)
	cp.CytBuffer.K = (1.0426e5 * 0.8) / 12
	cp.PSDBuffer.SetTargVol(0.05, PSDVol)
	cp.PSDBuffer.K = (1.7927e5 * 0.8) / 12

	cp.Diffuse.SetSym(600.0 / 0.0225)
	cp.Clamp = false
}

func (cp *CaParams) Init() {
	cp.InjectCa.Zero()
}

// SetBuffTarg sets buffered target level of calcium in terms of concentrations
func (cp *CaParams) SetBuffTarg(cyt, psd float64) {
	cp.CytBuffer.SetTargVol(cyt, CytVol)
	cp.PSDBuffer.SetTargVol(psd, PSDVol)
}

// SetClamp sets clamped calcium levels in terms of concentrations
func (cp *CaParams) SetClamp(cyt, psd float64) {
	cp.Clamp = true
	cp.ClampCa.Cyt = chem.CoToN(cyt, CytVol)
	cp.ClampCa.PSD = chem.CoToN(psd, PSDVol)
}

// SetInject sets injected calcium levels in terms of concentrations
func (cp *CaParams) SetInject(cyt, psd float64) {
	cp.InjectCa.Cyt = chem.CoToN(cyt, CytVol)
	cp.InjectCa.PSD = chem.CoToN(psd, PSDVol)
}

func (cp *CaParams) Step(c *CaState, d *CaState) {
	if cp.Clamp {
		*c = cp.ClampCa
		d.Zero()
		return
	}
	d.PSD += cp.InjectCa.PSD
	d.Cyt += cp.InjectCa.Cyt
	cp.CytBuffer.Step(c.Cyt, &d.Cyt)
	cp.PSDBuffer.Step(c.PSD, &d.PSD)
	cp.Diffuse.Step(c.Cyt, c.PSD, CytVol, PSDVol, &d.Cyt, &d.PSD)
}
