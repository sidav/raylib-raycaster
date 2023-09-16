package raycaster

import "math"

func (rend *Renderer) drawColumn(column *r_column, rayDirectionX, rayDirectionY float64) {
	// drawing the pixels column
	columnHeight := int(float64(rend.RenderHeight) / column.perpWallDist * rend.aspectFactor)
	highestPixelY := -columnHeight/2 + rend.RenderHeight/2
	if highestPixelY < 0 {
		highestPixelY = 0
	}
	lowestPixelY := columnHeight/2 + rend.RenderHeight/2

	var offset int
	if column.elevation > 0 {
		offset = int(float64(rend.RenderHeight) / column.perpWallDist * column.elevation * rend.aspectFactor)
		lowestPixelY -= offset
	}

	if lowestPixelY >= rend.RenderHeight {
		lowestPixelY = rend.RenderHeight - 1
	}

	if !rend.ApplyTexturing {
		rend.drawColumnUntextured(column, highestPixelY, lowestPixelY)
	} else {
		rend.drawColumnTextured(column, rayDirectionX, rayDirectionY, columnHeight, offset, lowestPixelY, highestPixelY)
	}
	rend.rayDistancesBuffer[column.x] = column.perpWallDist
}

func (rend *Renderer) drawColumnUntextured(column *r_column, lowestPixelY, highestPixelY int) {
	rend.backend.SetColor(128, 128, 128)
	if column.side == NS {
		rend.backend.SetColor(255, 255, 255)
	}
	rend.backend.VerticalLine(column.x, lowestPixelY+rend.cam.vBobOffset, highestPixelY+rend.cam.vBobOffset)
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
	texPos := (float64(highestPixelY+offset) - float64(rend.RenderHeight)/2 + float64(columnHeight)/2) * step
	for y := highestPixelY; y < lowestPixelY; y++ {
		// texY := int(texPos) % (texHeight-1) // analog of texPos % texHeight ONLY IF texHeight is a power of two!
		texY := int(texPos) % texHeight
		texPos += step
		// fmt.Printf("(%d,%d) OUT OF (%d,%d)\n", texX, texY, texWidth, texHeight)

		rend.setFoggedColorFromBitmapPixelAtCoords(texture.Bitmap, texX, texY, column.perpWallDist, column.side == NS)
		rend.backend.DrawPoint(int32(column.x), int32(y+rend.cam.vBobOffset))
	}
}
