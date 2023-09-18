package utils

import (
	"fmt"
	"math"
)

// CeilDiv returns the ceiling of the division of two values.
func CeilDiv(x, y int) int {
	CheckArgument(x >= 0 && y > 0, fmt.Sprintf("x must be non-negative and y must be positive, but got x = %d and y = %d", x, y))
	return (x + y - 1) / y
}

// Interpolate computes the y-coordinate of the interpolated point.
func Interpolate(y0, y1, x float64) float64 {
	slope := y1 - y0
	return math.FMA(slope, x, y0)
}

// ClampInt limits the value v to the interval [min, max].
func ClampInt(min, v, max int) int {
	CheckArgument(min <= max, fmt.Sprintf("min must be less than or equal to max, but got min = %d and max = %d", min, max))
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// ClampFloat64 limits the value v to the interval [min, max].
func ClampFloat64(min, v, max float64) float64 {
	CheckArgument(min <= max, fmt.Sprintf("min must be less than or equal to max, but got min = %f and max = %f", min, max))
	return math.Max(min, math.Min(max, v))
}

// DotProduct computes the dot product between two vectors u and v.
func DotProduct(uX, uY, vX, vY float64) float64 {
	return uX*vX + uY*vY
}

// SquaredNorm computes the square of the norm of a vector.
func SquaredNorm(uX, uY float64) float64 {
	return DotProduct(uX, uY, uX, uY)
}

// Norm computes the norm of a vector.
func Norm(uX, uY float64) float64 {
	return math.Sqrt(SquaredNorm(uX, uY))
}

// ProjectionLength computes the length of the projection of vector AP on vector AB.
func ProjectionLength(aX, aY, bX, bY, pX, pY float64) float64 {
	uX := pX - aX
	uY := pY - aY
	vX := bX - aX
	vY := bY - aY
	dotProd := DotProduct(uX, uY, vX, vY)
	vNorm := Norm(vX, vY)
	return dotProd / vNorm
}
