package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylib-raycaster/middleware"
	"raylib-raycaster/raycaster"
)

const (
	WINDOW_W = 1000
	WINDOW_H = 800

	IRES_W = 250
	IRES_H = 200
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
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
		MaxFogFraction:         0.9,
		RayLengthForMaximumFog: 5,
		FogR:                   64,
		FogG:                   64,
		FogB:                   32,
	}
	middleware.SetInternalResolution(IRES_W, IRES_H)
	loadResources()

	g := &game{}
	g.init()
	g.gameLoop()

	rl.CloseWindow()
}
