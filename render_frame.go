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
		x -= 128
	case 92:
		x += 24
	}
	drawBackend.DrawRlTextureAt(tex, x, RENDER_H-11*h/30)
}
