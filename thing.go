package main

import (
	"fmt"
	"raylib-raycaster/raycaster"
)

type thing struct {
	x, y float64
	dirX, dirY int
	spriteCode string
}

func (t *thing) GetCoords() (float64, float64) {
	return t.x, t.y
}

func (t *thing) GetSprite() *raycaster.SpriteStruct {
	if spritesAtlas[t.spriteCode] == nil {
		panic(fmt.Sprintf("WATAFUQ: %s, %v, %d", t.spriteCode, spritesAtlas, len(spritesAtlas)))
	}
	return spritesAtlas[t.spriteCode]
}
