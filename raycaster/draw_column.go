package raycaster

import (
	"math"
)

func (rend *Renderer) drawColumn(column *castedRay, rayDirectionX, rayDirectionY float64) {
	// drawing the pixels column
	camH := rend.cam.GetVerticalCoordWithBob()
	// columnHeight := int(float64(rend.RenderHeight) / column.perpWallDist * rend.aspectFactor)

	// lowestPixelY := columnHeight/2 + rend.RenderHeight/2
	lowestPixelY := int(float64(rend.RenderHeight)*(0.5-rend.aspectFactor*(0-camH)/column.perpWallDist)) + rend.cam.OnScreenVerticalOffset

	// highestPixelY := -columnHeight/2 + rend.RenderHeight/2
	highestPixelY := int(float64(rend.RenderHeight)*(0.5-rend.aspectFactor*(1-camH)/column.perpWallDist)) + rend.cam.OnScreenVerticalOffset
	columnHeight := lowestPixelY - highestPixelY
	// Note: highest and lowest PixelY both CAN be out of screen bounds (this behaviour is needed for texturing).

	// Vertical verticalSlideOffset for "slided up" walls
	var verticalSlideOffset int
	if column.vertSlide > 0 {
		newLowY := int(float64(rend.RenderHeight)*
			(0.5-rend.aspectFactor*(column.vertSlide-camH)/column.perpWallDist)) + rend.cam.OnScreenVerticalOffset
		verticalSlideOffset = lowestPixelY - newLowY
		lowestPixelY = newLowY
	}

	if !rend.ApplyTexturing {
		rend.wallsTimer.measure(func() {
			rend.drawColumnUntextured(column, lowestPixelY, highestPixelY)
		})
	} else {
		rend.wallsTimer.measure(func() {
			rend.drawColumnTextured(column, rayDirectionX, rayDirectionY, columnHeight, verticalSlideOffset, lowestPixelY, highestPixelY)
		})
		// Render floor/ceiling only once per column
		if !column.deferred {
			rend.floorCeilingTimer.measure(func() {
				rend.renderTexturedFloorAndCeilingColumn(column.x, lowestPixelY, highestPixelY)
			})
		}
	}
	rend.rayDistancesBuffer[column.x] = column.perpWallDist
}

func (rend *Renderer) drawColumnUntextured(column *castedRay, lowestPixelY, highestPixelY int) {
	if column.side == NS {
		rend.surface.verticalLine(column.x, lowestPixelY, highestPixelY, surfaceColor{128, 128, 128})
	} else {
		rend.surface.verticalLine(column.x, lowestPixelY, highestPixelY, surfaceColor{64, 64, 64})
	}
}

func (rend *Renderer) drawColumnTextured(column *castedRay, rayDirectionX, rayDirectionY float64, columnHeight, verticalSlideOffset, lowestPixelY, highestPixelY int) {
	camx, camy := rend.cam.getCoords()
	// TEXTURING
	texture := rend.scene.GetTextureForTile(column.hitTileX, column.hitTileY)
	texWidth := texture.W
	texHeight := texture.H

	var wallX float64 //where exactly the wall was hit
	if column.side == EW {
		wallX = camy + column.perpWallDist*rayDirectionY - column.horizSlide
	} else {
		wallX = camx + column.perpWallDist*rayDirectionX - column.horizSlide
	}
	wallX -= math.Floor(wallX)

	//X coordinate on the Texture
	texX := int(wallX * float64(texWidth))
	if column.side == EW && rayDirectionX > 0 {
		texX = texWidth - texX - 1
	}
	if column.side == NS && rayDirectionY < 0 {
		texX = texWidth - texX - 1
	}

	// How much to increase the Texture coordinate per screen pixel
	step := 1.0 * float64(texHeight) / float64(columnHeight)
	// Starting Texture coordinate
	// texPos := (float64(highestPixelY+offset-rend.cam.vBobOffset) - float64(rend.RenderHeight)/2 + float64(columnHeight)/2) * step
	texPos := float64(verticalSlideOffset) * step
	if highestPixelY < 0 {
		texPos += -float64(highestPixelY) * step
	}
	from := max(0, highestPixelY)
	to := min(rend.RenderHeight-1, lowestPixelY)
	for y := from; y <= to; y++ {
		// texY := int(texPos) % (texHeight-1) // analog of texPos % texHeight ONLY IF texHeight is a power of two!
		texY := int(texPos) % texHeight
		texPos += step
		// fmt.Printf("(%d,%d) OUT OF (%d,%d)\n", texX, texY, texWidth, texHeight)

		color := rend.setFoggedColorFromBitmapPixelAtCoords(texture.Bitmap, texX, texY, column.perpWallDist, column.side == NS)
		rend.surface.putPixel(column.x, y, color)
	}
}
