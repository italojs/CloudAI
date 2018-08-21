package domain

import (
	"image"
	"image/jpeg"
	"os"
)

func createFile(filePath string, img image.Image) (err error) {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	jpeg.Encode(f, img, nil)
	return
}

func createPath(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return
}

func verifyPath(path string) (exist bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}
	exist = true
	return
}

func WitheOnDisk(folderPath string, img image.Image, imgName string) (err error) {
	if err != nil {
		return err
	}

	if !verifyPath(folderPath) {
		err = createPath(folderPath)
		if err != nil {
			return err
		}
	}

	filePath := folderPath + "/" + imgName + ".jpg"

	if !verifyPath(filePath) {
		err = createFile(filePath, img)
		if err != nil {
			return err
		}
	}
	return
}
