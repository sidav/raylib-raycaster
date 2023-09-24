package main

import "math"

type hitScanAttack struct {
	hitDecorationSpriteCode string
	damage                  int
	maxLength               float64
}

// returns last UNHIT coords and hit mob if any
func (s *Scene) traceAttackRay(attacker *mob, dirx, diry, maxLength float64) (float64, float64, *mob) {
	const step = 0.1
	fromx, fromy, _ := attacker.GetCoords()
	length := 0.0
	var mob *mob
	for stepNum := 0; ; stepNum++ {
		nextX, nextY := fromx+dirx*step, fromy+diry*step
		if stepNum%2 == 0 { // slight optimization
			mob = s.GetMobInRadius(nextX, nextY, 0, attacker)
			if mob != nil {
				return fromx, fromy, mob
			}
		}
		if !s.areRealCoordsPassable(nextX, nextY) {
			return fromx, fromy, mob
		}
		fromx = nextX
		fromy = nextY
		length += step
		if length > maxLength {
			return fromx, fromy, mob
		}
	}
}

// returns last UNHIT coords
func (s *Scene) unobstructedLineExists(fromx, fromy, tox, toy, maxLength float64) bool {
	dirx, diry := tox-fromx, toy-fromy
	length := math.Sqrt(dirx*dirx + diry*diry)
	dirx /= length
	diry /= length
	const step = 0.1
	length = 0.0
	for {
		nextX, nextY := fromx+dirx*step, fromy+diry*step
		// close enough to destination
		if (nextX-tox)*(nextX-tox)+(nextY-toy)*(nextY-toy) <= 0.1 {
			return true
		}
		if !s.areRealCoordsPassable(nextX, nextY) {
			return false
		}
		fromx = nextX
		fromy = nextY
		length += step
		if length > maxLength {
			return false
		}
	}
}
