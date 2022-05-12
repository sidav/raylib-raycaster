package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *game) workPlayerInput() {
	if rl.IsKeyDown(rl.KeyUp) {
		g.movePlayerByFacing(false)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.rotatePlayer(false)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.rotatePlayer(true)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.movePlayerByFacing(true)
	}
	//if rl.IsKeyPressed(rl.KeySpace) {
	//	g.scene.things = append(g.scene.things, &thing{
	//		x:          g.scene.Camera.X,
	//		y:          g.scene.Camera.Y,
	//		spriteCode: "proj",
	//	})
	//}
}

func (g *game) movePlayerByFacing(backwards bool) {
	const MOVEFRAMES = 15.0
	factor := 1.0
	if backwards {
		factor = -1.0
	}
	tx, ty := trueCoordsToTileCoords(g.player.x+factor*g.player.facex, g.player.y+factor*g.player.facey)
	if g.scene.IsTilePassable(tx, ty) {
		g.player.x += factor*g.player.facex
		g.player.y += factor*g.player.facey
		for i := 0; i < int(MOVEFRAMES)-1; i++ {
			g.scene.Camera.MoveForward(factor*1/MOVEFRAMES)
			renderFrame(g.scene)
		}
		g.scene.Camera.X = g.player.x
		g.scene.Camera.Y = g.player.y
	}
}

func (g *game) rotatePlayer(clockwise bool) {
	const MOVEFRAMES = 15.0
	factor := -1.0
	if clockwise {
		g.player.facex, g.player.facey = -g.player.facey, g.player.facex
		factor = 1.0
	} else {
		g.player.facex, g.player.facey = g.player.facey, -g.player.facex
	}
	for i := 0; i < int(MOVEFRAMES); i++ {
		g.scene.Camera.Rotate(factor*(90/MOVEFRAMES)*3.14159265358 / 180.0)
		renderFrame(g.scene)
	}
}
