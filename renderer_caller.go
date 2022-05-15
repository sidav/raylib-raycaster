package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/middleware"
)

func renderFrame(s *Scene) {
	if rl.IsWindowResized() {
		renderer.RenderWidth, renderer.RenderHeight = rl.GetScreenWidth()/PIXEL_SIZE, rl.GetScreenHeight()/PIXEL_SIZE
		middleware.SetInternalResolution(int32(rl.GetScreenWidth()/PIXEL_SIZE), int32(rl.GetScreenHeight()/PIXEL_SIZE))
	}

	rl.BeginTextureMode(middleware.TargetTexture)
	rl.ClearBackground(rl.Black)

	renderer.RenderFrame(s)

	rl.EndTextureMode()
	middleware.Flush()
}
