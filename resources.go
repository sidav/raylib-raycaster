package main

import "raylib-raycaster/raycaster"

var (
	wallTexturesAtlas map[rune]*raycaster.Texture
	spritesAtlas      map[string]*raycaster.SpriteStruct
)

func loadResources() {
	// init textures (temp.)
	wallTexturesAtlas = make(map[rune]*raycaster.Texture, 0)
	wallTexturesAtlas['#'] = raycaster.InitTextureFromImageFile("textures/wall.png")
	wallTexturesAtlas['+'] = raycaster.InitTextureFromImageFile("textures/door.png")

	// init sprites
	spritesAtlas = make(map[string]*raycaster.SpriteStruct, 0)
	spritesAtlas["proj"] = raycaster.InitSpriteFromImageFile("sprites/projectile.png")
}
