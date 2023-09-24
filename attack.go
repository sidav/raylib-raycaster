package main

import "fmt"

func (g *game) doProjectileAttack(proj *projectileStatic, attacker *mob, dx, dy, spreadDegrees float64) {
	dx, dy = rotateVectorRandomlyGauss(dx, dy, spreadDegrees)
	g.scene.things.PushBack(
		g.newProjectile(
			attacker.x, attacker.y, g.scene.Camera.GetVerticalCoordWithBob()-0.1,
			dx, dy,
			proj,
			attacker,
		),
	)
}

func (g *game) doHitscanAttack(hits *hitScanAttack, attacker *mob, dx, dy, spreadDegrees float64) {
	dx, dy = rotateVectorRandomlyGauss(dx, dy, spreadDegrees)
	hitX, hitY, hitMob := g.scene.traceAttackRay(attacker, dx, dy, hits.maxLength)
	if hitMob != nil {
		fmt.Printf("Hit the %+v\n", hitMob)
		hitMob.hitpoints -= hits.damage
	}
	g.scene.things.PushBack(&decoration{
		x:                 hitX,
		y:                 hitY,
		remainingLifetime: 3,
		spriteCode:        hits.hitDecorationSpriteCode,
		width:             0.1,
		height:            0.1,
		blocksMovement:    false,
		blocksProjectiles: false,
	})
}
