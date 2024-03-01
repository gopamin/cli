package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiEnv []byte

func ApiEnvTemplate() ([]byte, string) {
	return apiEnv, ".env"
}

//go:embed files/readme/readme.tmpl
var apiReadme []byte

func ApiReadmeTemplate() ([]byte, string) {
	return apiReadme, "README.md"
}

//go:embed files/readme/readme-with-db.tmpl
var apiReadmeWithDb []byte

func ApiReadmeWithDbTemplate() ([]byte, string) {
	return apiReadmeWithDb, "README.md"
}

//go:embed files/cmd/main.tmpl
var apiMain []byte

func ApiMainTemplate() ([]byte, string) {
	return apiMain, "cmd/main.go"
}

//go:embed files/cmd/server/server.tmpl
var apiAppServer []byte

func ApiAppServerTemplate() ([]byte, string) {
	return apiAppServer, "cmd/server/server.go"
}

//go:embed files/cmd/main-with-db.tmpl
var apiMainWithDb []byte

func ApiMainWithDbTemplate() ([]byte, string) {
	return apiMainWithDb, "cmd/main.go"
}

//go:embed files/internal/adapters/handlers/chi/routes.tmpl
var apiChiRoutes []byte

func ApiChiRoutesTemplate() ([]byte, string) {
	return apiChiRoutes, "internal/adapters/handlers/chi/routes.go"
}

//go:embed files/internal/adapters/handlers/chi/users.tmpl
var apiChiUsers []byte

func ApiChiUsersTemplate() ([]byte, string) {
	return apiChiUsers, "internal/adapters/handlers/chi/users.go"
}

//go:embed files/internal/adapters/handlers/echo/routes.tmpl
var apiEchoRoutes []byte

func ApiEchoRoutesTemplate() ([]byte, string) {
	return apiEchoRoutes, "internal/adapters/handlers/echo/routes.go"
}

//go:embed files/internal/adapters/handlers/echo/users.tmpl
var apiEchoUsers []byte

func ApiEchoUsersTemplate() ([]byte, string) {
	return apiEchoUsers, "internal/adapters/handlers/echo/users.go"
}

//go:embed files/internal/adapters/handlers/gin/routes.tmpl
var apiGinRoutes []byte

func ApiGinRoutesTemplate() ([]byte, string) {
	return apiGinRoutes, "internal/adapters/handlers/gin/routes.go"
}

//go:embed files/internal/adapters/handlers/gin/users.tmpl
var apiGinUsers []byte

func ApiGinUsersTemplate() ([]byte, string) {
	return apiGinUsers, "internal/adapters/handlers/gin/users.go"
}

//go:embed files/internal/adapters/handlers/gorilla/routes.tmpl
var apiGorillaRoutes []byte

func ApiGorillaRoutesTemplate() ([]byte, string) {
	return apiGorillaRoutes, "internal/adapters/handlers/gorilla/routes.go"
}

//go:embed files/internal/adapters/handlers/gorilla/users.tmpl
var apiGorillaUsers []byte

func ApiGorillaUsersTemplate() ([]byte, string) {
	return apiGorillaUsers, "internal/adapters/handlers/gorilla/users.go"
}

//go:embed files/internal/adapters/handlers/http/routes.tmpl
var apiHttpRoutes []byte

func ApiHttpRoutesTemplate() ([]byte, string) {
	return apiHttpRoutes, "internal/adapters/handlers/http/routes.go"
}

//go:embed files/internal/adapters/handlers/http/users.tmpl
var apiHttpUsers []byte

func ApiHttpUsersTemplate() ([]byte, string) {
	return apiHttpUsers, "internal/adapters/handlers/http/users.go"
}

//go:embed files/internal/adapters/handlers/httprouter/routes.tmpl
var apiHttprouterRoutes []byte

func ApiHttprouterRoutesTemplate() ([]byte, string) {
	return apiHttprouterRoutes, "internal/adapters/handlers/httprouter/routes.go"
}

//go:embed files/internal/adapters/handlers/httprouter/users.tmpl
var apiHttprouterUsers []byte

func ApiHttprouterUsersTemplate() ([]byte, string) {
	return apiHttprouterUsers, "internal/adapters/handlers/httprouter/users.go"
}
