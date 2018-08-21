package controller

import (
	"net/http"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api"

	command "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/commands"
)

func Ping(w http.ResponseWriter, req *http.Request, ctx api.AppConfig) {
	response := command.Status{
		Code:    200,
		Message: "OK",
	}
	ctx.Render.JSON(w, http.StatusCreated, response)
}
