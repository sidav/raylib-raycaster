package raycaster

import (
	"raylib-raycaster/middleware"
)

func (r *Renderer) renderThings() {
	camx, camy := r.cam.getCoordsWithOffset()
	things := r.scene.GetListOfThings()
	
	// sort by distance from camera (descending)
	// WARNING: breaks things list order!
	for node1 := things.Front(); node1 != nil && node1.Next() != nil; node1 = node1.Next() {

		t1 := node1.Value.(Thing)
		t1x, t1y := t1.GetCoords()
		dist1 := (camx - t1x) * (camx - t1x) + (camy - t1y) * (camy - t1y)

		for node2 := node1.Next(); node2 != nil; node2 = node2.Next() {

			t2 := node2.Value.(Thing)
			t2x, t2y := t2.GetCoords()
			dist2 := (camx - t2x) * (camx - t2x) + (camy - t2y) * (camy - t2y)

			if dist2 > dist1 {
				// swap 2 and 1
				things.MoveAfter(node1, node2)
			}
		}
	}
	
	for node := things.Front(); node != nil; node = node.Next() {
		// unneeded? 
		//if node.Value == nil {
		//	continue
		//}
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
			// r.rayDistancesBuffer[screenXCoord] = transformY
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
