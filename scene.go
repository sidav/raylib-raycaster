package main

import (
	"raylib-raycaster/raycaster"
)

type Scene struct {
	gameMap     [][]tile
	projectiles []*projectile
	Camera      *raycaster.Camera
}

func (s *Scene) init(camX, camY float64) {
	s.Camera = raycaster.CreateCamera(camX, camY, VIEW_ANGLE, 0, 0, 4, 1)
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
}

func (s *Scene) AreGridCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(s.gameMap) && y < len(s.gameMap[0])
}

func (s *Scene) IsTilePassable(x, y int) bool {
	return s.gameMap[x][y].isPassable()
}
