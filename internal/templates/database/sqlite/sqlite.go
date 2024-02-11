package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var sqliteEnv []byte

func SqliteEnvTemplate() ([]byte, string) {
	return sqliteEnv, ".env"
}

//go:embed files/makefile.tmpl
var sqliteMakefile []byte

func SqliteMakefileTemplate() ([]byte, string) {
	return sqliteMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var sqliteReadme []byte

func SqliteReadmeTemplate() ([]byte, string) {
	return sqliteReadme, "README.md"
}

//go:embed files/internal/adapters/repositories/sqlite/repository.tmpl
var sqliteRepository []byte

func SqliteRepositoryTemplate() ([]byte, string) {
	return sqliteRepository, "internal/adapters/repositories/sqlite/repository.go"
}
