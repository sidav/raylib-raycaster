package main

func (g *game) actMob(m *mob) {
	const checkIdleStateEach = 100
	switch m.state {
	case mobStateSleeping:
		if g.currentTick%checkIdleStateEach == 0 {
			if g.doesMobSeePlayer(m) {
				m.changeState(mobStateIdle)
			}
		}
	case mobStateIdle:
		percent := rnd.Intn(100)
		if percent < m.static.aiData.chanceToAttack {
			m.changeState(mobStateAttacking)
		} else if percent < m.static.aiData.chanceToAttack+m.static.aiData.chanceToMove {
			m.changeState(mobStateMoving)
		}
	case mobStateMoving:
		g.actMobMoving(m)
	case mobStateAttacking:
		g.actMobAttacking(m)

	}
}

func (g *game) actMobAttacking(m *mob) {
	if m.ticksSinceStateChange == 1 {
		dx, dy := makeUnitVector(m.x, m.y, g.player.x, g.player.y)
		g.doProjectileAttack(m.static.firesProjectile, m, dx, dy, m.static.spreadDegrees)
	}
	if m.attackingAnimationEnded() {
		if rnd.Intn(100) < m.static.aiData.aggressiveness {
			m.changeState(mobStateAttacking)
		} else {
			m.changeState(mobStateIdle)
		}
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
		if rnd.Intn(m.static.aiData.averageMoveSteps) == 0 {
			m.changeState(mobStateIdle)
			return
		}
	}
	if g.scene.areRealCoordsPassable(m.x+m.intent.dirx, m.y+m.intent.diry) {
		m.x += m.intent.dirx * m.static.speedPerTick
		m.y += m.intent.diry * m.static.speedPerTick
	} else {
		m.changeState(mobStateIdle)
	}
}

func (g *game) doesMobSeePlayer(m *mob) bool {
	return g.scene.unobstructedLineExists(m.x, m.y, g.player.x, g.player.y, 10)
}
