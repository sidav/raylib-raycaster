package main

import (
	"math"
	"raylib-raycaster/raycaster"
)

type mobStateCode uint8

const (
	mobStateSleeping mobStateCode = iota
	mobStateIdle
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

func (t *mob) dyingAnimationEnded() bool {
	ticksNeeded := mobTicksPerDyingFrame * len(t.static.dyingFrames)
	return t.ticksSinceStateChange >= ticksNeeded
}

func (t *mob) GetSprite() *raycaster.SpriteStruct {
	var framesArr [][2]int
	switch t.state {
	case mobStateSleeping, mobStateIdle:
		framesArr = t.static.idleFrames
	case mobStateDying:
		return spritesAtlas[t.static.spriteCode][t.static.dyingFrames[t.ticksSinceStateChange/mobTicksPerDyingFrame]]
	case mobStateMoving:
		return spritesAtlas[t.static.spriteCode][t.static.movingFrames[t.ticksSinceStateChange/mobTicksPerMovingFrame%len(t.static.movingFrames)]]
	default:
		panic("State not implemented")
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
	name                      string
	spriteCode                string
	corpseSpriteCode          string
	maxHitpoints              int
	speedPerTick              float64
	idleFrames                [][2]int
	dyingFrames, movingFrames []int
}

const mobTicksPerDyingFrame = 5
const mobTicksPerMovingFrame = 5

var sTableMobs = []*mobStatic{
	{
		name:             "Soldier",
		spriteCode:       "soldier",
		corpseSpriteCode: "soldiercorpse",
		maxHitpoints:     30,
		speedPerTick:     0.07,
		idleFrames:       [][2]int{{0, 30}, {1, 60}},
		dyingFrames:      []int{2, 3, 4},
		movingFrames:     []int{5, 6, 7, 8},
	},
	{
		name:             "Elite",
		spriteCode:       "slayer",
		corpseSpriteCode: "slayercorpse",
		maxHitpoints:     50,
		speedPerTick:     0.04,
		idleFrames:       [][2]int{{0, 30}, {1, 60}},
		dyingFrames:      []int{2, 3, 4},
		movingFrames:     []int{5, 6, 7, 8},
	},
}
