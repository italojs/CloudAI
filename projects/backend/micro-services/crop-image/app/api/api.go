package api

import (
	"log"
	"net/http"

	"github.com/unrolled/render"
)

// AppConfig holds application configuration data
type AppConfig struct {
	Render *render.Render
	Env    map[string]string
}

func (app AppConfig) GetEnv(env string) string {
	return app.Env[env]
}

// makeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(ctx AppConfig, fn func(http.ResponseWriter, *http.Request, AppConfig)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	}
}

// PassportsHandler not implemented yet
func PassportsHandler(w http.ResponseWriter, req *http.Request, ctx AppConfig) {
	log.Println("Handling Passports - Not implemented yet")
	ctx.Render.Text(w, http.StatusNotImplemented, "")
}
