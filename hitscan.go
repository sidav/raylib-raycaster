package main

import "math"

type hitScanAttack struct {
	hitDecorationSpriteCode string
	damage                  int
	maxLength               float64
}

// returns last UNHIT coords and hit mob if any
func (s *Scene) traceAttackRay(fromx, fromy, dirx, diry, maxLength float64) (float64, float64, *mob) {
	const step = 0.1
	length := 0.0
	for {
		nextX, nextY := fromx+dirx*step, fromy+diry*step
		mob := s.GetMobInRadius(nextX, nextY, 0)
		if mob != nil || !s.areRealCoordsPassable(nextX, nextY) {
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
