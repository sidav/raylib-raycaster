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
	name               string
	spriteCode         string // unneeded for mobs' weapons
	ticksInFiringState int

	firingFrames [][2]int // array of: {frame number, max ticks since firing for this sprite to be the current one}
}

var sTableWeapons = []*weaponStatic{
	{
		name:               "Pistol",
		spriteCode:         "pWeaponPistol",
		ticksInFiringState: 20,
		firingFrames:       [][2]int{{1, 5}, {2, 10}},
	},
	{
		name:               "Gun",
		spriteCode:         "pWeaponGun",
		ticksInFiringState: 10,
		firingFrames:       [][2]int{{1, 3}, {2, 10}},
	},
}
