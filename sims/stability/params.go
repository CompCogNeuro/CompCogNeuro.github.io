// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stability

import "github.com/emer/axon/v2/axon"

// LayerParams sets the minimal non-default params.
// Base is always applied, and others can be optionally selected to apply on top of that.
var LayerParams = axon.LayerSheets{
	"Base": {
		{Sel: "Layer", Doc: "",
			Set: func(ly *axon.LayerParams) {
			}},
		{Sel: "#Layer0", Doc: "Input layer",
			Set: func(ly *axon.LayerParams) {
			}},
	},
}

// PathParams sets the minimal non-default params.
// Base is always applied, and others can be optionally selected to apply on top of that.
var PathParams = axon.PathSheets{
	"Base": {
		{Sel: "Path", Doc: "no learning",
			Set: func(pt *axon.PathParams) {
				pt.Learn.Learn.SetBool(false)
			}},
	},
}
