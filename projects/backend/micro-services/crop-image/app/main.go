package main

import (
	"os"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api"
	"github.com/unrolled/render"
)

const local string = "LOCAL"

func main() {
	// environment variables
	appConfig := make(map[string]string)
	appConfig["ENV"] = os.Getenv("ENV")                       // LOCAL, DEV, STG, PRD
	appConfig["PORT"] = os.Getenv("PORT")                     // server traffic on this port
	appConfig["VERSION"] = os.Getenv("VERSION")               // path to VERSION file
	appConfig["MONGO_SERVER"] = os.Getenv("MONGO_SERVER")     // MongoDB server DNS
	appConfig["MONGO_DATABASE"] = os.Getenv("MONGO_DATABASE") // Database name

	if appConfig["ENV"] == "" || appConfig["LOCAL"] == local {
		// running from localhost, so set some default values
		appConfig["ENV"] = local
		appConfig["PORT"] = "3001"
		appConfig["VERSION"] = "0.2.0"
	}

	// initialse application context
	ctx := api.AppConfig{
		Render: render.New(),
		Env:    appConfig,
	}
	// start application
	StartServer(ctx)
}
