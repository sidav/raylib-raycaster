package raycaster

type surfaceColor struct {
	r, g, b uint8
}

type surface struct {
	pixels [][][3]uint8
}

func (s *surface) create(w, h int) {
	s.pixels = make([][][3]uint8, w)
	for i := range s.pixels {
		s.pixels[i] = make([][3]uint8, h)
	}
}

func (r *Renderer) clearScreen() {
	r.fillRect(0, 0, r.RenderWidth, r.RenderHeight, surfaceColor{0, 0, 0})
}

func (r *Renderer) putPixel(x, y int, color surfaceColor) {
	r.surface.pixels[x][y][0] = color.r
	r.surface.pixels[x][y][1] = color.g
	r.surface.pixels[x][y][2] = color.b
}

func (r *Renderer) fillRect(x, y, w, h int, color surfaceColor) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			r.putPixel(i, j, color)
		}
	}
}

func (r *Renderer) verticalLine(x, y0, y1 int, color surfaceColor) {
	for j := y0; j < y1; j++ {
		r.putPixel(x, j, color)
	}
}

func (r *Renderer) ShowDrawnSurface() {
	for x := range r.surface.pixels {
		for y := range r.surface.pixels[x] {
			r.backend.SetColor(
				r.surface.pixels[x][y][0],
				r.surface.pixels[x][y][1],
				r.surface.pixels[x][y][2],
			)
			r.backend.DrawPoint(int32(x), int32(y))
		}
	}
}
