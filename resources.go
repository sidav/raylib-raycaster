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

	// non-renderer related: UI, weapons in hands etc
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
	loadSprite("projPlasma", "resources/sprites/projectile0.png")
	loadSprite("projPlasma", "resources/sprites/projectile1.png")
	loadSprite("projAcid", "resources/sprites/lab/sprites/acidspit.png")
	loadSprite("projFireball", "resources/sprites/lab/sprites/fireball0.png")
	loadSprite("enemy", "resources/sprites/characters/CommandoWalk1.png")
	loadSprite("enemy", "resources/sprites/characters/CommandoWalk2.png")

	loadUIImage("pWeaponPistol", "resources/sprites/lab/guns/gun2.png")
	loadUIImage("pWeaponPistol", "resources/sprites/lab/guns/gun2b.png")
	loadUIImage("pWeaponPistol", "resources/sprites/lab/guns/gun2c.png")

	loadUIImage("pWeaponGun", "resources/sprites/lab/guns/gun1a.png")
	loadUIImage("pWeaponGun", "resources/sprites/lab/guns/gun1b.png")
	loadUIImage("pWeaponGun", "resources/sprites/lab/guns/gun1c.png")

	loadUIImage("pWeaponSmg", "resources/sprites/lab/guns/gun5a.png")
	loadUIImage("pWeaponSmg", "resources/sprites/lab/guns/gun5b.png")
	loadUIImage("pWeaponSmg", "resources/sprites/lab/guns/gun5c.png")

	loadUIImage("pWeaponGun2", "resources/sprites/lab/guns/gun4.png")
	loadUIImage("pWeaponGun2", "resources/sprites/lab/guns/gun4b.png")
	loadUIImage("pWeaponGun2", "resources/sprites/lab/guns/gun4c.png")
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

func loadUIImage(code, filename string) {
	if len(uiAtlas) == 0 {
		uiAtlas = make(map[string][]rl.Texture2D)
	}
	uiAtlas[code] = append(uiAtlas[code], drawBackend.LoadImageAsRlTexture(filename))
}
