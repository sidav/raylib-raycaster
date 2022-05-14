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
		x:    5.5,
		y:    5.5,
		dirY: 1,
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
		case GSTATE_MOBS_DECISION:
			g.decideMobs()
		case GSTATE_MOBS_ACTION:
			g.actMobs()
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
			newX := proj.x + (proj.dirX / factor)
			newY := proj.y + (proj.dirY / factor)
			if !g.scene.areRealCoordsPassable(newX, newY) {
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

func (g *game) decideMobs() {
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		switch node.Value.(type) {
		case *mob:
			currMob := node.Value.(*mob)
			if currMob.intent == nil {
				dirx, diry := 1.0, 0.0
				rotate := rnd.Intn(4)
				for i := 0; i < rotate; i++ {
					dirx, diry = -diry, dirx // rotate 90 degrees
				}
				movx, movy := currMob.x+dirx, currMob.y+diry
				if g.scene.areRealCoordsPassable(movx, movy) {
					currMob.intent = &mobIntent{
						dirx:    dirx,
						diry:    diry,
						moveToX: movx,
						moveToY: movy,
					}
				}
			}
		}
	}
	g.gameState++
}

func (g *game) actMobs() {
	changeState := true
	const factor = 5
	for node := g.scene.things.Front(); node != nil; node = node.Next() {
		switch node.Value.(type) {
		case *mob:
			currMob := node.Value.(*mob)

			if currMob.intent == nil {
				continue
			}

			currMob.x += currMob.intent.dirx / factor
			currMob.y += currMob.intent.diry / factor
			currMob.intent.framesSpent++
			changeState = false

			if currMob.intent.framesSpent == factor {
				currMob.x = currMob.intent.moveToX
				currMob.y = currMob.intent.moveToY
				currMob.intent = nil
			}
		}
	}
	if changeState {
		g.gameState++
	}
}
