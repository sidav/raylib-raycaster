package raycaster

import (
	"fmt"
	"raylib-raycaster/middleware"
	"time"
)

const (
	EW uint8 = iota
	NS
)

type Renderer struct {
	cam   *Camera
	scene Scene

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

func (r *Renderer) RenderFrame(scene Scene) {
	fmt.Printf("=== FRAME START ===\n")
	startTimeTotal := time.Now()

	r.cam = scene.GetCamera()
	r.scene = scene

	r.aspectFactor = float64(r.RenderWidth) / float64(r.RenderHeight) // * r.cam.distToScreenPlane

	if len(r.rayDistancesBuffer) == 0 || len (r.rayDistancesBuffer) != r.RenderWidth {
		r.rayDistancesBuffer = make([]float64, r.RenderWidth)
	}

	if r.ApplyTexturing && r.RenderFloors {
		startTimeFloorsCeilings := time.Now()
		r.renderFloorAndCeiling()
		fmt.Printf("Floors/ceilings rendered in %d ms.\n", int(time.Since(startTimeFloorsCeilings) / time.Millisecond))
	} else {
		middleware.SetColor(64, 64, 64)
		middleware.FillRect(0, r.RenderHeight/2, r.RenderWidth, r.RenderHeight/2)
	}

	startTimeWalls := time.Now()
	r.renderWalls()
	fmt.Printf("Walls rendered in %d ms.\n", int(time.Since(startTimeWalls) / time.Millisecond))
	startTimeThings := time.Now()
	fmt.Printf("Things rendered in %d ms.\n", int(time.Since(startTimeThings) / time.Millisecond))
	r.renderThings()

	elapsedTime := int(time.Since(startTimeTotal) / time.Millisecond)
	if elapsedTime != 0 {
		fmt.Printf("Frame rendered in %d ms. (~ %d FPS) \n", elapsedTime, 1000/elapsedTime)
	} else {
		fmt.Printf("Frame rendered in 0 ms. Yay! \n")
	}
}
