package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderFrame(g *game) {
	s := g.scene
	if rl.IsWindowResized() {
		RENDER_W, RENDER_H = rl.GetScreenWidth()/PIXEL_SIZE, rl.GetScreenHeight()/PIXEL_SIZE
		renderer.RenderWidth, renderer.RenderHeight = RENDER_W, RENDER_H
		drawBackend.SetInternalResolution(int32(RENDER_W), int32(RENDER_H))
	}

	drawBackend.BeginFrame()
	rl.ClearBackground(rl.Black)

	renderer.RenderFrame(s)
	renderer.PresentDrawnSurfaceToBackend()

	drawHUD(g)
	drawWeaponInHands(g)

	drawBackend.EndFrame()

	drawBackend.Flush()
}

func drawHUD(g *game) {
	rl.DrawText(fmt.Sprintf("HEALTH %d/%d", g.player.hitpoints, 100), 0, 0, 16, rl.Gray)
}

func drawWeaponInHands(g *game) {
	weap := g.player.weaponInHands
	if weap == nil {
		return
	}
	tex := uiAtlas[weap.static.spriteCode][weap.getSpriteFrame(g.currentTick)]
	w, h := tex.Width, tex.Height
	drawBackend.DrawRlTextureAt(tex, int32(RENDER_W/2)-w/2, int32(RENDER_H)-h)
}
