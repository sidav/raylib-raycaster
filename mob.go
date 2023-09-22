package main

import (
	"math"
	"raylib-raycaster/raycaster"
)

type mobStateCode uint8

const (
	mobStateIdle mobStateCode = iota
	mobStateAttacking
	mobStateMoving
	mobStatePain
	mobStateDying
)

type mob struct {
	x, y, z         float64
	rotationRadians float64

	hitpoints int

	state                 mobStateCode
	ticksSinceStateChange int

	weaponInHands *weapon
	asPlayer      *playerStruct
	intent        *mobIntent

	static *mobStatic
}

func createMob(x, y float64, s *mobStatic) *mob {
	m := &mob{
		x:      x,
		y:      y,
		z:      0.5,
		static: s,
	}
	m.hitpoints = m.static.maxHitpoints
	return m
}

func (t *mob) GetCoords() (float64, float64, float64) {
	return t.x, t.y, 0.5
}

func (t *mob) GetWidthAndHeightFactors() (float64, float64) {
	return 1, 1
}

func (t *mob) GetDirectionVector() (float64, float64) {
	return math.Cos(t.rotationRadians), math.Sin(t.rotationRadians)
}

func (t *mob) GetSprite() *raycaster.SpriteStruct {
	var framesArr [][2]int
	switch t.state {
	case mobStateIdle:
		framesArr = t.static.idleFrames
	case mobStateDying:
		framesArr = t.static.dyingFrames
	}
	maxTick := framesArr[len(framesArr)-1][1]
	for _, sdata := range framesArr {
		if t.ticksSinceStateChange%maxTick < sdata[1] {
			return spritesAtlas[t.static.spriteCode][sdata[0]]
		}
	}
	panic("Frame calculation failure")
}

type mobStatic struct {
	name                    string
	spriteCode              string
	maxHitpoints            int
	idleFrames, dyingFrames [][2]int
}

var sTableMobs = []*mobStatic{
	{
		name:         "Soldier",
		spriteCode:   "soldier",
		maxHitpoints: 30,
		idleFrames:   [][2]int{{0, 30}, {1, 60}},
		dyingFrames:  [][2]int{{2, 10}, {3, 20}, {4, 30}, {5, 100000}},
	},
	{
		name:         "Elite",
		spriteCode:   "slayer",
		maxHitpoints: 50,
		idleFrames:   [][2]int{{0, 30}, {1, 60}},
		dyingFrames:  [][2]int{{2, 10}, {3, 20}, {4, 30}, {5, 100000}},
	},
}
