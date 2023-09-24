package main

import (
	"math/rand"
	"raylib-raycaster/backend"
	"raylib-raycaster/raycaster"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_W = 1366
	WINDOW_H = 768
	RENDER_W = WINDOW_W / PIXEL_SIZE
	RENDER_H = WINDOW_H / PIXEL_SIZE

	PIXEL_SIZE = 4

	VIEW_ANGLE = 110
)

var (
	gameIsRunning bool
	renderer      *raycaster.Renderer
	drawBackend   *backend.RaylibBackend // backend.RendererBackend
	rnd           *rand.Rand
	tick          int
)

func main() {
	drawBackend = &backend.RaylibBackend{}
	drawBackend.Init(WINDOW_W, WINDOW_H)
	drawBackend.SetInternalResolution(int32(RENDER_W), int32(RENDER_H))
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	renderer = &raycaster.Renderer{
		RenderWidth:             RENDER_W,
		RenderHeight:            RENDER_H,
		ApplyTexturing:          true,
		RenderFloors:            true,
		RenderCeilings:          true,
		MinRenderedSpriteHeight: 2,
		MaxRayLength:            25,
		MaxFogFraction:          1,
		RayLengthForMaximumFog:  20,
		FogR:                    64,
		FogG:                    48,
		FogB:                    32,
	}
	renderer.SetBackend(drawBackend)
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
