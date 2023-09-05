package raycaster

func (r *Renderer) renderFloorAndCeiling() {
	posX, posY := r.cam.getCoordsWithOffset()
	for y := 0; y < r.RenderHeight; y++ {
		// rayDir for leftmost ray (X = 0) and rightmost ray (X = W)
		rayDirX0 := r.cam.dirX - r.cam.planeX
		rayDirY0 := r.cam.dirY - r.cam.planeY
		rayDirX1 := r.cam.dirX + r.cam.planeX
		rayDirY1 := r.cam.dirY + r.cam.planeY

		// Current Y position compared to the center of the screen (the horizon)
		offsetFromCenterForFloor := y - r.RenderHeight/2 - r.cam.vBobOffset
		offsetFromCenterForCeiling := r.RenderHeight/2 - y + r.cam.vBobOffset

		// Vertical position of the Camera.
		floorPosZ := float64(r.RenderHeight/2) * r.aspectFactor // + float64(r.cam.vBobOffset)

		// Horizontal distance from the Camera to the floor for the current row.
		// 0.5 is the z position exactly in the middle between floor and ceiling.
		floorRowDistance := floorPosZ / float64(offsetFromCenterForFloor)
		ceilingRowDistance := floorPosZ / float64(offsetFromCenterForCeiling)
		// fmt.Printf("dist %f \n", floorRowDistance)

		// calculate the real world step vector we have to add for each X (parallel to Camera plane)
		// adding step by step avoids multiplications with a weight in the inner loop
		floorStepX := floorRowDistance * (rayDirX1 - rayDirX0) / float64(r.RenderWidth)
		floorStepY := floorRowDistance * (rayDirY1 - rayDirY0) / float64(r.RenderWidth)
		ceilingStepX := ceilingRowDistance * (rayDirX1 - rayDirX0) / float64(r.RenderWidth)
		ceilingStepY := ceilingRowDistance * (rayDirY1 - rayDirY0) / float64(r.RenderWidth)
		// fmt.Printf("sx %f, sy %f \n", floorStepX, floorStepY)

		// real world coordinates of the leftmost column. This will be updated as we step to the right.
		floorX := posX + floorRowDistance*rayDirX0
		floorY := posY + floorRowDistance*rayDirY0
		ceilingX := posX + ceilingRowDistance*rayDirX0
		ceilingY := posY + ceilingRowDistance*rayDirY0
		// fmt.Printf("fx %f, fy %f \n", floorX, floorY)

		for x := 0; x < r.RenderWidth; x++ {
			if y > r.RenderHeight/2+r.cam.vBobOffset {
				// the cell coord is simply got from the integer parts of floorX and floorY
				cellX := int(floorX)
				cellY := int(floorY)

				texture := r.scene.GetFloorTextureForCoords(cellX, cellY)
				texWidth := texture.W
				texHeight := texture.H

				// get the Texture coordinate from the fractional part
				tx := int(float64(texWidth) * (floorX - float64(cellX)))  // & (texWidth-1)
				ty := int(float64(texHeight) * (floorY - float64(cellY))) // & (texHeight-1)
				r.setFoggedColorFromBitmapPixelAtCoords(texture.Bitmap, tx, ty, floorRowDistance)
				r.backend.DrawPoint(int32(x), int32(y))
			} else if r.RenderCeilings {
				//ceiling
				cellX := int(ceilingX)
				cellY := int(ceilingY)
				texture := r.scene.GetCeilingTextureForCoords(cellX, cellY)
				texWidth := texture.W
				texHeight := texture.H
				tx := int(float64(texWidth) * (ceilingX - float64(cellX)))
				ty := int(float64(texHeight) * (ceilingY - float64(cellY)))
				r.setFoggedColorFromBitmapPixelAtCoords(texture.Bitmap, tx, ty, ceilingRowDistance)
				r.backend.DrawPoint(int32(x), int32(y))
			}

			floorX += floorStepX     // + 0.5
			floorY += floorStepY     // + 0.5
			ceilingX += ceilingStepX // + 0.5
			ceilingY += ceilingStepY // + 0.5
		}
	}
}
