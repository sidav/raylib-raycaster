package raycaster

import (
	"image"
	"os"
)

type SpriteStruct struct {
	bitmap image.Image
	w, h   int
}

func InitSpriteFromImageFile(filename string) *SpriteStruct {
	s := SpriteStruct{}
	// first, read the image file
	imgfile, _ := os.Open(filename)
	defer imgfile.Close()
	img, _, err := image.Decode(imgfile)
	if err != nil {
		panic(err)
	}
	// return the "reader carriage" to zero
	imgfile.Seek(0, 0)
	// read file config (needed for W, H)
	cfg, _, err := image.DecodeConfig(imgfile)
	if err != nil {
		panic(err)
	}

	s.bitmap = img
	s.w, s.h = cfg.Width, cfg.Height
	return &s
}
