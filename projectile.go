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

func (g *game) newProjectile(x, y, z, dirX, dirY float64, s *projectileStatic) *projectile {
	return &projectile{
		x:         x,
		y:         y,
		z:         z,
		dirX:      dirX,
		dirY:      dirY,
		createdAt: g.currentTick,
		static:    s,
	}
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

type projectileStatic struct {
	spriteCode            string
	totalFrames           int
	changeFrameEveryTicks int
	damage                int
	speed                 float64
	sizeFactor            float64
}
