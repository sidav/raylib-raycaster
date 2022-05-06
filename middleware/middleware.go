package middleware

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var (
	currColor     color.RGBA
	TargetTexture rl.RenderTexture2D
)

func SetInternalResolution(w, h int32) {
	TargetTexture = rl.LoadRenderTexture(w, h)
	rl.SetTextureFilter(TargetTexture.Texture, rl.FilterAnisotropic16x)
}

func Flush() {
	rl.BeginDrawing()
	rl.DrawTexturePro(TargetTexture.Texture, rl.Rectangle{
		X:      0,
		Y:      float32(TargetTexture.Texture.Height),
		Width:  float32(TargetTexture.Texture.Width),
		Height: float32(-TargetTexture.Texture.Height),
	},
	rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(rl.GetScreenWidth()),
		Height: float32(rl.GetScreenHeight()),
	},
	rl.Vector2{
		X: 0,
		Y: 0,
	},
	0,
	color.RGBA{255, 255, 255, 255})
	rl.EndDrawing()
}

func SetColor(r, g, b uint8) {
	currColor.R = r
	currColor.G = g
	currColor.B = b
	currColor.A = 255
	//currColor = color.RGBA{
	//	R: r,
	//	G: g,
	//	B: b,
	//	A: 255,
	//}
}

func DrawPoint(x, y int32) {
	rl.DrawPixel(x, y, currColor)
}

func FillRect(x, y, w, h int) {
	rl.DrawRectangle(int32(x), int32(y), int32(w), int32(h), currColor)
}

func VerticalLine(x, y0, y1 int) {
	rl.DrawLine(int32(x), int32(y0), int32(x), int32(y1), currColor)
}
