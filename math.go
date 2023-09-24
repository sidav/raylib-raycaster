package main

import "math"

func makeUnitVector(fx, fy, tx, ty float64) (float64, float64) {
	tx -= fx
	ty -= fy
	length := math.Sqrt(tx*tx + ty*ty)
	return tx / length, ty / length
}

func areFloatCoordsInRange(fx, fy, tx, ty, rang float64) bool {
	return (tx-fx)*(tx-fx)+(ty-fy)*(ty-fy) <= rang*rang
}

func rotateVectorRandomly(vx, vy, maxDegrees float64) (float64, float64) {
	randomDegrees := (float64(rnd.Rand(100001))/100000)*maxDegrees - maxDegrees/2
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
		randomDegrees += (float64(rnd.Rand(100001))/100000)*maxDegrees - maxDegrees/2
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
