package raycaster

import (
	"math"
)

type r_column struct {
	x                  int     // on-screen X coord
	perpWallDist       float64 // the distance to camera PLANE, not to the camera itself
	elevation, slide   float64
	side               uint8
	hitTileX, hitTileY int
	deferred           bool
}

func (r *Renderer) renderWalls() {
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
		column := r_column{x: x}
		deferredColumn := r_column{x: x, deferred: false}
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
				if r.scene.GetTileSlideAmount(mapX, mapY) != 0 {
					if side == EW {
						wallX = camY + rayToScreenLength*rayDirectionY
					} else {
						wallX = camX + rayToScreenLength*rayDirectionX
					}
					wallX -= math.Floor(wallX)
					if wallX < r.scene.GetTileSlideAmount(mapX, mapY) {
						continue
					}
				}
				// sliding tiles code ended
				if r.scene.GetTileElevation(mapX, mapY) == 0 {
					hit = true
					column.hitTileX, column.hitTileY = mapX, mapY
					column.slide = r.scene.GetTileSlideAmount(mapX, mapY)

				} else {
					deferredColumn.hitTileX, deferredColumn.hitTileY = mapX, mapY
					deferredColumn.elevation = r.scene.GetTileElevation(mapX, mapY)
					deferredColumn.deferred = true
					deferredColumn.perpWallDist = rayToScreenLength
					deferredColumn.side = side
					deferredColumn.slide = r.scene.GetTileSlideAmount(mapX, mapY)
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
