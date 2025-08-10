// Copyright (c) 2025, The CCN Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !generatehtml

package main

import (
	"cogentcore.org/core/tree"
	"github.com/compcogneuro/web/sims/stability"
	"github.com/compcogneuro/web/sims/urakubo"
	"github.com/emer/axon/v2/sims/bgdorsal"
	"github.com/emer/axon/v2/sims/bgventral"
	"github.com/emer/axon/v2/sims/deepfsa"
	"github.com/emer/axon/v2/sims/inhib"
	"github.com/emer/axon/v2/sims/neuron"
	"github.com/emer/emergent/v2/egui"
)

func e[S, C any](b tree.Node) {
	egui.Embed[S, C](b)
}

func init() {
	sims = map[string]func(tree.Node){
		"bgdorsal":  e[bgdorsal.Sim, bgdorsal.Config],
		"bgventral": e[bgventral.Sim, bgventral.Config],
		"deepfsa":   e[deepfsa.Sim, deepfsa.Config],
		"inhib":     e[inhib.Sim, inhib.Config],
		"neuron":    e[neuron.Sim, neuron.Config],
		"stability": e[stability.Sim, stability.Config],
		"urakubo":   e[urakubo.Sim, urakubo.Config],
	}
}
