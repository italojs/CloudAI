package main

import (
	"log"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api"
	route "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api/routes"
	"github.com/unrolled/secure"
)

// StartServer Wraps the mux Router and uses the Negroni Middleware
func StartServer(ctx api.AppConfig) {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range route.AppRoutes {
		httpHandler := api.MakeHandler(ctx, route.ControllerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(httpHandler)
	}
	// security
	var isDevelopment = false
	if ctx.GetEnv("ENV") == local {
		isDevelopment = true
	}
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      isDevelopment, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
		AllowedHosts:       []string{},    // AllowedHosts is a list of fully qualified domain names that are allowed (CORS)
		ContentTypeNosniff: true,          // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:   true,          // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
	})
	// start now
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)
	log.Println("===> Starting app (v" + ctx.GetEnv("VERSION") + ") on port " + ctx.GetEnv("PORT") + " in " + ctx.GetEnv("ENV") + " mode.")
	n.Run(":" + ctx.GetEnv("PORT"))
}
