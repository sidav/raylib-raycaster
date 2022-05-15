package main

import "raylib-raycaster/raycaster"

var (
	wallTexturesAtlas    map[string][]*raycaster.Texture
	floorTexturesAtlas   map[string][]*raycaster.Texture
	ceilingTexturesAtlas map[string][]*raycaster.Texture
	spritesAtlas         map[string][]*raycaster.SpriteStruct
)

func loadResources() {
	// init textures (temp.)
	loadTexture(&wallTexturesAtlas, "WALL", "resources/textures/tile091.png")
	loadTexture(&wallTexturesAtlas, "DOOR", "resources/textures/door.png")

	loadTexture(&floorTexturesAtlas, "DEFAULT", "resources/textures/tile126.png")
	loadTexture(&floorTexturesAtlas, "DOOR", "resources/textures/tile110.png")

	loadTexture(&ceilingTexturesAtlas, "DEFAULT", "resources/textures/tile110.png")
	// init sprites
	loadSprite("proj", "resources/sprites/projectile0.png")
	loadSprite("proj", "resources/sprites/projectile1.png")
	loadSprite("enemy", "resources/sprites/cobra1.png")
	loadSprite("enemy", "resources/sprites/cobra2.png")
}

func loadTexture(atlas *map[string][]*raycaster.Texture, code string, filename string) {
	if *atlas == nil {
		*atlas = make(map[string][]*raycaster.Texture, 0)
	}
	(*atlas)[code] = append((*atlas)[code], raycaster.InitTextureFromImageFile(filename))
}

func loadSprite(code string, filename string) {
	if spritesAtlas == nil {
		spritesAtlas = make(map[string][]*raycaster.SpriteStruct, 0)
	}
	spritesAtlas[code] = append(spritesAtlas[code], raycaster.InitSpriteFromImageFile(filename))
}
