package utils

// DoubleUnaryOperator is a function that takes a double and returns a double.
type DoubleUnaryOperator func(float64) float64

// Constant represents a constant function.
type constant struct {
	value float64
}

// ApplyAsDouble applies the constant function to the input.
func (c constant) ApplyAsDouble(x float64) float64 {
	return c.value
}

// Sampled represents a function obtained by linear interpolation between samples.
type sampled struct {
	samples []float32
	xMax    float64
}

// ApplyAsDouble applies the sampled function to the input.
func (s sampled) ApplyAsDouble(x float64) float64 {
	// x out of bounds
	if x < 0 {
		return float64(s.samples[0])
	}
	if x >= s.xMax {
		return float64(s.samples[len(s.samples)-1])
	}

	nbInterval := len(s.samples) - 1
	sizeInterval := s.xMax / float64(nbInterval)

	leftBoundIndex := int(x / sizeInterval)
	leftBoundY := float64(s.samples[leftBoundIndex])
	rightBoundY := float64(s.samples[leftBoundIndex+1])

	leftBoundX := float64(leftBoundIndex) * sizeInterval
	dx := (x - leftBoundX) / sizeInterval

	return Interpolate(leftBoundY, rightBoundY, dx)
}

// Constant returns a constant function, whose value is always y.
func Constant(y float64) DoubleUnaryOperator {
	return constant{value: y}.ApplyAsDouble
}

// Sampled returns a function obtained by linear interpolation between samples.
func Sampled(samples []float32, xMax float64) DoubleUnaryOperator {
	CheckArgument(xMax > 0, "xMax must be positive")
	CheckArgument(len(samples) > 1, "samples must contain at least 2 elements")

	samplesCopy := make([]float32, len(samples))
	copy(samplesCopy, samples)

	return sampled{samples: samplesCopy, xMax: xMax}.ApplyAsDouble
}
