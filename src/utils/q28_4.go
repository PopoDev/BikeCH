package utils

import (
	"math"
)

// Q28_4 is a utility class to convert numbers between the Q28.4 (int32) representation and other representations.
const (
	bitsAfterDecimalPoint = 4
	power2ScaleFactor     = -bitsAfterDecimalPoint
)

// OfInt returns a Q28.4 value corresponding to the given integer by doing a 4-bit left shift.
func OfInt(i int32) int32 {
	return i << bitsAfterDecimalPoint
}

// AsDouble returns a number of type double corresponding to the given Q28.4 number.
func AsDouble(q28_4 int32) float64 {
	return math.Pow(2, power2ScaleFactor) * float64(q28_4)
}

// AsFloat returns a number of type float corresponding to the given Q28.4 number.
func AsFloat(q28_4 int32) float32 {
	return float32(math.Pow(2, power2ScaleFactor) * float64(q28_4))
}
