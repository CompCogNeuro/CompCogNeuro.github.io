// Copyright (c) 2025, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cogentcore.org/core/cli"
	"github.com/CompCogNeuro/CompCogNeuro.github.io/sims/urakubo"
)

func main() {
	cfg := urakubo.NewConfig()
	opts := cli.DefaultOptions(cfg.Name, cfg.Title)
	opts.DefaultFiles = append(opts.DefaultFiles, "config.toml")
	cli.Run(opts, cfg, urakubo.RunSim)
}
