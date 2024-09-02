package main

import (
	"embed"

	"cogentcore.org/core/core"
	"cogentcore.org/core/pages"
)

//go:embed content
var content embed.FS

func main() {
	b := core.NewBody("CompCogNeuro")
	pg := pages.NewPage(b).SetContent(content)
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(pg.MakeToolbar)
	})
	b.RunMainWindow()
}
