package main

import "container/list"

func (g *game) actThings() {
	var next *list.Element
	for node := g.scene.things.Front(); node != nil; node = next {
		next = node.Next()
		needToRemove := false
		switch node.Value.(type) {
		case (*projectile):
			needToRemove = g.actProjectiles(node.Value.(*projectile))
		case (*decoration):
			needToRemove = g.actDecoration(node.Value.(*decoration))
		case (*pickupable):
			needToRemove = g.actPickupable(node.Value.(*pickupable))
		case (*mob):
			if node.Value.(*mob) == g.player {
				continue
			}
			needToRemove = g.pushMobState(node.Value.(*mob))
			if !needToRemove {
				g.actMob(node.Value.(*mob))
			}
		}
		if needToRemove {
			g.scene.things.Remove(node)
		}
	}
}

func (g *game) actProjectiles(proj *projectile) bool {
	speed := proj.static.speed
	if proj.static.speedIsRandom {
		speed *= float64(rnd.Rand(10000)) / 10000
	}
	newX := proj.x + (proj.dirX * speed)
	newY := proj.y + (proj.dirY * speed)
	hitMob := g.scene.GetMobInRadius(newX, newY, proj.static.sizeFactor/2, proj.creator)
	if hitMob != nil {
		hitMob.hitpoints -= proj.static.damage
		if proj.static.destroysOnMobHit {
			return true
		}
	}
	if !g.scene.areRealCoordsPassable(newX, newY) {
		return true
	} else {
		proj.x, proj.y = newX, newY
		if proj.static.changeFrameEveryTicks > 0 {
			if (g.currentTick-proj.createdAt)%proj.static.changeFrameEveryTicks == 0 {
				proj.frameNum = (proj.frameNum + 1) % proj.static.totalFrames
			}
		}
	}
	return false
}

func (g *game) actDecoration(dec *decoration) bool {
	if dec.remainingLifetime > 0 {
		dec.remainingLifetime--
	}
	if dec.remainingLifetime == 0 {
		return true
	}
	return false
}

func (g *game) actPickupable(pk *pickupable) bool {
	if !areFloatCoordsInRange(pk.x, pk.y, g.player.x, g.player.y, 0.6) {
		return false
	}
	g.player.hitpoints += pk.static.givesHealth
	g.player.hitpoints += pk.static.givesArmor
	return true
}

func (g *game) pushMobState(mob *mob) bool {
	if mob.hitpoints <= 0 && mob.state != mobStateDying {
		mob.changeState(mobStateDying)
		return false
	}
	mob.ticksSinceStateChange++
	if mob.state == mobStateDying && mob.dyingAnimationEnded() {
		g.scene.things.PushBack(&decoration{
			x:                 mob.x,
			y:                 mob.y,
			width:             1,
			height:            1,
			remainingLifetime: -1,
			spriteCode:        mob.static.corpseSpriteCode,
			blocksMovement:    false,
			blocksProjectiles: false,
		})
		return true
	}
	return false
}
