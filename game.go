package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GSTATE_PLAYER_INPUT = iota
	GSTATE_PROJECTILE_MOVEMENT
	GSTATE_MOBS_DECISION
	GSTATE_MOBS_ACTION
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
		x: 5.5,
		y: 5.5,
	}
	g.scene.init(g.player.x, g.player.y)
	gameIsRunning = true
}

func (g *game) gameLoop() {
	for gameIsRunning && !rl.WindowShouldClose() {
		g.workPlayerInput()
		g.actProjectiles()
		g.decideMobs()
		g.actMobs()
		g.actTiles()
		tick++
		renderFrame(g.scene)
	}
}

func (g *game) actProjectiles() {
	speed := 0.75
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		switch node.Value.(type) {
		case *projectile:
			proj := node.Value.(*projectile)
			newX := proj.x + (proj.dirX * speed)
			newY := proj.y + (proj.dirY * speed)
			if !g.scene.areRealCoordsPassable(newX, newY) {
				g.scene.things.Remove(node)
			} else {
				proj.x, proj.y = newX, newY
			}
		}
	}
}

func (g *game) actTiles() {
	for x := range g.scene.gameMap {
		for y := range g.scene.gameMap[x] {
			g.scene.gameMap[x][y].actOnState()
		}
	}
}

func (g *game) decideMobs() {

}

func (g *game) actMobs() {

}
