package migrationsSQL

import (
	"embed"
)

// Todos arquivos da pasta atual serão embutidos na variavel abaixo.
//
//go:embed *.sql
var MigrationsDir embed.FS
