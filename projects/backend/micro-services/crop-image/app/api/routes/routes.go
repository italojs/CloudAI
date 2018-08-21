package route

import (
	"net/http"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api"
	controller "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api/controllers"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type ControllerFunc func(http.ResponseWriter, *http.Request, api.AppConfig)

// Route is the model for the router setup
type Route struct {
	Name           string
	Method         string
	Pattern        string
	ControllerFunc ControllerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var AppRoutes = Routes{
	//=== IMAGE ===
	Route{"Create", "POST", "/image", controller.Create},
	Route{"Ping", "GET", "/ping", controller.Ping},
}
