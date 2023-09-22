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
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		switch node.Value.(type) {
		case *projectile:
			proj := node.Value.(*projectile)
			newX := proj.x + (proj.dirX * proj.static.speed)
			newY := proj.y + (proj.dirY * proj.static.speed)
			hitMob := g.scene.GetMobAtRealCoords(newX, newY)
			if hitMob != nil {
				g.scene.things.Remove(node)
				hitMob.hitpoints -= proj.static.damage
				continue
			}
			if !g.scene.areRealCoordsPassable(newX, newY) {
				g.scene.things.Remove(node)
			} else {
				proj.x, proj.y = newX, newY
				if proj.static.changeFrameEveryTicks > 0 {
					if (g.currentTick-proj.createdAt)%proj.static.changeFrameEveryTicks == 0 {
						proj.frameNum = (proj.frameNum + 1) % proj.static.totalFrames
					}
				}
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
	ticksSinceFiring := g.currentTick - wpn.lastTickShot
	if ticksSinceFiring > wpn.static.ticksInFiringState {
		g.scene.Camera.OnScreenVerticalOffset = 0
		wpn.state = wStateIdle
	}
	if wpn.state == wStateFiring && ticksSinceFiring < 3 {
		g.scene.Camera.OnScreenVerticalOffset += 3
	}
	if wpn.state == wStateIdle || ticksSinceFiring > 3 {
		g.scene.Camera.OnScreenVerticalOffset = 0
	}
}

func (g *game) decideMobs() {
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		if mob, ok := node.Value.(*mob); ok {
			if mob.hitpoints <= 0 && mob.state != mobStateDying {
				mob.state = mobStateDying
				mob.ticksSinceStateChange = 0
			} else {
				mob.ticksSinceStateChange++
				if mob.state == mobStateDying && mob.dyingAnimationEnded() {
					g.scene.things.PushBack(&decoration{
						x:                 mob.x,
						y:                 mob.y,
						spriteCode:        mob.static.corpseSpriteCode,
						blocksMovement:    false,
						blocksProjectiles: false,
					})
					g.scene.things.Remove(node)
				}
			}
		}
	}
}

func (g *game) actMobs() {

}
