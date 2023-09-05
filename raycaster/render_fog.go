package raycaster

import (
	"image"
)

func (r *Renderer) setFoggedColorFromBitmapPixelAtCoords(bitmap image.Image, x, y int, distance float64) {
	// fade out calculation
	fadeoutFraction := 1 - (distance / r.RayLengthForMaximumFog)
	if fadeoutFraction < (1 - r.MaxFogFraction) {
		fadeoutFraction = 1 - r.MaxFogFraction
	}

	red, g, b, _ := bitmap.At(x, y).RGBA()

	rbyte := uint8(fadeoutFraction*float64(uint8(red)) + (1-fadeoutFraction)*float64(r.FogR))
	gbyte := uint8(fadeoutFraction*float64(uint8(g)) + (1-fadeoutFraction)*float64(r.FogG))
	bbyte := uint8(fadeoutFraction*float64(uint8(b)) + (1-fadeoutFraction)*float64(r.FogB))

	r.backend.SetColor(rbyte, gbyte, bbyte)
}
