package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderFrame(s *Scene) {
	if rl.IsWindowResized() {
		renderer.RenderWidth, renderer.RenderHeight = rl.GetScreenWidth()/PIXEL_SIZE, rl.GetScreenHeight()/PIXEL_SIZE
		drawBackend.SetInternalResolution(int32(rl.GetScreenWidth()/PIXEL_SIZE), int32(rl.GetScreenHeight()/PIXEL_SIZE))
	}

	drawBackend.BeginFrame()
	rl.ClearBackground(rl.Black)

	renderer.RenderFrame(s)

	drawWeaponInHands()

	drawBackend.EndFrame()
	drawWeaponInHands()
	drawBackend.Flush()
}

func drawWeaponInHands() {
	tex := uiAtlas["pWeaponPistol"][0]
	tex = uiAtlas["pWeaponGun"][0]
	w, h := tex.Width, tex.Height
	fmt.Printf("%d, %d\n", w, h)
	// I dunno where those magic numbers come from, something wrong with RayLib texture
	drawBackend.DrawRlTextureAt(tex, RENDER_W-2*w, RENDER_H-11*h/30)
}
