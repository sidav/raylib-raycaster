package main

import (
	"math/rand"
	"raylib-raycaster/backend"
	"raylib-raycaster/raycaster"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_W = 1360
	WINDOW_H = 768
	RENDER_W = WINDOW_W / PIXEL_SIZE
	RENDER_H = WINDOW_H / PIXEL_SIZE

	PIXEL_SIZE = 6

	VIEW_ANGLE = 120
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
	drawBackend   *backend.RaylibBackend
	rnd           *rand.Rand
	tick          int
)

func main() {
	drawBackend = &backend.RaylibBackend{}
	rl.InitWindow(WINDOW_W, WINDOW_H, "RENDERER")
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(30)
	rl.SetExitKey(rl.KeyEscape)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	renderer = &raycaster.Renderer{
		RenderWidth:            RENDER_W,
		RenderHeight:           RENDER_H,
		ApplyTexturing:         true,
		RenderFloors:           true,
		RenderCeilings:         true,
		MaxRayLength:           20,
		MaxFogFraction:         0.9,
		RayLengthForMaximumFog: 7,
		FogR:                   64,
		FogG:                   32,
		FogB:                   32,
	}
	renderer.SetBackend(drawBackend)
	drawBackend.SetInternalResolution(int32(rl.GetScreenWidth()/PIXEL_SIZE), int32(rl.GetScreenHeight()/PIXEL_SIZE))
	loadResources()

	g := &game{}
	g.init()
	g.gameLoop()

	rl.CloseWindow()
}

func trueCoordsToTileCoords(tx, ty float64) (int, int) {
	return int(tx), int(ty)
}

func tileCoordsToTrueCoords(tx, ty int) (float64, float64) {
	return float64(tx) + 0.5, float64(ty) + 0.5
}
