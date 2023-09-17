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

	// time measure
	columnsTimer, wallsTimer, floorCeilingTimer, thingsTimer timer
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

	r.renderUntexturedFloorAndCeiling()

	r.columnsTimer.newMeasure()
	r.wallsTimer.newMeasure()
	r.floorCeilingTimer.newMeasure()
	// r.renderWalls()
	r.columnsTimer.measure(func() { r.castRays() })
	debugPrintf("Columns rendered in %d ms (mean %d ms).\n",
		int(r.columnsTimer.getMeasuredPassedTime()/time.Millisecond),
		int(r.columnsTimer.getMeanPassedTime()/time.Millisecond),
	)
	debugPrintf(" -> Walls: %d ms (mean %dms), floors/ceilings %dms (mean %d ms).\n",
		int(r.wallsTimer.getMeasuredPassedTime()/time.Millisecond),
		int(r.wallsTimer.getMeanPassedTime()/time.Millisecond),
		int(r.floorCeilingTimer.getMeasuredPassedTime()/time.Millisecond),
		int(r.floorCeilingTimer.getMeanPassedTime()/time.Millisecond),
	)

	r.thingsTimer.newMeasure()
	r.thingsTimer.measure(func() { r.renderThings() })
	debugPrintf("Things rendered in %d ms.\n", int(r.thingsTimer.getMeasuredPassedTime()/time.Millisecond))

	elapsedTime := int(time.Since(startTimeTotal) / time.Millisecond)
	if elapsedTime != 0 {
		debugPrintf("Frame rendered in %d ms. (~ %d FPS) \n", elapsedTime, 1000/elapsedTime)
	} else {
		debugPrintf("Frame rendered in 0 ms. Yay! \n")
	}
}
