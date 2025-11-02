package static

import "embed"

var (
	//go:embed www/*
	Files embed.FS
)
