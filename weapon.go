package main

type weaponState uint8

const (
	wStateIdle weaponState = iota
	wStateFiring
)

type weapon struct {
	state        weaponState
	static       *weaponStatic
	lastTickShot int
}

func (w *weapon) getSpriteFrame(currentTick int) int {
	switch w.state {
	case wStateIdle:
		return 0
	case wStateFiring:
		for _, arr := range w.static.firingFrames {
			if currentTick-w.lastTickShot <= arr[1] {
				return arr[0]
			}
		}
		return 0
	}
	panic("Unimplemented")
}

func (w *weapon) canShoot() bool {
	return w.state == wStateIdle
}

type weaponStatic struct {
	name       string
	spriteCode string // unneeded for mobs' weapons

	ticksInFiringState int
	firingFrames       [][2]int // array of: {frame number, max ticks since firing for this sprite to be the current one}

	spreadDegrees   float64 // spread cone angle IN DEGREES
	shotsPerShot    int
	firesProjectile *projectileStatic
	firesHitscan    *hitScanAttack
}

var sTableWeapons = []*weaponStatic{
	{
		name:               "Pistol",
		spriteCode:         "pWeaponPistol",
		ticksInFiringState: 20,
		firingFrames:       [][2]int{{1, 5}, {2, 10}},
		spreadDegrees:      3,
		shotsPerShot:       1,
		firesProjectile: &projectileStatic{
			spriteCode:  "projPlasma",
			totalFrames: 1,
			speed:       0.85,
			damage:      30,
			sizeFactor:  0.1,
		},
	},
	{
		name:               "SMG",
		spriteCode:         "pWeaponSmg",
		ticksInFiringState: 2,
		spreadDegrees:      6,
		shotsPerShot:       1,
		firingFrames:       [][2]int{{1, 1}, {2, 10}},
		firesHitscan: &hitScanAttack{
			hitDecorationSpriteCode: "projFireball",
			damage:                  4,
			maxLength:               10,
		},
	},
	{
		name:               "Gun",
		spriteCode:         "pWeaponGun",
		ticksInFiringState: 10,
		spreadDegrees:      10,
		shotsPerShot:       7,
		firingFrames:       [][2]int{{1, 1}, {2, 10}},
		firesHitscan: &hitScanAttack{
			hitDecorationSpriteCode: "projFireball",
			damage:                  3,
			maxLength:               10,
		},
	},
	{
		name:               "Gun2",
		spriteCode:         "pWeaponGun2",
		ticksInFiringState: 60,
		firingFrames:       [][2]int{{1, 10}, {2, 35}},
		spreadDegrees:      1,
		shotsPerShot:       1,
		firesProjectile: &projectileStatic{
			spriteCode:  "projAcid",
			totalFrames: 1,
			speed:       0.2,
			damage:      100,
			sizeFactor:  0.75,
		},
	},
}
