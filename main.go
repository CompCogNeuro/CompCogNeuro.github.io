package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	_ "cogentcore.org/lab/yaegilab"
)

//go:embed content
var econtent embed.FS

//go:embed icon.svg
var icon string

func main() {
	core.AppIcon = icon
	b := core.NewBody("CompCogNeuro")
	ct := content.NewContent(b).SetContent(econtent)
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}
