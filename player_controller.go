package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	playerMovementSpeed = 0.075
	playerRotationSpeed = 5 // degrees
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
	if rl.IsKeyDown(rl.KeyEnter) || rl.IsKeyDown(rl.KeyLeftControl) {
		g.shootAsPlayer()
	}
	if rl.IsKeyDown(rl.KeyPageUp) {
		g.scene.Camera.OnScreenVerticalOffset--
	}
	if rl.IsKeyDown(rl.KeyPageDown) {
		g.scene.Camera.OnScreenVerticalOffset++
	}
	if rl.IsKeyPressed(rl.KeyOne) {
		g.player.weaponInHands = &weapon{static: sTableWeapons[0]}
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		g.player.weaponInHands = &weapon{static: sTableWeapons[1]}
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		g.player.weaponInHands = &weapon{static: sTableWeapons[2]}
	}
	if rl.IsKeyPressed(rl.KeyFour) {
		g.player.weaponInHands = &weapon{static: sTableWeapons[3]}
	}
}

func (g *game) movePlayerByFacing(backwards bool) {
	const checkDistanceFactor = 5.0
	factor := playerMovementSpeed
	if backwards {
		factor = -factor
	}
	dx, dy := g.player.GetDirectionVector()
	checkX, checkY := g.player.x+checkDistanceFactor*factor*dx, g.player.y+checkDistanceFactor*factor*dy
	tx, ty := trueCoordsToTileCoords(checkX, checkY)
	if g.scene.IsTilePassable(tx, ty) && g.scene.GetMobInRadius(checkX, checkY, 0.6) == nil {
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

func (g *game) shootAsPlayer() {
	if g.player.weaponInHands == nil || !g.player.weaponInHands.canShoot() {
		return
	}
	g.player.weaponInHands.lastTickShot = g.currentTick
	g.player.weaponInHands.state = wStateFiring
	weapStatic := g.player.weaponInHands.static
	for i := 0; i < g.player.weaponInHands.static.shotsPerShot; i++ {
		if weapStatic.firesProjectile != nil {
			dx, dy := g.player.GetDirectionVector()
			dx, dy = rotateVectorRandomlyGauss(dx, dy, weapStatic.spreadDegrees)
			g.scene.things.PushBack(
				g.newProjectile(
					g.player.x, g.player.y, g.scene.Camera.GetVerticalCoordWithBob()-0.1,
					dx, dy,
					weapStatic.firesProjectile,
				),
			)
		} else if weapStatic.firesHitscan != nil {
			dx, dy := g.player.GetDirectionVector()
			dx, dy = rotateVectorRandomlyGauss(dx, dy, weapStatic.spreadDegrees)
			hitX, hitY, hitMob := g.scene.traceAttackRay(g.player.x, g.player.y, dx, dy, weapStatic.firesHitscan.maxLength)
			if hitMob != nil {
				fmt.Printf("Hit the %+v\n", hitMob)
				hitMob.hitpoints -= weapStatic.firesHitscan.damage
			}
			g.scene.things.PushBack(&decoration{
				x:                 hitX,
				y:                 hitY,
				remainingLifetime: 3,
				spriteCode:        weapStatic.firesHitscan.hitDecorationSpriteCode,
				width:             0.1,
				height:            0.1,
				blocksMovement:    false,
				blocksProjectiles: false,
			})
		} else {
			panic("Wat")
		}
	}
}
