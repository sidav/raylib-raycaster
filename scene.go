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

func (s *Scene) AreGridCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(s.gameMap) && y < len(s.gameMap[0])
}

func (s *Scene) GetMobAtRealCoords(x, y float64) *mob {
	tx, ty := trueCoordsToTileCoords(x, y)
	return s.GetMobAtTileCoords(tx, ty)
}

func (s *Scene) removeMob(mb *mob) {
	for m := s.things.Front(); m != nil; m = m.Next() {
		switch m.Value.(type) {
		case *mob:
			if m.Value.(*mob) == mb {
				s.things.Remove(m)
				return
			}
		}
	}
}

func (s *Scene) GetMobAtTileCoords(tx, ty int) *mob {
	for m := s.things.Front(); m != nil; m = m.Next() {
		switch m.Value.(type) {
		case *mob:
			mrx, mry, _ := m.Value.(*mob).GetCoords()
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
