package main

import "raylib-raycaster/raycaster"

type decoration struct {
	x, y, z           float64
	spriteCode        string
	width, height     float64
	blocksMovement    bool
	blocksProjectiles bool
	remainingLifetime int // -1 means "forever"
}

func (d *decoration) GetCoords() (float64, float64, float64) {
	return d.x, d.y, d.z
}

func (d *decoration) GetWidthAndHeightFactors() (float64, float64) {
	return d.width, d.height
}

func (d *decoration) GetDirectionVector() (float64, float64) {
	return 1, 0 // math.Cos(t.rotationRadians), math.Sin(t.rotationRadians)
}

func (d *decoration) GetSprite() *raycaster.SpriteStruct {
	return spritesAtlas[d.spriteCode][0]
}
