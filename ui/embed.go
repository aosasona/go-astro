package ui

import (
	"embed"
)

//go:generate yarn
//go:generate yarn build
//go:embed dist/*
var UI embed.FS
