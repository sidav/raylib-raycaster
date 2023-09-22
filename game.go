package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
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
		g.actThings()
		g.actTiles()
		g.updatePlayerWeaponState()
		tick++
		g.currentTick++
		fmt.Printf("TOTAL %d THINGS\n", g.scene.things.Len())
		renderFrame(g)
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
	ticksSinceFiring := g.currentTick - wpn.lastTickShot
	if ticksSinceFiring > wpn.static.ticksInFiringState {
		g.scene.Camera.OnScreenVerticalOffset = 0
		wpn.state = wStateIdle
	}
	if wpn.state == wStateFiring && ticksSinceFiring < 3 {
		g.scene.Camera.OnScreenVerticalOffset += rnd.Intn(4) + 1
	}
	if wpn.state == wStateIdle || ticksSinceFiring > 3 {
		g.scene.Camera.OnScreenVerticalOffset = 0
	}
}
