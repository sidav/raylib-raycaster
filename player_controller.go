package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	playerMovementSpeed = 0.05
	playerRotationSpeed = 3.5 // degrees
)

func (g *game) workPlayerInput() {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.movePlayerByFacing(false)
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyQ) {
		g.rotatePlayer(false)
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyE) {
		g.rotatePlayer(true)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		g.movePlayerByFacing(true)
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.movePlayerSideways(false)
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.movePlayerSideways(true)
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.tryOpenDoorAsPlayer()
	}
	if rl.IsKeyDown(rl.KeyEnter) {
		dx, dy := g.player.GetDirectionVector()
		g.scene.things.PushBack(&projectile{
			x:          g.scene.Camera.X,
			y:          g.scene.Camera.Y,
			dirX:       dx,
			dirY:       dy,
			spriteCode: "proj",
		})
		g.gameState++
	}
	if rl.IsKeyDown(rl.KeyPageUp) {
		g.scene.Camera.OnScreenVerticalOffset--
	}
	if rl.IsKeyDown(rl.KeyPageDown) {
		g.scene.Camera.OnScreenVerticalOffset++
	}
}

func (g *game) movePlayerByFacing(backwards bool) {
	factor := playerMovementSpeed
	if backwards {
		factor = -factor
	}
	dx, dy := g.player.GetDirectionVector()
	tx, ty := trueCoordsToTileCoords(g.player.x+5*factor*dx, g.player.y+5*factor*dy)
	if g.scene.IsTilePassable(tx, ty) && g.scene.GetMobAtTileCoords(tx, ty) == nil {
		g.player.x += factor * dx
		g.player.y += factor * dy
		g.scene.Camera.MoveForward(factor)
		g.scene.Camera.X = g.player.x
		g.scene.Camera.Y = g.player.y
	}
}

func (g *game) movePlayerSideways(right bool) {
	factor := playerMovementSpeed
	if right {
		factor = -factor
	}
	dx, dy := g.player.GetDirectionVector()
	moveDirX, moveDirY := factor*dy, -factor*dx
	tx, ty := trueCoordsToTileCoords(g.player.x+5*moveDirX, g.player.y+5*moveDirY)
	if g.scene.IsTilePassable(tx, ty) {
		g.player.x += moveDirX
		g.player.y += moveDirY
		g.scene.Camera.MoveByVector(moveDirX, moveDirY)
		g.scene.Camera.X = g.player.x
		g.scene.Camera.Y = g.player.y
	}
}

func (g *game) rotatePlayer(clockwise bool) {
	factor := -1.0
	if clockwise {
		factor = 1.0
	}
	rot := factor * playerRotationSpeed * 3.14159265358 / 180.0
	g.player.rotationRadians += rot
	g.scene.Camera.Rotate(rot)
}

func (g *game) tryOpenDoorAsPlayer() {
	dx, dy := g.player.GetDirectionVector()
	tx, ty := trueCoordsToTileCoords(g.player.x+dx, g.player.y+dy)
	if g.scene.gameMap[tx][ty].getStaticData().openable {
		if g.scene.gameMap[tx][ty].isOpened() {
			g.scene.gameMap[tx][ty].state = tileStateClosing
		} else {
			g.scene.gameMap[tx][ty].state = tileStateOpening
		}
	}
}
