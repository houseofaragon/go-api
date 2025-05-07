package migrations

import "embed"

// find all sql files that exist in directory path
//go:embed *.sql
var FS embed.FS