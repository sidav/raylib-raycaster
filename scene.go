package main

import (
	"container/list"
	"raylib-raycaster/raycaster"
)

type Scene struct {
	gameMap [][]tile
	things  *list.List
	Camera  *raycaster.Camera
}

func (s *Scene) init(camX, camY float64) {
	s.Camera = raycaster.CreateCamera(camX, camY, VIEW_ANGLE, 0, 0, 4, 1)
	s.things = list.New()
	mp := []string{
		"#############################################",
		"#             #       #       #             #",
		"#             #       +       #             #",
		"#             #       #       #             #",
		"#        #    #       #       #             #",
		"#             #       #       #             #",
		"#             #       #       +             #",
		"#####+#########       #       #             #",
		"#            #        #       #             #",
		"#            +        #       #             #",
		"#            #        #       #             #",
		"#            #        #       #             #",
		"###+###################       ######+########",
		"#                             #             #",
		"#                             #             #",
		"#                             #             #",
		"#                             #             #",
		"#                             #             #",
		"#                             #             #",
		"#############################################",
	}
	s.gameMap = make([][]tile, 0)
	for i := 0; i < len(mp); i++ {
		s.gameMap = append(s.gameMap, make([]tile, 0))
		for j := 0; j < len(mp[i]); j++ {
			char := rune(mp[i][j])
			var code string
			switch char {
			case '.', ' ':
				code = "FLOOR"
			case '#':
				code = "WALL"
			case '+':
				code = "DOOR"
			}
			s.gameMap[i] = append(s.gameMap[i], tile{tileCode: code})
		}
	}

	for i := 0; i < 15; i++ {
		x, y := 0, 0
		for !s.IsTilePassable(x, y) {
			x = rnd.Intn(len(s.gameMap))
			y = rnd.Intn(len(s.gameMap[0]))
		}
		rx, ry := tileCoordsToTrueCoords(x, y)
		s.things.PushBack(&mob{
			x:          rx,
			y:          ry,
			dirX:       1,
			dirY:       0,
			spriteCode: "enemy",
		})
	}
	s.things.PushBack(&mob{
		x:          4.5,
		y:          4.5,
		dirX:       1,
		dirY:       0,
		spriteCode: "enemy",
	})
}

func (s *Scene) AreGridCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(s.gameMap) && y < len(s.gameMap[0])
}

func (s *Scene) GetMobAtRealCoords(x, y float64) *mob {
	tx, ty := trueCoordsToTileCoords(x, y)
	return s.GetMobAtTileCoords(tx, ty)
}

func (s *Scene) GetMobAtTileCoords(tx, ty int) *mob {
	for m := s.things.Front(); m != nil; m = m.Next() {
		switch m.Value.(type) {
		case *mob:
			mrx, mry := m.Value.(*mob).GetCoords()
			mtx, mty := trueCoordsToTileCoords(mrx, mry)
			if mtx == tx && mty == ty {
				return m.Value.(*mob)
			}
		}
	}
	return nil
}

func (s *Scene) areRealCoordsPassable(x, y float64) bool {
	tx, ty := trueCoordsToTileCoords(x, y)
	return s.gameMap[tx][ty].isPassable()
}

func (s *Scene) IsTilePassable(x, y int) bool {
	return s.gameMap[x][y].isPassable()
}
