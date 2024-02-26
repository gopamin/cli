package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var webAppEnv []byte

func WebAppEnvTemplate() ([]byte, string) {
	return webAppEnv, ".env"
}

//go:embed files/makefile.tmpl
var webAppMakefile []byte

func WebAppMakefileTemplate() ([]byte, string) {
	return webAppMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var webAppReadme []byte

func WebAppReadmeTemplate() ([]byte, string) {
	return webAppReadme, "README.md"
}

//go:embed files/internal/adapters/handlers/http/routes.tmpl
var webAppHttpRoutes []byte

func WebAppHttpRoutesTemplate() ([]byte, string) {
	return webAppHttpRoutes, "internal/adapters/handlers/http/routes.go"
}

//go:embed files/internal/adapters/handlers/http/users.tmpl
var webAppHttpUsers []byte

func WebAppHttpUsersTemplate() ([]byte, string) {
	return webAppHttpUsers, "internal/adapters/handlers/http/users.go"
}

//go:embed files/internal/adapters/handlers/chi/routes.tmpl
var webAppChiRoutes []byte

func WebAppChiRoutesTemplate() ([]byte, string) {
	return webAppChiRoutes, "internal/adapters/handlers/chi/routes.go"
}

//go:embed files/internal/adapters/handlers/chi/users.tmpl
var webAppHttpChiUsers []byte

func WebAppChiUsersTemplate() ([]byte, string) {
	return webAppHttpChiUsers, "internal/adapters/handlers/chi/users.go"
}

//go:embed files/internal/adapters/handlers/echo/routes.tmpl
var webAppEchoRoutes []byte

func WebAppEchoRoutesTemplate() ([]byte, string) {
	return webAppEchoRoutes, "internal/adapters/handlers/echo/routes.go"
}

//go:embed files/internal/adapters/handlers/echo/users.tmpl
var webAppHttpEchoUsers []byte

func WebAppEchoUsersTemplate() ([]byte, string) {
	return webAppHttpEchoUsers, "internal/adapters/handlers/echo/users.go"
}

//go:embed files/internal/adapters/handlers/gin/routes.tmpl
var webAppGinRoutes []byte

func WebAppGinRoutesTemplate() ([]byte, string) {
	return webAppGinRoutes, "internal/adapters/handlers/gin/routes.go"
}

//go:embed files/internal/adapters/handlers/gin/users.tmpl
var webAppHttpGinUsers []byte

func WebAppGinUsersTemplate() ([]byte, string) {
	return webAppHttpGinUsers, "internal/adapters/handlers/gin/users.go"
}

//go:embed files/internal/adapters/handlers/gorilla/routes.tmpl
var webAppGorillaRoutes []byte

func WebAppGorillaRoutesTemplate() ([]byte, string) {
	return webAppGorillaRoutes, "internal/adapters/handlers/gorilla/routes.go"
}

//go:embed files/internal/adapters/handlers/gorilla/users.tmpl
var webAppHttpGorillaUsers []byte

func WebAppGorillaUsersTemplate() ([]byte, string) {
	return webAppHttpGorillaUsers, "internal/adapters/handlers/gorilla/users.go"
}

//go:embed files/internal/adapters/handlers/httprouter/routes.tmpl
var webAppHttprouterRoutes []byte

func WebAppHttprouterRoutesTemplate() ([]byte, string) {
	return webAppHttprouterRoutes, "internal/adapters/handlers/httprouter/routes.go"
}

//go:embed files/internal/adapters/handlers/httprouter/users.tmpl
var webAppHttpHttprouterUsers []byte

func WebAppHttprouterUsersTemplate() ([]byte, string) {
	return webAppHttpHttprouterUsers, "internal/adapters/handlers/httprouter/users.go"
}
