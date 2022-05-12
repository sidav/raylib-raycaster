package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	GSTATE_PLAYER_INPUT = iota
	GSTATE_PLAYER_MOVEMENT
	GSTATE_WORLD_MOVEMENT
	GSTATE_RESET_TO_ZERO // this state should be reset to 0
)

type game struct {
	gameState int
	scene     *Scene
}

func (g *game) init() {
	g.scene = &Scene{}
	g.scene.init()
	gameIsRunning = true
}

func (g *game) gameLoop() {
	for gameIsRunning && !rl.WindowShouldClose() {
		switch g.gameState {
		case GSTATE_PLAYER_INPUT:
			g.workPlayerInput()
		}
		renderFrame(g.scene)
	}
}
