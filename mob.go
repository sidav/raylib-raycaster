package main

import (
	"fmt"
	"raylib-raycaster/raycaster"
)

type mob struct {
	x, y       float64
	dirX, dirY float64
	intent     *mobIntent
	spriteCode string
}

func (t *mob) GetCoords() (float64, float64) {
	return t.x, t.y
}

func (t *mob) GetSprite() *raycaster.SpriteStruct {
	if spritesAtlas[t.spriteCode] == nil {
		panic(fmt.Sprintf("WATAFUQ: %s, %v, %d", t.spriteCode, spritesAtlas, len(spritesAtlas)))
	}
	const changeFrameEveryTicks = 15
	return spritesAtlas[t.spriteCode][(tick/changeFrameEveryTicks) % len(spritesAtlas[t.spriteCode])]
}
