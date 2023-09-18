package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderFrame(g *game) {
	s := g.scene
	if rl.IsWindowResized() {
		renderer.RenderWidth, renderer.RenderHeight = rl.GetScreenWidth()/PIXEL_SIZE, rl.GetScreenHeight()/PIXEL_SIZE
		drawBackend.SetInternalResolution(int32(rl.GetScreenWidth()/PIXEL_SIZE), int32(rl.GetScreenHeight()/PIXEL_SIZE))
	}

	drawBackend.BeginFrame()
	rl.ClearBackground(rl.Black)

	renderer.RenderFrame(s)

	drawWeaponInHands(g)

	drawBackend.EndFrame()
	drawBackend.Flush()
}

func drawWeaponInHands(g *game) {
	weap := g.player.weaponInHands
	if weap == nil {
		return
	}
	tex := uiAtlas[weap.static.spriteCode][weap.getSpriteFrame(g.currentTick)]
	w, h := tex.Width, tex.Height
	fmt.Printf("%d, %d\n", w, h)
	var x int32 = RENDER_W / 2
	// I dunno where those magic numbers come from, something wrong with my RayLib texture-mode code
	switch w {
	case 64:
		x += 38
	case 92:
		x += 24
	case 256:
		x -= 64
	default:
		panic("Rl texture coords unfinished")
	}
	var y int32 = RENDER_H
	switch h {
	case 92:
		y -= 32
	case 128:
		y -= 68
	default:
		panic("Rl texture coords unfinished")
	}
	drawBackend.DrawRlTextureAt(tex, x, y)
}
