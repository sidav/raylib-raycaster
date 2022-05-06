package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/middleware"
	"raylib-raycaster/raycaster"
)

const (
	WINDOW_W = 800
	WINDOW_H = 600

	IRES_W = 320
	IRES_H = 240
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
		RenderWidth:            IRES_W,
		RenderHeight:           IRES_H,
		ApplyTexturing:         true,
		RenderFloors:           false,
		RenderCeilings:         false,
		MaxRayLength:           25,
		MaxFogFraction:         0,
		RayLengthForMaximumFog: 0,
		FogR:                   0,
		FogG:                   0,
		FogB:                   0,
	}
	s := &Scene{}
	s.init()
	gameIsRunning = true
	middleware.SetInternalResolution(IRES_W, IRES_H)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyUp) {
			s.Camera.MoveForward(0.03)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			s.Camera.Rotate(2 * -3.141592654/180)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			s.Camera.Rotate(2 * 3.141592654/180)
		}


		renderFrame(s)
	}

	rl.CloseWindow()
}
