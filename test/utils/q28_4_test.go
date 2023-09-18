package utils_test

import (
	"BikeCH/src/utils"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ28_4_OfIntWorksWithRandomValues(t *testing.T) {
	rng := rand.New(rand.NewSource(0)) // Seed for reproducibility
	for i := 0; i < RANDOM_ITERATIONS; i++ {
		n := rng.Int31n(1 << 28)
		assert.Equal(t, uint32(n), uint32(utils.OfInt(n))>>4)
	}
}

func TestQ28_4_AsDoubleWorksOnKnownValues(t *testing.T) {
	assert.Equal(t, 1.0, utils.AsDouble(0b1_0000))
	assert.Equal(t, 1.5, utils.AsDouble(0b1_1000))
	assert.Equal(t, 1.25, utils.AsDouble(0b1_0100))
	assert.Equal(t, 1.125, utils.AsDouble(0b1_0010))
	assert.Equal(t, 1.0625, utils.AsDouble(0b1_0001))
	assert.Equal(t, 1.9375, utils.AsDouble(0b1_1111))
}

func TestQ28_4_AsFloatWorksOnKnownValues(t *testing.T) {
	assert.Equal(t, float32(1.0), utils.AsFloat(0b1_0000))
	assert.Equal(t, float32(1.5), utils.AsFloat(0b1_1000))
	assert.Equal(t, float32(1.25), utils.AsFloat(0b1_0100))
	assert.Equal(t, float32(1.125), utils.AsFloat(0b1_0010))
	assert.Equal(t, float32(1.0625), utils.AsFloat(0b1_0001))
	assert.Equal(t, float32(1.9375), utils.AsFloat(0b1_1111))
}

func TestQ28_4_OfIntAndAsFloatDoubleAreInverse(t *testing.T) {
	rng := rand.New(rand.NewSource(0)) // Seed for reproducibility
	for i := 0; i < RANDOM_ITERATIONS; i++ {
		n := rng.Int31n(1 << 24)
		assert.Equal(t, n, int32(utils.AsFloat(utils.OfInt(n))))
		assert.Equal(t, n, int32(utils.AsDouble(utils.OfInt(n))))
	}
}
