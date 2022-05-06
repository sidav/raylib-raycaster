package raycaster

import "math"

type Camera struct {
	X, Y           float64
	dirX, dirY     float64
	planeX, planeY float64
	movDirX, movDirY float64 // same as dirX/dirY but always of length 1 

	vBobOffset, maxVBobOffset, vBobSpeed int
	
	hBobOffset, hBobSpeed, maxHBobOffset float64
}

func CreateCamera(x, y,  viewAngle, maxHBobOffset, hBobSpeed float64, maxVBobOffset, vBobSpeed int) *Camera {
	// planeX := math.Tan(float64(viewAngle)*3.14159265358979323 / (2*180.0))
	cam := &Camera{
		X:             x,
		Y:             y,
		maxVBobOffset: maxVBobOffset,
		vBobSpeed:     vBobSpeed,
		maxHBobOffset: maxHBobOffset,
		hBobSpeed:     hBobSpeed,
		movDirX: 0,
		movDirY: 1,
	}
	cam.ChangeViewWidth(viewAngle)
	return cam
}

func (c *Camera) getCoordsWithOffset() (float64, float64) {
	return c.X + c.dirY * c.hBobOffset, c.Y + c.dirX * c.hBobOffset
}

func (c *Camera) getIntCoords() (int, int) {
	cx, cy := c.getCoordsWithOffset()
	return int(cx), int(cy)
}

func (c *Camera) getSquareDistanceTo(x, y float64) float64 {
	return (x-c.X)*(x-c.X) + (y-c.Y)*(y-c.Y)
}

func (c *Camera) ChangeViewWidth(degrees float64) {
	// reset the Camera dir
	c.dirX = 0
	c.movDirX = 0
	c.dirY = 1
	c.movDirY = 1
	c.planeX = -0.5
	c.planeY = 0
	c.dirY = 1.0 / math.Tan(degrees * math.Pi / 360.0)
}

func (c *Camera) Rotate(radians float64) {
	oldDirX := c.dirX
	c.dirX = c.dirX*math.Cos(radians) - c.dirY*math.Sin(radians)
	c.dirY = oldDirX*math.Sin(radians) + c.dirY*math.Cos(radians)
	oldMovDirX := c.movDirX
	c.movDirX = c.movDirX*math.Cos(radians) - c.movDirY*math.Sin(radians)
	c.movDirY = oldMovDirX*math.Sin(radians) + c.movDirY*math.Cos(radians)
	oldPlaneX := c.planeX
	c.planeX = c.planeX*math.Cos(radians) - c.planeY*math.Sin(radians)
	c.planeY = oldPlaneX*math.Sin(radians) + c.planeY*math.Cos(radians)
}

func (c *Camera) MoveForward(fraction float64) {
	c.X += c.movDirX * fraction
	c.Y += c.movDirY * fraction
	c.vBobOffset += c.vBobSpeed
	if c.vBobOffset >= c.maxVBobOffset {
		c.vBobSpeed = -c.vBobSpeed
		c.vBobOffset = c.maxVBobOffset
	}
	if c.vBobOffset <= -c.maxVBobOffset {
		c.vBobSpeed = -c.vBobSpeed
		c.vBobOffset = -c.maxVBobOffset
	}

	c.hBobOffset += c.hBobSpeed
	if c.hBobOffset >= c.maxHBobOffset {
		c.hBobSpeed = -c.hBobSpeed
		c.hBobOffset = c.maxHBobOffset
	}
	if c.hBobOffset <= -c.maxHBobOffset {
		c.hBobSpeed = -c.hBobSpeed
		c.hBobOffset = -c.maxHBobOffset
	}
}
