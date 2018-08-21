package helper

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"strings"
)

func Decode(base64Str string) (imageDecoded image.Image, err error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Str))
	imageDecoded, _, err = image.Decode(reader)
	if err != nil {
		return
	}
	return
}

func Encode(img image.Image) (base64str string, err error) {
	// Here we allocate and create a buffer that can easily be
	// turned into a []byte
	out := new(bytes.Buffer)

	// We now encode the image we created to the buffer
	err = png.Encode(out, img)
	if err != nil {
		return
	}

	// This now takes a []byte of the buffer and base64 encodes it to a string
	// Never needing to create the image file all done in memory
	base64str = base64.StdEncoding.EncodeToString(out.Bytes())
	return
}
