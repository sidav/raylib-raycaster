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

func (t *mob) changeState(s mobStateCode) {
	t.state = s
	t.intent = nil
	t.ticksSinceStateChange = 0
}

func (t *mob) dyingAnimationEnded() bool {
	ticksNeeded := mobTicksPerDyingFrame * len(t.static.dyingFrames)
	return t.ticksSinceStateChange >= ticksNeeded
}

func (t *mob) attackingAnimationEnded() bool {
	ticksNeeded := mobTicksPerAttackingFrame * len(t.static.attackingFrames)
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
	case mobStateAttacking:
		return spritesAtlas[t.static.spriteCode][t.static.attackingFrames[t.ticksSinceStateChange/mobTicksPerAttackingFrame%len(t.static.attackingFrames)]]
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
	name                                       string
	spriteCode                                 string
	corpseSpriteCode                           string
	maxHitpoints                               int
	speedPerTick                               float64
	idleFrames                                 [][2]int
	dyingFrames, movingFrames, attackingFrames []int

	spreadDegrees   float64
	firesProjectile *projectileStatic
}

const mobTicksPerDyingFrame = 5
const mobTicksPerMovingFrame = 5
const mobTicksPerAttackingFrame = 7

var sTableMobs = []*mobStatic{
	{
		name:             "Soldier",
		spriteCode:       "soldier",
		corpseSpriteCode: "soldiercorpse",
		maxHitpoints:     30,
		speedPerTick:     0.065,
		idleFrames:       [][2]int{{0, 30}, {1, 60}},
		dyingFrames:      []int{2, 3, 4},
		movingFrames:     []int{5, 6, 7, 8},
		attackingFrames:  []int{9, 10},

		spreadDegrees: 5,
		firesProjectile: &projectileStatic{
			spriteCode:  "projFireball",
			totalFrames: 1,
			speed:       0.7,
			damage:      15,
			sizeFactor:  0.1,
		},
	},
	{
		name:             "Elite",
		spriteCode:       "slayer",
		corpseSpriteCode: "slayercorpse",
		maxHitpoints:     50,
		speedPerTick:     0.035,
		idleFrames:       [][2]int{{0, 30}, {1, 60}},
		dyingFrames:      []int{2, 3, 4},
		movingFrames:     []int{5, 6, 7, 8},
		attackingFrames:  []int{9, 10},

		spreadDegrees: 5,
		firesProjectile: &projectileStatic{
			spriteCode:  "projFireball",
			totalFrames: 1,
			speed:       0.7,
			damage:      15,
			sizeFactor:  0.1,
		},
	},
}
