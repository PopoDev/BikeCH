package projection

import (
	"math"
)

const (
	O_E = 2600000.0 // Origin E
	O_N = 1200000.0 // Origin N
)

// E calculates the East coordinates in terms of longitude (lon) and latitude (lat) in the WGS84 system.
// latitudes and longitudes are in radians.
func E(lon, lat float64) float64 {
	lambda1 := 1e-4 * (3600*radiansToDegrees(lon) - 26782.5)
	phi1 := 1e-4 * (3600*radiansToDegrees(lat) - 169028.66)

	E := 2600072.37 +
		211455.93*lambda1 -
		10938.51*lambda1*phi1 -
		0.36*lambda1*math.Pow(phi1, 2) -
		44.54*math.Pow(lambda1, 3)

	return E
}

// N calculates the North coordinates in terms of longitude (lon) and latitude (lat) in the WGS84 system.
// latitudes and longitudes are in radians.
func N(lon, lat float64) float64 {
	lambda1 := 1e-4 * (3600*radiansToDegrees(lon) - 26782.5)
	phi1 := 1e-4 * (3600*radiansToDegrees(lat) - 169028.66)

	N := 1200147.07 +
		308807.95*phi1 +
		3745.25*math.Pow(lambda1, 2) +
		76.63*math.Pow(phi1, 2) -
		194.56*lambda1*math.Pow(phi1, 3) +
		119.79*math.Pow(phi1, 3)

	return N
}

// Lon calculates the longitude in radians in terms of east coordinate (e) and north coordinate (n).
func Lon(e, n float64) float64 {
	x := 1e-6 * (e - O_E)
	y := 1e-6 * (n - O_N)

	lambda0 := 2.6779094 +
		4.728982*x +
		0.791484*x*y +
		0.1306*x*math.Pow(y, 2) -
		0.0436*math.Pow(x, 3)

	lambda := lambda0 * (100.0 / 36.0)

	return degreesToRadians(lambda)
}

// Lat calculates the latitude in radians in terms of east coordinate (e) and north coordinate (n).
func Lat(e, n float64) float64 {
	x := 1e-6 * (e - O_E)
	y := 1e-6 * (n - O_N)

	phi0 := 16.9023892 +
		3.238272*y -
		0.270978*math.Pow(x, 2) -
		0.002528*math.Pow(y, 2) -
		0.0447*x*math.Pow(y, 3) -
		0.0140*math.Pow(y, 3)

	phi := phi0 * (100.0 / 36.0)

	return degreesToRadians(phi)
}

func radiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
