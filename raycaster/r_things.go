package raycaster

import (
	"raylib-raycaster/middleware"
)

func (r *Renderer) renderThings() {
	camx, camy := r.cam.getCoordsWithOffset()
	things := r.scene.GetListOfThings()
	for node := things.Front(); node != nil; node = node.Next() {
		t := node.Value.(Thing)
		tx, ty := t.GetCoords()
		// check if the Sprite is faced by Camera
		xRelative, yRelative := tx-camx, ty-camy
		invDet := 1.0 / (r.cam.planeX*r.cam.dirY - r.cam.dirX*r.cam.planeY) // vector projection fuckery
		transformX := invDet * (r.cam.dirY*xRelative - r.cam.dirX*yRelative)
		transformY := invDet * (-r.cam.planeY*xRelative + r.cam.planeX*yRelative)
		if transformY < 0.01 { // close enough to zero
			continue
		}

		osx := int((float64(r.RenderWidth) / 2) * (1 + transformX/transformY))
		osy := r.RenderHeight / 2
		osw := int(float64(r.RenderWidth) / transformY)
		osh := int(r.aspectFactor * float64(r.RenderHeight) / transformY)
		if osw > r.RenderWidth {
			continue
		}

		// render the Sprite column-wise, like a Texture
		currSprite := t.GetSprite()
		for x := 0; x < osw; x++ {
			screenXCoord := x + osx - osw/2
			if screenXCoord < 0 || screenXCoord > r.RenderWidth-1 || r.rayDistancesBuffer[screenXCoord] < transformY {
				continue
			}
			spriteX := x * currSprite.w / osw

			for y := 0; y < osh; y++ {
				spriteY := (y * currSprite.h / osh) % currSprite.h
				_, _, _, a := currSprite.bitmap.At(spriteX, spriteY).RGBA()
				if a == 0 {
					continue
				}
				r.setFoggedColorFromBitmapPixelAtCoords(currSprite.bitmap, spriteX, spriteY, transformY)
				middleware.DrawPoint(int32(x+osx-osw/2), int32(y+osy-osh/2)+int32(r.cam.vBobOffset))
			}
		}
	}
}
