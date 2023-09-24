package main

import (
	"container/list"
	bspdung "raylib-raycaster/lib/bsp_dung"
	"raylib-raycaster/raycaster"
)

func (s *Scene) init() (float64, float64) {
	s.things = list.New()
	gen := bspdung.Generator{
		MinRoomSide:            3,
		RoomWForRandomDropping: 8,
	}
	mp := gen.Generate(rnd, 40, 40)
	s.gameMap = make([][]tile, 0)
	camX, camY := 0.0, 0.0
	for i := 0; i < len(mp); i++ {
		s.gameMap = append(s.gameMap, make([]tile, 0))
		for j := 0; j < len(mp[i]); j++ {
			char := rune(mp[i][j])
			var code string
			switch char {
			case '.', ' ':
				code = "FLOOR"
			case '<':
				camX, camY = float64(i)+0.5, float64(j)+0.5
				code = "FLOOR"
			case '#', '"':
				code = "WALL"
			case '+':
				if rnd.Rand(5) == 0 {
					code = "DOORHORIZ"
				} else {
					code = "DOORVERT"
				}
			}
			s.gameMap[i] = append(s.gameMap[i], tile{tileCode: code})
		}
	}

	s.Camera = raycaster.CreateCamera(camX, camY, VIEW_ANGLE, 4, 1)
	s.placeMobs()
	s.placeItems()
	s.finalizeTiles()
	return camX, camY
}

func (s *Scene) placeMobs() {
	for i := 0; i < 50; i++ {
		x, y := 0, 0
		for !s.IsTilePassable(x, y) || s.GetMobAtTileCoords(x, y) != nil {
			x = rnd.Rand(len(s.gameMap))
			y = rnd.Rand(len(s.gameMap[0]))
		}
		rx, ry := tileCoordsToTrueCoords(x, y)
		s.things.PushBack(createMob(rx, ry, sTableMobs[rnd.Rand(len(sTableMobs))]))
	}
}

func (s *Scene) placeItems() {
	for i := 0; i < 25; i++ {
		x, y := 0, 0
		// TODO: don't place items on items
		for !s.IsTilePassable(x, y) {
			x = rnd.Rand(len(s.gameMap))
			y = rnd.Rand(len(s.gameMap[0]))
		}
		rx, ry := tileCoordsToTrueCoords(x, y)
		s.things.PushBack(createPickupable(rx, ry, sTablePickupables[rnd.Rand(len(sTablePickupables))]))
	}
}

func (s *Scene) finalizeTiles() {
	for x := range s.gameMap {
		for y := range s.gameMap[x] {
			tileCode := s.gameMap[x][y].tileCode
			if tileCode == "DOORHORIZ" || tileCode == "DOORVERT" {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if i*j == 0 && s.AreGridCoordsValid(x+i, y+j) && s.gameMap[x+i][y+j].tileCode == "WALL" {
							s.gameMap[x+i][y+j].tileCode = "WALLELEC"
						}
					}
				}
			}
		}
	}
}
