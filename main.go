package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/text/csl"
	_ "cogentcore.org/core/text/tex" // include this to get math
	_ "github.com/CompCogNeuro/CompCogNeuro.github.io/sims/yaegisims"
)

// NOTE: you must make a symbolic link to the zotero CCNLab CSL file as ccnlab.json
// in this directory, to generate references and have the generated reference links
// use the official APA style. https://www.zotero.org/groups/340666/ccnlab
// Must configure using BetterBibTeX for zotero: https://retorque.re/zotero-better-bibtex/
// todo: include link for configuring here

//go:generate mdcite -vv -refs ./ccnlab.json -d ./content

//go:embed content
var econtent embed.FS

//go:embed icon.svg
var icon string

func main() {
	core.AppIcon = icon
	b := core.NewBody("CompCogNeuro")
	ct := content.NewContent(b).SetContent(econtent)
	refs, err := csl.Open("citedrefs.json")
	if err == nil {
		ct.References = csl.NewKeyList(refs)
	}
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}
