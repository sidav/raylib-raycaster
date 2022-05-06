package main

import (
	"image"
	"image/png"
	"os"
	"raylib-raycaster/raycaster"
)

type Scene struct {
	wallTexturesAtlas map[rune]*raycaster.Texture
	gameMap           [][]rune
	Camera            *raycaster.Camera
}

func (s *Scene) init() {
	s.Camera = &raycaster.Camera{
		X: 5,
		Y: 5,
	}
	mp := []string{
		"##########",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"##########",
	}
	s.gameMap = make([][]rune, 0)
	for i := 0; i < len(mp); i++ {
		s.gameMap = append(s.gameMap, make([]rune, 0))
		for j := 0; j < len(mp[i]); j++ {
			s.gameMap[i] = append(s.gameMap[i], rune(mp[i][j]))
		}
	}

	// init textures (temp.)
	s.wallTexturesAtlas = make(map[rune]*raycaster.Texture, 0)
	s.wallTexturesAtlas['#'] = s.readTextureFromFile("textures/wall.png")
}

func (s *Scene) AreGridCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(s.gameMap) && y < len(s.gameMap[0])
}

func (s *Scene) IsTileOpaque(x, y int) bool {
	return s.gameMap[x][y] != '.'
}

func (s *Scene) GetTileElevation(x, y int) float64 {
	return 0
}

func (s *Scene) GetCamera() *raycaster.Camera {
	return s.Camera
}

func (s *Scene) GetTileSlideAmount(x, y int) float64 {
	//if s.gameMap[x][y] == '-' || s.gameMap[x][y] == '|' {
	//	return s.CurrElevation
	//}
	return 0
}

func (s *Scene) IsTileThin(x, y int) bool {
	return s.gameMap[x][y] == '|' || s.gameMap[x][y] == '-' || s.gameMap[x][y] == '+'
}

func (s *Scene) GetTextureForTile(x, y int) *raycaster.Texture {
	t := s.gameMap[x][y]
	tex := s.wallTexturesAtlas[t]
	if tex == nil {
		tex = s.wallTexturesAtlas['#']
	}
	if tex == nil {
		panic("NO Texture FOR " + string(t))
	}
	return tex
}

func (s *Scene) GetFloorTextureForCoords(x, y int) *raycaster.Texture {
	return nil
	//t := '#'
	//if s.AreGridCoordsValid(x, y) {
	//	t = s.gameMap[x][y]
	//}
	//tex := floorTexture // TextureAtlas[t]
	//if tex == nil {
	//	tex = TextureAtlas['#']
	//}
	//if tex == nil {
	//	panic("NO Texture FOR " + string(t))
	//}
	//return tex
}

func (s *Scene) GetCeilingTextureForCoords(x, y int) *raycaster.Texture {
	return nil
	//t := '#'
	//if s.AreGridCoordsValid(x, y) {
	//	t = s.gameMap[x][y]
	//}
	//tex := ceilingTexture // TextureAtlas[t]
	//if tex == nil {
	//	tex = TextureAtlas['#']
	//}
	//if tex == nil {
	//	panic("NO Texture FOR " + string(t))
	//}
	//return tex
}

func (s *Scene) GetListOfThings() []*raycaster.Thing {
	return make([]*raycaster.Thing, 0)
}

func (s *Scene) readTextureFromFile(filename string) *raycaster.Texture {
	// first, read the image file
	imgfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()
	img, err := png.Decode(imgfile)
	imgfile.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	// return the "reader carriage" to zero
	imgfile.Seek(0, 0)
	// read file config (needed for w, h)
	cfg, _, err := image.DecodeConfig(imgfile)
	if err != nil {
		panic(err)
	}
	// init the map if needed

	return &raycaster.Texture{
		Bitmap: img,
		W:      cfg.Width,
		H:      cfg.Height,
	}
}
