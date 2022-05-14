package main

import "raylib-raycaster/raycaster"

var (
	wallTexturesAtlas map[string]*raycaster.Texture
	spritesAtlas      map[string]*raycaster.SpriteStruct
)

func loadResources() {
	// init textures (temp.)
	wallTexturesAtlas = make(map[string]*raycaster.Texture, 0)
	wallTexturesAtlas["WALL"] = raycaster.InitTextureFromImageFile("resources/textures/wall.png")
	wallTexturesAtlas["DOOR"] = raycaster.InitTextureFromImageFile("resources/textures/door.png")

	// init sprites
	spritesAtlas = make(map[string]*raycaster.SpriteStruct, 0)
	spritesAtlas["proj"] = raycaster.InitSpriteFromImageFile("resources/sprites/projectile.png")
	spritesAtlas["enemy"] = raycaster.InitSpriteFromImageFile("resources/sprites/enemy.png")
}
