package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *game) workPlayerInput() {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.movePlayerByFacing(false)
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.rotatePlayer(false)
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.rotatePlayer(true)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		g.movePlayerByFacing(true)
	}
	if rl.IsKeyDown(rl.KeyQ) {
		g.movePlayerSideways(false)
	}
	if rl.IsKeyDown(rl.KeyE) {
		g.movePlayerSideways(true)
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.tryOpenDoorAsPlayer()
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.scene.things.PushBack(&projectile{
			x:          g.scene.Camera.X,
			y:          g.scene.Camera.Y,
			dirX:       g.player.dirX,
			dirY:       g.player.dirY,
			spriteCode: "proj",
		})
		g.gameState++
	}
}

func (g *game) movePlayerByFacing(backwards bool) {
	const MOVEFRAMES = 15.0
	factor := 1.0
	if backwards {
		factor = -1.0
	}
	tx, ty := trueCoordsToTileCoords(g.player.x+factor*g.player.dirX, g.player.y+factor*g.player.dirY)
	if g.scene.IsTilePassable(tx, ty) {
		g.player.x += factor * g.player.dirX
		g.player.y += factor * g.player.dirY
		for i := 0; i < int(MOVEFRAMES)-1; i++ {
			g.scene.Camera.MoveForward(factor * 1 / MOVEFRAMES)
			renderFrame(g.scene)
		}
		g.scene.Camera.X = g.player.x
		g.scene.Camera.Y = g.player.y
	}
}

func (g *game) movePlayerSideways(right bool) {
	const MOVEFRAMES = 15.0
	factor := 1.0
	if right {
		factor = -1.0
	}
	moveDirX, moveDirY := factor*g.player.dirY, -factor*g.player.dirX
	tx, ty := trueCoordsToTileCoords(g.player.x+moveDirX, g.player.y+moveDirY)
	if g.scene.IsTilePassable(tx, ty) {
		g.player.x += moveDirX
		g.player.y += moveDirY
		for i := 0; i < int(MOVEFRAMES)-1; i++ {
			g.scene.Camera.MoveByVector(moveDirX*1/MOVEFRAMES, moveDirY*1/MOVEFRAMES)
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
		g.player.dirX, g.player.dirY = -g.player.dirY, g.player.dirX
		factor = 1.0
	} else {
		g.player.dirX, g.player.dirY = g.player.dirY, -g.player.dirX
	}
	for i := 0; i < int(MOVEFRAMES); i++ {
		g.scene.Camera.Rotate(factor * (90 / MOVEFRAMES) * 3.14159265358 / 180.0)
		renderFrame(g.scene)
	}
}

func (g *game) tryOpenDoorAsPlayer() {
	const MOVEFRAMES = 15.0
	tx, ty := trueCoordsToTileCoords(g.player.x+g.player.dirX, g.player.y+g.player.dirY)
	if g.scene.gameMap[tx][ty].getStaticData().openable {
		if g.scene.gameMap[tx][ty].isOpened() {
			for !g.scene.gameMap[tx][ty].isClosed() {
				g.scene.gameMap[tx][ty].tileSlideAmount -= 1 / MOVEFRAMES
				renderFrame(g.scene)
			}
		} else {
			for !g.scene.gameMap[tx][ty].isOpened() {
				g.scene.gameMap[tx][ty].tileSlideAmount += 1 / MOVEFRAMES
				renderFrame(g.scene)
			}
		}
		// g.scene.gameMap[tx][ty].tileSlideAmount = math.Round(g.scene.gameMap[tx][ty].tileSlideAmount)
	} else if !g.scene.IsTilePassable(tx, ty) { // zoom effect for "pushing" the wall
		initialAngle := VIEW_ANGLE / 2.0
		g.scene.Camera.ChangeViewWidth(initialAngle)
		angleIncrement := (VIEW_ANGLE - initialAngle) / MOVEFRAMES

		for i := 0; i < MOVEFRAMES; i++ {
			g.scene.Camera.ChangeViewWidth(initialAngle + float64(i)*angleIncrement)
			renderFrame(g.scene)
		}
	}
	g.gameState++
}
