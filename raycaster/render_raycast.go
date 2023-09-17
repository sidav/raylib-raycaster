package raycaster

import (
	"math"
)

// corresponds to each on-screen column
type castedRay struct {
	x                     int     // on-screen X coord
	perpWallDist          float64 // the distance to camera PLANE, not to the camera itself
	vertSlide, horizSlide float64 // vertical and horizontal slide amount for hit tile
	side                  uint8
	hitTileX, hitTileY    int

	// "True" if the ray hit wall, but was cast further anyway.
	// There will be two castedRay structs: one for first hit (deferred=true), and one for final hit (deferred=false).
	// This may occur with, for example, half-transparent walls.
	// Also needed for walls with vertical slide.
	deferred bool
}

func (r *Renderer) castRays() {
	for x := 0; x < r.RenderWidth; x++ {
		cameraX := 2*float64(x)/float64(r.RenderWidth) - 1
		rayDirectionX := r.cam.dirX + r.cam.planeX*cameraX
		rayDirectionY := r.cam.dirY + r.cam.planeY*cameraX

		camX, camY := r.cam.getCoords()
		mapX, mapY := r.cam.getIntCoords()

		var sideDistX, sideDistY, rayToScreenLength float64
		deltaDistX := math.Abs(1 / rayDirectionX)
		deltaDistY := math.Abs(1 / rayDirectionY)

		var stepX, stepY int
		var side uint8

		hit := false

		if rayDirectionX < 0 {
			stepX = -1
			sideDistX = (camX - float64(mapX)) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (float64(mapX) + 1.0 - camX) * deltaDistX
		}
		if rayDirectionY < 0 {
			stepY = -1
			sideDistY = (camY - float64(mapY)) * deltaDistY
		} else {
			stepY = 1
			sideDistY = (float64(mapY) + 1.0 - camY) * deltaDistY
		}

		// tracing the ray
		column := castedRay{x: x}
		deferredColumn := castedRay{x: x, deferred: false}
		for !hit {
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = EW
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = NS
			}
			if side == EW {
				rayToScreenLength = (float64(mapX) - camX + (1-float64(stepX))/2) / rayDirectionX
			} else {
				rayToScreenLength = (float64(mapY) - camY + (1-float64(stepY))/2) / rayDirectionY
			}
			// break tracing of the ray is too long
			if r.MaxRayLength != 0 && rayToScreenLength > r.MaxRayLength {
				return
			}
			// ray is out of map bounds
			if !r.scene.AreGridCoordsValid(mapX, mapY) {
				continue
			}

			// hit detected
			if r.scene.IsTileOpaque(mapX, mapY) {
				// THIN WALLS BAD CODE
				if r.scene.IsTileThin(mapX, mapY) {
					if side == EW {
						if sideDistX <= sideDistY {
							rayToScreenLength += deltaDistX / 2
						} else {
							nperpWallDist := (float64(mapY+stepY) - camY + (1-float64(stepY))/2) / rayDirectionY
							xCoordOfNextIntersection := camX + (nperpWallDist)*rayDirectionX
							xCoordOfNextIntersection -= math.Floor(xCoordOfNextIntersection)
							if stepX < 0 && xCoordOfNextIntersection <= 0.5 ||
								stepX > 0 && xCoordOfNextIntersection >= 0.5 {
								rayToScreenLength += deltaDistX / 2
							} else {
								continue
							}
						}
					} else {
						if sideDistY <= sideDistX {
							rayToScreenLength += deltaDistY / 2
						} else {
							nperpWallDist := (float64(mapX+stepX) - camX + (1-float64(stepX))/2) / rayDirectionX
							xCoordOfNextIntersection := camY + (nperpWallDist)*rayDirectionY
							xCoordOfNextIntersection -= math.Floor(xCoordOfNextIntersection)
							if stepY < 0 && xCoordOfNextIntersection <= 0.5 ||
								stepY > 0 && xCoordOfNextIntersection >= 0.5 {
								rayToScreenLength += deltaDistY / 2
							} else {
								continue
							}
						}

					}
				}
				// THIN WALLS BAD CODE ENDED

				// sliding tiles code
				var wallX float64
				if r.scene.GetTileHorizontalSlide(mapX, mapY) != 0 {
					if side == EW {
						wallX = camY + rayToScreenLength*rayDirectionY
					} else {
						wallX = camX + rayToScreenLength*rayDirectionX
					}
					wallX -= math.Floor(wallX)
					if wallX < r.scene.GetTileHorizontalSlide(mapX, mapY) {
						continue
					}
				}
				// sliding tiles code ended
				if r.scene.GetTileVerticalSlide(mapX, mapY) == 0 {
					hit = true
					column.hitTileX, column.hitTileY = mapX, mapY
					column.horizSlide = r.scene.GetTileHorizontalSlide(mapX, mapY)

				} else {
					deferredColumn.hitTileX, deferredColumn.hitTileY = mapX, mapY
					deferredColumn.vertSlide = r.scene.GetTileVerticalSlide(mapX, mapY)
					deferredColumn.deferred = true
					deferredColumn.perpWallDist = rayToScreenLength
					deferredColumn.side = side
					deferredColumn.horizSlide = r.scene.GetTileHorizontalSlide(mapX, mapY)
				}
			}
		}

		column.perpWallDist = rayToScreenLength
		column.side = side
		r.drawColumn(&column, rayDirectionX, rayDirectionY)
		if deferredColumn.deferred {
			r.drawColumn(&deferredColumn, rayDirectionX, rayDirectionY)
		}
	}
}
