package main

import (
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

	drawBackend.EndFrame()
	drawBackend.Flush()
}
