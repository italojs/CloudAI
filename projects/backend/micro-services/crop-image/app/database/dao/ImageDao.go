package dao

import "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/entities"

type Image struct {
	OriginalImage entity.Image
	ClassObjects  []entity.ClassObject
}
