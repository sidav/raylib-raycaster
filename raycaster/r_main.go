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
	startTime := time.Now()

	r.cam = scene.GetCamera()
	r.scene = scene

	r.aspectFactor = float64(r.RenderWidth) / float64(r.RenderHeight) // * r.cam.distToScreenPlane

	if len(r.rayDistancesBuffer) == 0 {
		r.rayDistancesBuffer = make([]float64, r.RenderWidth)
	}

	if r.ApplyTexturing && r.RenderFloors {
		r.renderFloorAndCeiling()
	} else {
		middleware.SetColor(64, 64, 64)
		middleware.FillRect(0, r.RenderHeight/2, r.RenderWidth, r.RenderHeight/2)
	}

	r.renderWalls()
	r.renderThings()

	elapsedTime := int(time.Since(startTime) / time.Millisecond)
	if elapsedTime != 0 {
		fmt.Printf("Frame rendered in %d ms. (~ %d FPS) \n", elapsedTime, 1000/elapsedTime)
	} else {
		fmt.Printf("Frame rendered in 0 ms. Yay! \n")
	}
}
