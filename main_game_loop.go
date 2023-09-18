package main

import (
	"fmt"

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
	gameState   int
	player      *mob
	scene       *Scene
	currentTick int
}

func (g *game) init() {
	g.scene = &Scene{}
	px, py := g.scene.init()
	g.player = &mob{
		x: px,
		y: py,
		weaponInHands: &weapon{
			static: sTableWeapons[0],
		},
	}
	gameIsRunning = true
}

func (g *game) gameLoop() {
	for gameIsRunning && !rl.WindowShouldClose() {
		g.workPlayerInput()
		g.actProjectiles()
		g.decideMobs()
		g.actMobs()
		g.actTiles()
		g.updatePlayerWeaponState()
		tick++
		g.currentTick++
		fmt.Printf("TOTAL %d THINGS\n", g.scene.things.Len())
		renderFrame(g)
	}
}

func (g *game) actProjectiles() {
	speed := 0.5
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

func (g *game) updatePlayerWeaponState() {
	wpn := g.player.weaponInHands
	if wpn == nil {
		return
	}
	if g.currentTick-wpn.lastTickShot > wpn.static.ticksInFiringState {
		wpn.state = wStateIdle
	}
}

func (g *game) decideMobs() {

}

func (g *game) actMobs() {

}
