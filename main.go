package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"raylib-raycaster/middleware"
	"raylib-raycaster/raycaster"
	"time"
)

const (
	WINDOW_W = 1000
	WINDOW_H = 800

	IRES_W = WINDOW_W/4 // 320
	IRES_H = WINDOW_H/4 // 240

	VIEW_ANGLE = 135
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
	rnd           *rand.Rand
	tick          int
)

func main() {
	rl.InitWindow(WINDOW_W, WINDOW_H, "RENDERER")
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(30)
	rl.SetExitKey(rl.KeyEscape)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	renderer = &raycaster.Renderer{
		RenderWidth:            IRES_W,
		RenderHeight:           IRES_H,
		ApplyTexturing:         true,
		RenderFloors:           true,
		RenderCeilings:         true,
		MaxRayLength:           2500,
		MaxFogFraction:         0.9,
		RayLengthForMaximumFog: 7,
		FogR:                   64,
		FogG:                   32,
		FogB:                   32,
	}
	middleware.SetInternalResolution(IRES_W, IRES_H)
	loadResources()

	g := &game{}
	g.init()
	g.gameLoop()

	rl.CloseWindow()
}

func trueCoordsToTileCoords(tx, ty float64) (int, int) {
	return int(tx), int(ty)
}

func tileCoordsToPhysicalCoords(tx, ty int) (float64, float64) {
	return float64(tx) + 0.5, float64(ty) + 0.5
}
