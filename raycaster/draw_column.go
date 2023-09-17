package raycaster

import "math"

func (rend *Renderer) drawColumn(column *r_column, rayDirectionX, rayDirectionY float64) {
	// drawing the pixels column
	columnHeight := int(float64(rend.RenderHeight) / column.perpWallDist * rend.aspectFactor)

	lowestPixelY := columnHeight/2 + rend.RenderHeight/2
	var offset int
	if column.elevation > 0 {
		offset = int(float64(rend.RenderHeight) / column.perpWallDist * column.elevation * rend.aspectFactor)
		lowestPixelY -= offset
	}

	highestPixelY := -columnHeight/2 + rend.RenderHeight/2
	// Note: highest and lowest PixelY both CAN be out of screen bounds (this behaviour is needed for texturing).
	highestPixelY += rend.cam.vBobOffset
	lowestPixelY += rend.cam.vBobOffset

	if !rend.ApplyTexturing {
		rend.drawColumnUntextured(column, lowestPixelY, highestPixelY)
	} else {
		rend.drawColumnTextured(column, rayDirectionX, rayDirectionY, columnHeight, offset, lowestPixelY, highestPixelY)
		rend.renderTexturedFloorAndCeilingColumn(column.x, lowestPixelY, highestPixelY)
	}
	rend.rayDistancesBuffer[column.x] = column.perpWallDist
}

func (rend *Renderer) drawColumnUntextured(column *r_column, lowestPixelY, highestPixelY int) {
	rend.backend.SetColor(128, 128, 128)
	if column.side == NS {
		rend.backend.SetColor(255, 255, 255)
	}
	rend.backend.VerticalLine(column.x, lowestPixelY, highestPixelY)
}

func (rend *Renderer) drawColumnTextured(column *r_column, rayDirectionX, rayDirectionY float64, columnHeight, offset, lowestPixelY, highestPixelY int) {
	camx, camy := rend.cam.getCoordsWithOffset()
	// TEXTURING
	texture := rend.scene.GetTextureForTile(column.hitTileX, column.hitTileY)
	texWidth := texture.W
	texHeight := texture.H

	var wallX float64 //where exactly the wall was hit
	if column.side == EW {
		wallX = camy + column.perpWallDist*rayDirectionY - column.slide
	} else {
		wallX = camx + column.perpWallDist*rayDirectionX - column.slide
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
	texPos := 0.0
	if highestPixelY+offset < 0 {
		texPos = -float64(highestPixelY+offset) * step
	}
	from := max(0, highestPixelY)
	to := min(rend.RenderHeight-1, lowestPixelY)
	for y := from; y <= to; y++ {
		// texY := int(texPos) % (texHeight-1) // analog of texPos % texHeight ONLY IF texHeight is a power of two!
		texY := int(texPos) % texHeight
		texPos += step
		// fmt.Printf("(%d,%d) OUT OF (%d,%d)\n", texX, texY, texWidth, texHeight)

		rend.setFoggedColorFromBitmapPixelAtCoords(texture.Bitmap, texX, texY, column.perpWallDist, column.side == NS)
		rend.backend.DrawPoint(int32(column.x), int32(y))
	}
}
