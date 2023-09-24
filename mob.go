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
	ticksNeeded := t.static.ticksPerDyingFrame * len(t.static.dyingFrames)
	return t.ticksSinceStateChange >= ticksNeeded
}

func (t *mob) attackingAnimationEnded() bool {
	ticksNeeded := t.static.ticksPerAttackingFrame * len(t.static.attackingFrames)
	return t.ticksSinceStateChange >= ticksNeeded
}

func (t *mob) GetSprite() *raycaster.SpriteStruct {
	if t.static.spriteCode == "" {
		// may be useful for unrenderable things (for now it is Player)
		return nil
	}
	var framesArr []int
	var ticksPerFrame int
	switch t.state {
	case mobStateSleeping, mobStateIdle:
		framesArr = t.static.idleFrames
		ticksPerFrame = t.static.ticksPerIdleFrame
	case mobStateDying:
		framesArr = t.static.dyingFrames
		ticksPerFrame = t.static.ticksPerDyingFrame
	case mobStateMoving:
		framesArr = t.static.movingFrames
		ticksPerFrame = t.static.ticksPerMovingFrame
	case mobStateAttacking:
		framesArr = t.static.attackingFrames
		ticksPerFrame = t.static.ticksPerAttackingFrame
	default:
		panic("State not implemented")
	}
	return spritesAtlas[t.static.spriteCode][framesArr[t.ticksSinceStateChange/ticksPerFrame%len(framesArr)]]
}

type mobStatic struct {
	name                                                                               string
	spriteCode                                                                         string
	corpseSpriteCode                                                                   string
	maxHitpoints                                                                       int
	speedPerTick                                                                       float64
	idleFrames, dyingFrames, movingFrames, attackingFrames                             []int
	ticksPerIdleFrame, ticksPerDyingFrame, ticksPerMovingFrame, ticksPerAttackingFrame int

	aiData *mobAiStatic

	spreadDegrees   float64
	firesProjectile *projectileStatic
}

type mobAiStatic struct {
	aggressiveness               int // chance (%) to repeat attack right after previous one
	chanceToAttack, chanceToMove int
	averageMoveSteps             int
}

var sTableMobs = []*mobStatic{
	{
		name:                   "Soldier",
		spriteCode:             "soldier",
		corpseSpriteCode:       "soldiercorpse",
		maxHitpoints:           30,
		speedPerTick:           0.045,
		idleFrames:             []int{0, 1},
		ticksPerIdleFrame:      25,
		dyingFrames:            []int{2, 3, 4},
		ticksPerDyingFrame:     5,
		movingFrames:           []int{5, 6, 7, 8},
		ticksPerMovingFrame:    5,
		attackingFrames:        []int{9, 10},
		ticksPerAttackingFrame: 7,

		aiData: &mobAiStatic{
			aggressiveness:   10,
			chanceToAttack:   40,
			chanceToMove:     60,
			averageMoveSteps: 15,
		},

		spreadDegrees: 5,
		firesProjectile: &projectileStatic{
			spriteCode:  "projFireball",
			totalFrames: 1,
			speed:       0.3,
			damage:      6,
			sizeFactor:  0.1,
		},
	},
	{
		name:                   "Elite",
		spriteCode:             "slayer",
		corpseSpriteCode:       "slayercorpse",
		maxHitpoints:           50,
		speedPerTick:           0.03,
		idleFrames:             []int{0, 1},
		ticksPerIdleFrame:      25,
		dyingFrames:            []int{2, 3, 4},
		ticksPerDyingFrame:     5,
		movingFrames:           []int{5, 6, 7, 8},
		ticksPerMovingFrame:    5,
		attackingFrames:        []int{9, 10},
		ticksPerAttackingFrame: 4,

		aiData: &mobAiStatic{
			aggressiveness:   65,
			chanceToAttack:   5,
			chanceToMove:     10,
			averageMoveSteps: 25,
		},

		spreadDegrees: 5,
		firesProjectile: &projectileStatic{
			spriteCode:  "projAcid",
			totalFrames: 1,
			speed:       0.3,
			damage:      14,
			sizeFactor:  0.1,
		},
	},
}
