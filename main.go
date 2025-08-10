// Copyright (c) 2025, The CCN Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"fmt"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/text/csl"
	_ "cogentcore.org/core/text/tex" // include this to get math
	"cogentcore.org/core/tree"
	_ "github.com/compcogneuro/web/sims/yaegisims"
)

// NOTE: you must make a symbolic link to the zotero CCNLab CSL file as ccnlab.json
// in this directory, to generate references and have the generated reference links
// use the official APA style. https://www.zotero.org/groups/340666/ccnlab
// Must configure using BetterBibTeX for zotero: https://retorque.re/zotero-better-bibtex/
// todo: include link for configuring here

//go:generate mdcite -vv -refs ./ccnlab.json -d ./content

//go:embed content citedrefs.json
var econtent embed.FS

//go:embed icon.svg
var icon string

// sims is a map of sim names to functions that embed a sim GUI.
// This is only set on non-generatehtml platforms so that xyz does
// not interfere with html generation.
var sims map[string]func(tree.Node)

func main() {
	core.AppIcon = icon
	b := core.NewBody("Computational Cognitive Neuroscience")
	ct := content.NewContent(b).SetContent(econtent)
	refs, err := csl.OpenFS(econtent, "citedrefs.json")
	if err == nil {
		ct.References = csl.NewKeyList(refs)
	}

	ct.Context.ElementHandlers["ccn-sim"] = simHandler

	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}

func simHandler(ctx *htmlcore.Context) bool {
	if sims == nil {
		return true // for generatehtml
	}

	name := htmlcore.GetAttr(ctx.Node, "sim")
	fn := sims[name]
	if fn == nil {
		errors.Log(fmt.Errorf("sim %q not found", name))
		return true
	}

	fn(ctx.BlockParent)
	return true
}
