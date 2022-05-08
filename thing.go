package main

import "raylib-raycaster/raycaster"

type thing struct {
	x, y float64
	dirX, dirY int
	spriteCode string
}

func (t *thing) GetCoords() (float64, float64) {
	return t.x, t.y
}

func (t *thing) GetSprite() *raycaster.SpriteStruct {
	return nil
}
