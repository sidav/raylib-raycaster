package raycaster

import (
	"raylib-raycaster/backend"
	"time"
)

const (
	EW uint8 = iota
	NS
)

type Renderer struct {
	backend backend.RendererBackend
	cam     *Camera
	scene   Scene

	RenderWidth                  int
	RenderHeight                 int
	ApplyTexturing               bool
	RenderFloors, RenderCeilings bool
	MaxRayLength                 float64

	MaxFogFraction, RayLengthForMaximumFog float64
	FogR, FogG, FogB                       uint8 // color for ambient "fog", (0,0,0) is for darkness

	aspectFactor float64
	// currentColumn, deferredColumn r_column

	// private
	rayDistancesBuffer []float64
}

func (r *Renderer) SetBackend(b backend.RendererBackend) {
	r.backend = b
}

func (r *Renderer) RenderFrame(scene Scene) {
	debugPrint("=== FRAME START ===")
	startTimeTotal := time.Now()

	r.cam = scene.GetCamera()
	r.scene = scene

	r.aspectFactor = float64(r.RenderWidth) / (1.35 * float64(r.RenderHeight)) // * r.cam.distToScreenPlane

	if len(r.rayDistancesBuffer) == 0 || len(r.rayDistancesBuffer) != r.RenderWidth {
		r.rayDistancesBuffer = make([]float64, r.RenderWidth)
	}

	if r.ApplyTexturing && r.RenderFloors {
		startTimeFloorsCeilings := time.Now()
		r.renderFloorAndCeiling()
		debugPrintf("Floors/ceilings rendered in %d ms.\n", int(time.Since(startTimeFloorsCeilings)/time.Millisecond))
	} else {
		r.backend.SetColor(64, 64, 64)
		r.backend.FillRect(0, r.RenderHeight/2, r.RenderWidth, r.RenderHeight/2)
	}

	startTimeWalls := time.Now()
	r.renderWalls()
	debugPrintf("Walls rendered in %d ms.\n", int(time.Since(startTimeWalls)/time.Millisecond))
	startTimeThings := time.Now()
	debugPrintf("Things rendered in %d ms.\n", int(time.Since(startTimeThings)/time.Millisecond))
	r.renderThings()

	elapsedTime := int(time.Since(startTimeTotal) / time.Millisecond)
	if elapsedTime != 0 {
		debugPrintf("Frame rendered in %d ms. (~ %d FPS) \n", elapsedTime, 1000/elapsedTime)
	} else {
		debugPrintf("Frame rendered in 0 ms. Yay! \n")
	}
}
