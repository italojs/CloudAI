package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/api"
	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database"
	storage "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/storages"

	command "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/commands"
	handler "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/commands/handlers"
)

func Create(w http.ResponseWriter, req *http.Request, ctx api.AppConfig) {
	body := json.NewDecoder(req.Body)

	var imageCmd command.Image
	err := body.Decode(&imageCmd)

	conn := database.GetConnection(ctx.GetEnv("MONGO_SERVER"), ctx.GetEnv("MONGO_DATABASE"))
	imgStorage := storage.Image{
		conn.DB,
	}

	imgHldr := handler.Image{
		imgStorage,
		ctx,
	}

	err = imgHldr.Handle(imageCmd)
	if err != nil {
		response := command.Status{
			Code:    400,
			Message: "malformed user object",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	response := command.Status{
		Code:    200,
		Message: "OK",
	}
	ctx.Render.JSON(w, http.StatusCreated, response)
}
