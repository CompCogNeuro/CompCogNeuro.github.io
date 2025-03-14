package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
)

//go:embed content
var econtent embed.FS

func main() {
	b := core.NewBody("CompCogNeuro")
	ct := content.NewContent(b).SetContent(econtent)
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}
