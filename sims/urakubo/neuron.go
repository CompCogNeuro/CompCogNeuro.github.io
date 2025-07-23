// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

// Extra state for neuron -- Vgcc and AK
type NeuronEx struct {
	// Vgcc Ca calcium contribution to PSD
	VgccJcaPSD float32 `desc:"Vgcc Ca calcium contribution to PSD"`
	// Vgcc Ca calcium contribution to Cyt
	VgccJcaCyt float32 `desc:"Vgcc Ca calcium contribution to Cyt"`
	// AK M gate -- activates with increasing Vm
	AKm float32 `desc:"AK M gate -- activates with increasing Vm"`
	// AK H gate -- deactivates with increasing Vm
	AKh float32 `desc:"AK H gate -- deactivates with increasing Vm"`
}

func (nex *NeuronEx) Init() {
	nex.VgccJcaPSD = 0
	nex.VgccJcaCyt = 0
	nex.AKm = 0
	nex.AKh = 1
}
