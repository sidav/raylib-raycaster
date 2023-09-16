package main

import (
	"raylib-raycaster/raycaster"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	wallTexturesAtlas    map[string][]*raycaster.Texture
	floorTexturesAtlas   map[string][]*raycaster.Texture
	ceilingTexturesAtlas map[string][]*raycaster.Texture
	spritesAtlas         map[string][]*raycaster.SpriteStruct

	uiAtlas map[string][]rl.Texture2D
)

func loadResources() {
	// init textures (temp.)
	loadTexture(&wallTexturesAtlas, "WALL", "resources/textures/scifi/suprt2.png")
	loadTexture(&wallTexturesAtlas, "DOOR", "resources/textures/door.png")

	loadTexture(&floorTexturesAtlas, "DEFAULT", "resources/textures/tile113.png")
	loadTexture(&floorTexturesAtlas, "DOOR", "resources/textures/tile110.png")

	loadTexture(&ceilingTexturesAtlas, "DEFAULT", "resources/textures/tile110.png")
	// init sprites
	loadSprite("proj", "resources/sprites/projectile0.png")
	loadSprite("proj", "resources/sprites/projectile1.png")
	loadSprite("enemy", "resources/sprites/characters/CommandoWalk1.png")
	loadSprite("enemy", "resources/sprites/characters/CommandoWalk2.png")

	uiAtlas = make(map[string][]rl.Texture2D)
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
