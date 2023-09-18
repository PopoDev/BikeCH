package projection

const MIN_E = 2485000.0 // MIN_E is the smallest E coordinate in Switzerland (2,485,000).
const MAX_E = 2834000.0 // MAX_E is the biggest E coordinate in Switzerland (2,834,000).

const MIN_N = 1075000.0 // MIN_N is the smallest N coordinate in Switzerland (1,075,000).
const MAX_N = 1296000.0 // MAX_N is the biggest N coordinate in Switzerland (1,296,000).

const WIDTH = MAX_E - MIN_E  // WIDTH is the width of Switzerland in meters (349,000).
const HEIGHT = MAX_N - MIN_N // HEIGHT is the height of Switzerland in meters (221,000).

// ContainsEN checks if the coordinates (in terms of N and E) are located in Switzerland.
func ContainsEN(e, n float64) bool {
	return (MIN_E <= e && e <= MAX_E) && (MIN_N <= n && n <= MAX_N)
}
