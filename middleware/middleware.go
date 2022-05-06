package middleware

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var (
	currColor color.RGBA
)

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
