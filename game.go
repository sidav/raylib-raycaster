package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	GSTATE_PLAYER_INPUT = iota
	GSTATE_PROJECTILE_MOVEMENT
	GSTATE_WORLD_MOVEMENT
	GSTATE_RESET_TO_ZERO // this state should be reset to 0
)

type game struct {
	gameState int
	player    *mob
	scene     *Scene
}

func (g *game) init() {
	g.scene = &Scene{}
	g.player = &mob{
		x:     5.5,
		y:     5.5,
		facey: 1,
	}
	g.scene.init(g.player.x, g.player.y)
	gameIsRunning = true
}

func (g *game) gameLoop() {
	for gameIsRunning && !rl.WindowShouldClose() {
		switch g.gameState {
		case GSTATE_PLAYER_INPUT:
			g.workPlayerInput()
		case GSTATE_PROJECTILE_MOVEMENT:
			g.actProjectiles()
		case GSTATE_RESET_TO_ZERO:
			g.gameState = 0
		default:
			g.gameState++
		}
		renderFrame(g.scene)
	}
}

func (g *game) actProjectiles() {
	changeState := true
	const factor = 5
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		switch node.Value.(type) {
		case *projectile:
			proj := node.Value.(*projectile)
			newX := proj.x + (proj.dirX/factor)
			newY := proj.y + (proj.dirY/factor)
			tx, ty := trueCoordsToTileCoords(newX, newY)
			if !g.scene.gameMap[tx][ty].isPassable() {
				g.scene.things.Remove(node)
			} else {
				changeState = false
				proj.x, proj.y = newX, newY
			}
		}
	}
	if changeState {
		g.gameState++
	}
}
