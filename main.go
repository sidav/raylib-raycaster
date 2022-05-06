package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/raycaster"
)

const (
	WINDOW_W = 800
	WINDOW_H = 600
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
)

func main() {
	rl.InitWindow(WINDOW_W, WINDOW_H, "RENDERER")
	rl.SetTargetFPS(30)
	rl.SetExitKey(rl.KeyEscape)

	renderer = &raycaster.Renderer{
		RenderWidth:            WINDOW_W,
		RenderHeight:           WINDOW_H,
		ApplyTexturing:         true,
		RenderFloors:           false,
		RenderCeilings:         false,
		MaxRayLength:           1000000000,
		MaxFogFraction:         0,
		RayLengthForMaximumFog: 0,
		FogR:                   0,
		FogG:                   0,
		FogB:                   0,
	}
	s := &Scene{}
	s.init()
	gameIsRunning = true

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		s.Camera.Rotate(3.141592654/180)
		s.Camera.MoveForward(0.03)
		renderer.RenderFrame(s)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
