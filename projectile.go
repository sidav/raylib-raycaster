package main

import (
	"fmt"
	"raylib-raycaster/raycaster"
)

type projectile struct {
	x, y, z    float64
	dirX, dirY float64
	spriteCode string
}

func (t *projectile) GetCoords() (float64, float64, float64) {
	return t.x, t.y, t.z
}

func (t *projectile) GetWidthAndHeightFactors() (float64, float64) {
	return 0.1, 0.1
}

func (t *projectile) GetSprite() *raycaster.SpriteStruct {
	if spritesAtlas[t.spriteCode] == nil {
		panic(fmt.Sprintf("WATAFUQ: %s, %v, %d", t.spriteCode, spritesAtlas, len(spritesAtlas)))
	}
	const changeFrameEveryTicks = 5
	return spritesAtlas[t.spriteCode][(tick/changeFrameEveryTicks)%len(spritesAtlas[t.spriteCode])]
}
