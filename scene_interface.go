package main

import "raylib-raycaster/raycaster"

func (s *Scene) IsTileOpaque(x, y int) bool {
	return s.gameMap[x][y].getStaticData().opaque
}

func (s *Scene) GetTileElevation(x, y int) float64 {
	return 0
}

func (s *Scene) GetCamera() *raycaster.Camera {
	return s.Camera
}

func (s *Scene) GetTileSlideAmount(x, y int) float64 {
	return s.gameMap[x][y].tileSlideAmount
}

func (s *Scene) IsTileThin(x, y int) bool {
	return s.gameMap[x][y].getStaticData().thin
}

func (s *Scene) GetTextureForTile(x, y int) *raycaster.Texture {
	t := s.gameMap[x][y]
	tex := wallTexturesAtlas[t.tileCode]
	if tex == nil {
		tex = wallTexturesAtlas["WALL"]
	}
	if tex == nil {
		panic("NO Texture FOR " + t.tileCode)
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
