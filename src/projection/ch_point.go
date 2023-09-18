package projection

import (
	"BikeCH/src/utils"
	"fmt"
	"math"
)

// ChPoint represents an immutable point in the Swiss coordinate system (E, N).
type ChPoint struct {
	E float64
	N float64
}

// NewChPoint creates a new immutable ChPoint instance with the given coordinates (E, N).
func NewChPoint(e, n float64) ChPoint {
	utils.CheckArgument(ContainsEN(e, n), fmt.Sprintf("Coordinates (%v, %v) provided are not within Swiss limits", e, n))
	return ChPoint{E: e, N: n}
}

// SquaredDistanceTo returns the squared distance between two points.
func (p ChPoint) SquaredDistanceTo(that ChPoint) float64 {
	uX := p.E - that.E
	uY := p.N - that.N

	return uX*uX + uY*uY
}

// DistanceTo returns the distance between two points.
func (p ChPoint) DistanceTo(that ChPoint) float64 {
	squaredDistance := p.SquaredDistanceTo(that)
	return math.Sqrt(squaredDistance)
}

// Lon returns the longitude of the ChPoint.
func (p ChPoint) Lon() float64 {
	return Lon(p.E, p.N)
}

// Lat returns the latitude of the ChPoint.
func (p ChPoint) Lat() float64 {
	return Lat(p.E, p.N)
}
