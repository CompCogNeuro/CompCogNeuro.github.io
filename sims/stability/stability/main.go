// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/CompCogNeuro/CompCogNeuro.github.io/sims/stability"
	"github.com/emer/emergent/v2/egui"
)

func main() { egui.Run[stability.Sim, stability.Config]() }
