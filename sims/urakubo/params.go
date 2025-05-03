// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"github.com/emer/axon/v2/axon"
)

var LayerParams = axon.LayerSheets{
	"Base": {
		{Sel: "Layer", Doc: "all defaults",
			Set: func(ly *axon.LayerParams) {
				ly.Acts.Spikes.Tr = 7
				ly.Acts.Spikes.RTau = 3 // maybe could go a bit wider even
				ly.Acts.Dt.VmC = 100
				ly.Acts.Dt.VmDendC = 100
				ly.Acts.Dt.VmSteps = 2
				ly.Acts.Dt.GeTau = 1 // not natural but fits spike current injection
				ly.Acts.VmRange.Max = -2
				ly.Acts.Spikes.ExpThr = -10 // note: critical to keep < Max!
				ly.Acts.Spikes.Thr = -45    // also bump up
				ly.Acts.Spikes.VmR = -55
				ly.Acts.Init.Vm = -65
				ly.Acts.Erev.L = -65
				// ly.Acts.Erev.I = -65
				// ly.Acts.Erev.K = -65 // todo: not clear if need this
				ly.Acts.AK.Gk = 0.05 // most distal = .48 per Migliore et al, 1999, but rescale in similar way to Vgccc
				ly.Acts.VGCC.Ge = 1.2
				ly.Acts.Mahp.Gk = 0
				ly.Acts.Sahp.Gk = 0
			}},
	},
}
