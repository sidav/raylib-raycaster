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

func (s *surface) clear() {
	s.fillRect(0, 0, len(s.pixels), len(s.pixels[0]), surfaceColor{0, 0, 0})
}

func (s *surface) putPixel(x, y int, color surfaceColor) {
	s.pixels[x][y][0] = color.r
	s.pixels[x][y][1] = color.g
	s.pixels[x][y][2] = color.b
}

func (s *surface) fillRect(x, y, w, h int, color surfaceColor) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			s.putPixel(i, j, color)
		}
	}
}

func (s *surface) verticalLine(x, y0, y1 int, color surfaceColor) {
	for j := y0; j < y1; j++ {
		s.putPixel(x, j, color)
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
