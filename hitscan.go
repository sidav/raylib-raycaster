package main

type hitScanAttack struct {
	hitDecorationSpriteCode string
	damage                  int
	maxLength               float64
}

// returns last UNHIT coords and hit mob if any
func (s *Scene) traceRay(fromx, fromy, dirx, diry, maxLength float64) (float64, float64, *mob) {
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
