package main

import "math"

func rotateVectorRandomly(vx, vy, maxDegrees float64) (float64, float64) {
	randomDegrees := rnd.Float64()*maxDegrees - maxDegrees/2
	radians := randomDegrees * 3.1415926 / 180.0
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	t := vx
	vx = vx*cos - vy*sin
	vy = t*sin + vy*cos
	return vx, vy
}

func rotateVectorRandomlyGauss(vx, vy, maxDegrees float64) (float64, float64) {
	const sums = 2
	randomDegrees := 0.0
	for i := 0; i < sums; i++ {
		randomDegrees += rnd.Float64()*maxDegrees - maxDegrees/2
	}
	randomDegrees /= sums
	radians := randomDegrees * 3.1415926 / 180.0
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	t := vx
	vx = vx*cos - vy*sin
	vy = t*sin + vy*cos
	return vx, vy
}
