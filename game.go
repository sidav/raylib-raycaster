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
		}
		renderFrame(g.scene)
	}
}

func (g *game) actProjectiles() {
	for _, proj := range g.scene.projectiles {
		proj.x += proj.dirX
		proj.y += proj.dirY
	}
}
