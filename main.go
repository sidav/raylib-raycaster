package main

import (
	"raylib-raycaster/backend"
	"raylib-raycaster/lib/random"
	"raylib-raycaster/lib/random/pcgrandom"
	"raylib-raycaster/raycaster"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PIXEL_SIZE = 4
	VIEW_ANGLE = 110
)

var (
	WINDOW_W      = 1366
	WINDOW_H      = 768
	RENDER_W      = WINDOW_W / PIXEL_SIZE
	RENDER_H      = WINDOW_H / PIXEL_SIZE
	gameIsRunning bool
	renderer      *raycaster.Renderer
	drawBackend   *backend.RaylibBackend // backend.RendererBackend
	rnd           random.PRNG
	tick          int
)

func main() {
	drawBackend = &backend.RaylibBackend{}
	drawBackend.Init(int32(WINDOW_W), int32(WINDOW_H))
	drawBackend.SetInternalResolution(int32(RENDER_W), int32(RENDER_H))
	rnd = pcgrandom.NewPCG64()
	rnd.SetSeed(int(time.Now().UnixNano()))

	renderer = &raycaster.Renderer{
		RenderThreads:           8,
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
