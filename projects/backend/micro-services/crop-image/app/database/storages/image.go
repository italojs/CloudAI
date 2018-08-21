package storage

import (
	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/dao"
	mgo "gopkg.in/mgo.v2"
)

const COLLECTION = "image"

type Image struct {
	DB *mgo.Database
}

func (img *Image) Save(imageDao dao.Image) (err error) {
	err = img.DB.C(COLLECTION).Insert(&imageDao)
	return
}
