package main

import (
	"container/list"
	bspdung "raylib-raycaster/lib/bsp_dung"
	pcgr "raylib-raycaster/lib/random/pcgrandom"
	"raylib-raycaster/raycaster"
	"time"
)

func (s *Scene) init() (float64, float64) {
	s.things = list.New()
	gen := bspdung.Generator{
		MinRoomSide:            3,
		RoomWForRandomDropping: 8,
	}
	rnd := pcgr.NewPCG64()
	rnd.SetSeed(int(time.Now().UnixNano()))
	// rnd.SetSeed(1)
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
				code = "DOOR"
			}
			s.gameMap[i] = append(s.gameMap[i], tile{tileCode: code})
		}
	}

	s.Camera = raycaster.CreateCamera(camX, camY, VIEW_ANGLE, 4, 1)
	for i := 0; i < 15; i++ {
		x, y := 0, 0
		for !s.IsTilePassable(x, y) {
			x = rnd.Rand(len(s.gameMap))
			y = rnd.Rand(len(s.gameMap[0]))
		}
		rx, ry := tileCoordsToTrueCoords(x, y)
		s.things.PushBack(&mob{
			x:          rx,
			y:          ry,
			spriteCode: "enemy",
		})
	}
	return camX, camY
}
