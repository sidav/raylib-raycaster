package main

import "raylib-raycaster/raycaster"

var (
	wallTexturesAtlas map[string]*raycaster.Texture
	spritesAtlas      map[string][]*raycaster.SpriteStruct
)

func loadResources() {
	// init textures (temp.)
	wallTexturesAtlas = make(map[string]*raycaster.Texture, 0)
	wallTexturesAtlas["WALL"] = raycaster.InitTextureFromImageFile("resources/textures/wall.png")
	wallTexturesAtlas["DOOR"] = raycaster.InitTextureFromImageFile("resources/textures/door.png")

	// init sprites
	loadSprite("proj", "resources/sprites/projectile0.png")
	loadSprite("proj", "resources/sprites/projectile1.png")
	loadSprite("enemy", "resources/sprites/cobra1.png")
	loadSprite("enemy", "resources/sprites/cobra2.png")
}

func loadSprite(code string, filename string) {
	if spritesAtlas == nil {
		spritesAtlas = make(map[string][]*raycaster.SpriteStruct, 0)
	}
	spritesAtlas[code] = append(spritesAtlas[code], raycaster.InitSpriteFromImageFile(filename))
}
