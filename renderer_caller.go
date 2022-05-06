package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/middleware"
)

func renderFrame(s *Scene) {
	rl.BeginTextureMode(middleware.TargetTexture)
	rl.ClearBackground(rl.Black)

	renderer.RenderFrame(s)

	rl.EndTextureMode()
	middleware.Flush()
}
