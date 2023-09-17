package raycaster

import "math"

type Camera struct {
	X, Y             float64
	dirX, dirY       float64
	planeX, planeY   float64
	movDirX, movDirY float64 // same as dirX/dirY but always of length 1

	vBobOffset, maxVBobOffset, vBobSpeed int
}

func CreateCamera(x, y, viewAngle float64, maxVBobOffset, vBobSpeed int) *Camera {
	// planeX := math.Tan(float64(viewAngle)*3.14159265358979323 / (2*180.0))
	cam := &Camera{
		X:             x,
		Y:             y,
		maxVBobOffset: maxVBobOffset,
		vBobSpeed:     vBobSpeed,
		movDirX:       1,
		movDirY:       0,
	}
	cam.ChangeViewWidth(viewAngle)
	return cam
}

func (c *Camera) getCoords() (float64, float64) {
	return c.X, c.Y
}

func (c *Camera) getIntCoords() (int, int) {
	cx, cy := c.getCoords()
	return int(cx), int(cy)
}

func (c *Camera) getSquareDistanceTo(x, y float64) float64 {
	return (x-c.X)*(x-c.X) + (y-c.Y)*(y-c.Y)
}

func (c *Camera) reset() {
	// reset the Camera dir
	c.dirX = 1
	c.movDirX = 1
	c.dirY = 0
	c.movDirY = 0
	c.planeX = 0 //-0.5
	c.planeY = 0.5
}

func (c *Camera) ChangeViewWidth(degrees float64) {
	// remember the camera rotation angle
	lastAngle := math.Atan2(-c.dirX, c.dirY)

	c.reset()

	// setting the degrees
	c.dirX = 1.0 / math.Tan(degrees*math.Pi/360.0)

	// return the Camera rotation angle
	c.Rotate(lastAngle)
}

func (c *Camera) Rotate(radians float64) {
	sin := math.Sin(radians)
	cos := math.Cos(radians)
	oldDirX := c.dirX
	c.dirX = c.dirX*cos - c.dirY*sin
	c.dirY = oldDirX*sin + c.dirY*cos
	oldMovDirX := c.movDirX
	c.movDirX = c.movDirX*cos - c.movDirY*sin
	c.movDirY = oldMovDirX*sin + c.movDirY*cos
	oldPlaneX := c.planeX
	c.planeX = c.planeX*cos - c.planeY*sin
	c.planeY = oldPlaneX*sin + c.planeY*cos
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
}

func (c *Camera) MoveByVector(vx, vy float64) {
	c.X += vx
	c.Y += vy
	c.vBobOffset += c.vBobSpeed
	if c.vBobOffset >= c.maxVBobOffset {
		c.vBobSpeed = -c.vBobSpeed
		c.vBobOffset = c.maxVBobOffset
	}
	if c.vBobOffset <= -c.maxVBobOffset {
		c.vBobSpeed = -c.vBobSpeed
		c.vBobOffset = -c.maxVBobOffset
	}
}
