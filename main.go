package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/middleware"
	"raylib-raycaster/raycaster"
)

const (
	WINDOW_W = 800
	WINDOW_H = 600

	IRES_W = WINDOW_W / 2
	IRES_H = WINDOW_H / 2
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
	scene         *Scene
)

func main() {
	rl.InitWindow(WINDOW_W, WINDOW_H, "RENDERER")
	rl.SetWindowState(rl.FlagWindowResizable)
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
		RayLengthForMaximumFog: 4,
		FogR:                   0,
		FogG:                   0,
		FogB:                   0,
	}
	scene = &Scene{}
	scene.init()
	gameIsRunning = true
	middleware.SetInternalResolution(IRES_W, IRES_H)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyUp) {
			scene.Camera.MoveForward(0.03)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			scene.Camera.Rotate(3 * -3.141592654 / 180)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			scene.Camera.Rotate(3 * 3.141592654 / 180)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			scene.Camera.MoveForward(-1)
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			scene.things = append(scene.things, &thing{
				x:          scene.Camera.X,
				y:          scene.Camera.Y,
				spriteCode: "proj",
			})
		}

		renderFrame(scene)
	}

	rl.CloseWindow()
}
