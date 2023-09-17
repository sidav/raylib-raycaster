package main

import (
	"container/list"
	"fmt"
	"raylib-raycaster/raycaster"
)

func (s *Scene) IsTileOpaque(x, y int) bool {
	return s.gameMap[x][y].getStaticData().opaque
}

func (s *Scene) GetTileVerticalSlide(x, y int) float64 {
	return 0
}

func (s *Scene) GetCamera() *raycaster.Camera {
	return s.Camera
}

func (s *Scene) GetTileHorizontalSlide(x, y int) float64 {
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
	if tex == nil || len(tex) == 0 {
		panic("NO Texture FOR " + t.tileCode)
	}
	const changeFrameEveryTicks = 30
	return tex[tick/changeFrameEveryTicks%len(tex)]
}

func (s *Scene) GetFloorTextureForCoords(x, y int) *raycaster.Texture {
	var tex []*raycaster.Texture
	if !s.AreGridCoordsValid(x, y) {
		tex = floorTexturesAtlas["DEFAULT"]
	} else {
		t := s.gameMap[x][y]
		tex = floorTexturesAtlas[t.tileCode]
		// fmt.Printf("x,y %d,%d; tex is %v, code %s\n", x, y, tex, t.tileCode)
		if tex == nil || len(tex) == 0 {
			tex = floorTexturesAtlas["DEFAULT"]
		}
		if tex == nil || len(tex) == 0 {
			panic(fmt.Sprintf("tex is %v, NO Texture FOR %s", tex, t.tileCode))
		}
	}
	const changeFrameEveryTicks = 30
	return tex[tick/changeFrameEveryTicks%len(tex)]
}

func (s *Scene) GetCeilingTextureForCoords(x, y int) *raycaster.Texture {
	var tex []*raycaster.Texture
	if !s.AreGridCoordsValid(x, y) {
		tex = ceilingTexturesAtlas["DEFAULT"]
	} else {
		t := s.gameMap[x][y]
		tex = ceilingTexturesAtlas[t.tileCode]
		// fmt.Printf("x,y %d,%d; tex is %v, code %s\n", x, y, tex, t.tileCode)
		if tex == nil || len(tex) == 0 {
			tex = ceilingTexturesAtlas["DEFAULT"]
		}
		if tex == nil || len(tex) == 0 {
			panic(fmt.Sprintf("tex is %v, NO Texture FOR %s", tex, t.tileCode))
		}
	}
	const changeFrameEveryTicks = 30
	return tex[tick/changeFrameEveryTicks%len(tex)]
}

func (s *Scene) GetListOfThings() *list.List {
	return s.things
}
