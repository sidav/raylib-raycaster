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
	g.player = createMob(
		px, py, &mobStatic{
			name:         "Player",
			maxHitpoints: 100,
			speedPerTick: 0.065,
		},
	)
	g.player.weaponInHands = &weapon{
		static: sTableWeapons[0],
	}
	g.scene.things.PushFront(g.player)
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
			t := &(g.scene.gameMap[x][y])
			if t.state == tileStateIdle {
				continue
			}
			speed := 0.1
			switch t.state {
			case tileStateOpening:
				t.tileSlideAmount += speed
				if t.isOpened() {
					t.state = tileStateWaitsToClose
					t.tileStateCooldown = 75
				}
			case tileStateWaitsToClose:
				t.tileStateCooldown--
				if t.tileStateCooldown == 0 {
					t.state = tileStateClosing
				}
			case tileStateClosing:
				t.tileSlideAmount -= speed
				px, py := trueCoordsToTileCoords(g.player.x, g.player.y)
				if px == x && py == y || g.scene.GetMobAtTileCoords(x, y) != nil {
					t.state = tileStateOpening
					continue
				}
				if t.isClosed() {
					t.tileSlideAmount = 0
					t.state = tileStateIdle
				}
			}
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
