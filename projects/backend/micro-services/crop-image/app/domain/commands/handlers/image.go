package handler

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"strconv"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/interfaces"
	"github.com/oliamb/cutter"

	"github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/dao"
	entity "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/entities"
	storage "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/database/storages"
	command "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/commands"
	helper "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/helpers"
	io "github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/domain/io"
)

// Image is contain the handler method to handler an image that
// are received in /image controller
type Image struct {
	Storage storage.Image
	Ctx     interfaces.AppConfig
}

// Handle Method receive an image and handler it to crop, write it on disk,
// send this images to a image storage service and save it in database
// TODO: apply depency injection
func (imgHdlr *Image) Handle(imageCmd command.Image) (err error) {
	originalImage, err := helper.Decode(imageCmd.Base64)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	//send it to some storage cloud
	folderClassPath := fmt.Sprintf("/images/%v/%v/", imageCmd.User, imageCmd.Label)
	io.WitheOnDisk(folderClassPath, originalImage, strconv.Itoa(rand.Int()))

	imgDao := dao.Image{}
	imgDao.OriginalImage = entity.Image{
		StorageLink: folderClassPath,
		Label:       imageCmd.Label,
		Length: entity.XY{
			X: imageCmd.Length.X,
			Y: imageCmd.Length.Y,
		},
	}

	for _, classObj := range imageCmd.ClassObjects {
		x := classObj.Crop.To.X - classObj.Crop.From.X
		y := classObj.Crop.To.Y - classObj.Crop.From.Y

		classImage, err := cutter.Crop(originalImage, cutter.Config{
			Width:  imageCmd.Length.X,
			Height: imageCmd.Length.Y,
			Anchor: image.Point{x, y},
			Mode:   cutter.TopLeft,
		})
		if err != nil {
			log.Panic(err)
			panic(err)
		}

		length := entity.XY{
			X: classObj.Crop.To.X - classObj.Crop.From.X,
			Y: classObj.Crop.To.Y - classObj.Crop.From.Y,
		}

		//send it to some storage cloud
		folderClassPath := fmt.Sprintf("/images/%v/%v/crops/%v/", imageCmd.User, imageCmd.Label, classObj.Label)
		io.WitheOnDisk(folderClassPath, classImage, strconv.Itoa(rand.Int()))

		classImgEntity := entity.Image{
			StorageLink: folderClassPath,
			Label:       classObj.Label,
			Length:      length,
		}

		classObjEntity := entity.ClassObject{
			Image: classImgEntity,
			Crop: entity.Crop{
				From: entity.XY{
					X: classObj.Crop.From.X,
					Y: classObj.Crop.From.Y,
				},
				To: entity.XY{
					X: classObj.Crop.To.X,
					Y: classObj.Crop.To.Y,
				},
			},
		}

		imgDao.ClassObjects = append(imgDao.ClassObjects, classObjEntity)

	}

	imgHdlr.Storage.Save(imgDao)
	return

}
