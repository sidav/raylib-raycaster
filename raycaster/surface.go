package raycaster

import (
	"sync"
)

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

func (s *surface) clear(threads int) {
	wg := &sync.WaitGroup{}
	for x := 0; x < len(s.pixels); x++ {
		wg.Add(1)

		go s.clearColumn(x, wg)

		if (x+1)%threads == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
}

func (s *surface) clearColumn(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	s.verticalLine(x, 0, len(s.pixels[0]), surfaceColor{0, 0, 0})
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

func (r *Renderer) PresentDrawnSurfaceToBackend() {
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
