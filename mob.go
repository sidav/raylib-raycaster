package main

import (
	"fmt"
	"raylib-raycaster/raycaster"
)

type mob struct {
	x, y         float64
	facex, facey float64
	spriteCode   string
}

func (t *mob) GetCoords() (float64, float64) {
	return t.x, t.y
}

func (t *mob) GetSprite() *raycaster.SpriteStruct {
	if spritesAtlas[t.spriteCode] == nil {
		panic(fmt.Sprintf("WATAFUQ: %s, %v, %d", t.spriteCode, spritesAtlas, len(spritesAtlas)))
	}
	return spritesAtlas[t.spriteCode]
}

