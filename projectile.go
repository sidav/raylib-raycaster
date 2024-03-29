package main

import (
	"raylib-raycaster/raycaster"
)

type projectile struct {
	x, y, z    float64
	dirX, dirY float64
	createdAt  int // tick
	frameNum   int
	static     *projectileStatic
}

func (t *projectile) GetCoords() (float64, float64, float64) {
	return t.x, t.y, t.z
}

func (t *projectile) GetWidthAndHeightFactors() (float64, float64) {
	return t.static.sizeFactor, t.static.sizeFactor
}

func (t *projectile) GetSprite() *raycaster.SpriteStruct {
	return spritesAtlas[t.static.spriteCode][t.frameNum]
}

type codeProjectile uint8

const (
	projectilePlasma codeProjectile = iota
	projectileAcid
	projectileFireball
)

type projectileStatic struct {
	spriteCode            string
	totalFrames           int
	changeFrameEveryTicks int
	sizeFactor            float64
}

var sTableProjectiles = map[codeProjectile]*projectileStatic{
	projectilePlasma: {
		spriteCode:            "projPlasma",
		totalFrames:           2,
		changeFrameEveryTicks: 5,
		sizeFactor:            0.25,
	},
	projectileAcid: {
		spriteCode:  "projAcid",
		totalFrames: 1,
		sizeFactor:  0.75,
	},
	projectileFireball: {
		spriteCode:  "projFireball",
		totalFrames: 1,
		sizeFactor:  0.1,
	},
}
