package migrations

import "embed"

// Files holds embedded migration SQL files.
//
//go:embed *.sql
var Files embed.FS
