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
	pages.NewPage(b).SetContent(content)
	b.RunMainWindow()
}
