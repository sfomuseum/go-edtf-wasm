package www

import (
	"embed"
)

//go:embed *.html wasm/* css/* javascript/*
var FS embed.FS
