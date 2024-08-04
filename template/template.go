package template

import (
	"embed"
	_ "embed"
)

//go:embed *.html
var Template embed.FS
