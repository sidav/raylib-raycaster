package main

import (
	"raylib-raycaster/raycaster"
)

type Scene struct {
	gameMap           [][]rune
	things            []raycaster.Thing
	Camera            *raycaster.Camera
}

func (s *Scene) init() {
	s.Camera = raycaster.CreateCamera(5.5, 5.5, 120, 0, 0, 4, 1)
	mp := []string{
		"##########",
		"#........#",
		"#........#",
		"####+#####",
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
	tex := wallTexturesAtlas[t]
	if tex == nil {
		tex = wallTexturesAtlas['#']
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

func (s *Scene) GetListOfThings() []raycaster.Thing {
	return s.things
}
