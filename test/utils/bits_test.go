package utils_test

import (
	"BikeCH/src/utils"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const RANDOM_ITERATIONS = 1000
const INTEGER_SIZE = 32

func TestBits_ExtractThrowsWithInvalidStart(t *testing.T) {
	for _, start := range []int{-1, 32} {
		t.Run("", func(t *testing.T) {
			assertPanics(t, func() {
				utils.ExtractUnsigned(0, start, 1)
			})
			assertPanics(t, func() {
				utils.ExtractSigned(0, start, 1)
			})
		})
	}
}

func TestBits_ExtractThrowsWithInvalidLength(t *testing.T) {
	for _, length := range []int{-1, 32, 33} {
		t.Run("", func(t *testing.T) {
			assertPanics(t, func() {
				utils.ExtractUnsigned(0, 10, length)
			})
			assertPanics(t, func() {
				utils.ExtractSigned(0, 10, length)
			})
		})
	}
}

func TestBits_ExtractWorksOnFullLength(t *testing.T) {
	rng := rand.New(rand.NewSource(0)) // Seed for reproducibility
	for i := 0; i < RANDOM_ITERATIONS; i++ {
		v := rng.Int31()
		assert.Equal(t, v, utils.ExtractSigned(v, 0, INTEGER_SIZE))
	}
	for i := 0; i < RANDOM_ITERATIONS; i++ {
		v := 1 + rng.Int31n(math.MaxInt32)
		assert.Equal(t, uint32(v), utils.ExtractUnsigned(v, 0, INTEGER_SIZE-1))
	}
}

func TestBits_ExtractWorksOnRandomValues(t *testing.T) {
	rng := rand.New(rand.NewSource(0)) // Seed for reproducibility
	for i := 0; i < RANDOM_ITERATIONS; i++ {
		value := rng.Int31()
		start := rng.Intn(INTEGER_SIZE - 1)
		length := rng.Intn(INTEGER_SIZE-start) + 1

		expectedU := (value >> start) & ((1 << length) - 1)
		mask := int32(1 << (length - 1))
		expectedS := (expectedU ^ mask) - mask

		assert.Equal(t, uint32(expectedU), utils.ExtractUnsigned(value, start, length))
		assert.Equal(t, expectedS, utils.ExtractSigned(value, start, length))
	}
}

func assertPanics(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic, but the function completed successfully")
		}
	}()
	f()
}
