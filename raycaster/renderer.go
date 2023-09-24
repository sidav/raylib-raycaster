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
	MinRenderedSpriteHeight      int

	MaxFogFraction, RayLengthForMaximumFog float64
	FogR, FogG, FogB                       uint8 // color for ambient "fog", (0,0,0) is for darkness

	aspectFactor float64
	// currentColumn, deferredColumn r_column

	// private
	rayDistancesBuffer []float64

	// time measure
	columnsTimer, wallsTimer, floorCeilingTimer, thingsTimer Timer

	// surface to draw in
	surface surface
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
		r.surface.create(r.RenderWidth, r.RenderHeight)
	}

	tmr := Timer{}
	tmr.NewMeasure()
	tmr.Measure(func() {
		r.surface.clear()
	})
	debugPrintf("Surface cleared in %d ms.\n", int(r.columnsTimer.GetMeasuredPassedTime()/time.Millisecond))

	r.renderUntexturedFloorAndCeiling()

	r.columnsTimer.NewMeasure()
	r.wallsTimer.NewMeasure()
	r.floorCeilingTimer.NewMeasure()
	// r.renderWalls()
	r.columnsTimer.Measure(func() { r.castRays() })
	debugPrintf("Columns rendered in %d ms (mean %d ms).\n",
		int(r.columnsTimer.GetMeasuredPassedTime()/time.Millisecond),
		int(r.columnsTimer.GetMeanPassedTime()/time.Millisecond),
	)
	debugPrintf(" -> Walls: %d ms (mean %dms), floors/ceilings %dms (mean %d ms).\n",
		int(r.wallsTimer.GetMeasuredPassedTime()/time.Millisecond),
		int(r.wallsTimer.GetMeanPassedTime()/time.Millisecond),
		int(r.floorCeilingTimer.GetMeasuredPassedTime()/time.Millisecond),
		int(r.floorCeilingTimer.GetMeanPassedTime()/time.Millisecond),
	)

	r.thingsTimer.NewMeasure()
	r.thingsTimer.Measure(func() { r.renderThings() })
	debugPrintf("Things rendered in %d ms.\n", int(r.thingsTimer.GetMeasuredPassedTime()/time.Millisecond))

	elapsedTime := int(time.Since(startTimeTotal) / time.Millisecond)
	if elapsedTime != 0 {
		debugPrintf("Frame rendered in %d ms. (~ %d FPS) \n", elapsedTime, 1000/elapsedTime)
	} else {
		debugPrintf("Frame rendered in 0 ms. Yay! \n")
	}
}
