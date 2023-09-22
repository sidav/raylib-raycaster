package main

func (g *game) actMob(m *mob) {
	const checkIdleStateEach = 100
	switch m.state {
	case mobStateSleeping:
		if g.currentTick%checkIdleStateEach == 0 {
			if g.doesMobSeePlayer(m) {
				m.state = mobStateIdle
			}
		}
	case mobStateIdle:
		percent := rnd.Intn(100)
		if percent < 25 {
			// m.state = mobStateAttacking
		} else if percent < 50 {
			m.state = mobStateMoving
		}
	case mobStateMoving:
		g.actMobMoving(m)

	}
}

func (g *game) actMobMoving(m *mob) {
	if m.intent == nil {
		dx, dy := rotateVectorRandomly(1, 0, 360)
		m.intent = &mobIntent{
			dirx: dx,
			diry: dy,
		}
	} else {
		if rnd.Intn(25) == 0 {
			m.state = mobStateIdle
			m.intent = nil
			return
		}
	}
	if g.scene.areRealCoordsPassable(m.x+m.intent.dirx, m.y+m.intent.diry) {
		m.x += m.intent.dirx * m.static.speedPerTick
		m.y += m.intent.diry * m.static.speedPerTick
	} else {
		m.state = mobStateIdle
		m.intent = nil
	}
}

func (g *game) doesMobSeePlayer(m *mob) bool {
	return g.scene.unobstructedLineExists(m.x, m.y, g.player.x, g.player.y, 10)
}
