package main

import "raylib-raycaster/raycaster"

type pickupable struct {
	x, y   float64
	static *pickupableStatic
}

func createPickupable(x, y float64, s *pickupableStatic) *pickupable {
	return &pickupable{
		x:      x,
		y:      y,
		static: s,
	}
}

func (p *pickupable) GetCoords() (float64, float64, float64) {
	return p.x, p.y, p.static.sizeFactor * 0.45
}

func (p *pickupable) GetWidthAndHeightFactors() (float64, float64) {
	return p.static.sizeFactor, p.static.sizeFactor
}

func (p *pickupable) GetSprite() *raycaster.SpriteStruct {
	return spritesAtlas[p.static.spriteCode][0]
}

type pickupableStatic struct {
	spriteCode string
	sizeFactor float64

	givesHealth int
	givesArmor  int
}

var sTablePickupables = []*pickupableStatic{
	{
		spriteCode:  "food",
		sizeFactor:  0.3,
		givesHealth: 15,
	},
	{
		spriteCode: "armor",
		sizeFactor: 0.3,
		givesArmor: 100,
	},
}
