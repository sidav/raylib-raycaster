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
	middleware.SetInternalResolution(IRES_W, IRES_H)
	loadResources()

	g := &game{}
	g.init()
	g.gameLoop()

	rl.CloseWindow()
}
