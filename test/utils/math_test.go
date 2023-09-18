package utils_test

import (
	"BikeCH/src/utils"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const delta = 1e-7
const randomIterations = 100

func TestCeilDivThrowsOnNegativeX(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		utils.CeilDiv(-1, 2)
	})
}

func TestCeilDivThrowsOnZeroY(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		utils.CeilDiv(1, 0)
	})
}

func TestCeilDivWorksOnPositiveValues(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		x := rng.Intn(1000)
		y := rng.Intn(1000) + 1
		expected := int(math.Ceil(float64(x) / float64(y)))
		actual := utils.CeilDiv(x, y)
		assert.Equal(expected, actual)
	}
}

func TestInterpolateWorksWith0(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		v1 := rng.Float64()*2000 - 1000
		v2 := rng.Float64()*2000 - 1000
		assert.InDelta(v1, utils.Interpolate(v1, v2, 0), delta)
	}
}

func TestInterpolateWorksWith1(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		v1 := rng.Float64()*2000 - 1000
		v2 := rng.Float64()*2000 - 1000
		assert.InDelta(v2, utils.Interpolate(v1, v2, 1), delta)
	}
}

func TestInterpolateWorksWithRandomValues(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		y0 := rng.Float64()*2000 - 1000
		y1 := rng.Float64()*2000 - 1000
		x := rng.Float64()*40 - 20
		y := utils.Interpolate(y0, y1, x)
		expectedSlope := y1 - y0
		actualSlope := (y - y0) / x
		assert.InDelta(expectedSlope, actualSlope, delta)
	}
}

func TestClampIntClampsValueBelowMin(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Intn(200000) - 100000
		max := min + rng.Intn(100000)
		v := min - rng.Intn(500)
		assert.Equal(min, utils.ClampInt(min, v, max))
	}
}

func TestClampIntClampsValueAboveMax(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Intn(200000) - 100000
		max := min + rng.Intn(100000)
		v := max + rng.Intn(500)
		assert.Equal(max, utils.ClampInt(min, v, max))
	}
}

func TestClampIntPreservesValuesBetweenMinAndMax(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Intn(200000) - 100000
		v := min + rng.Intn(100000)
		max := v + rng.Intn(100000)
		assert.Equal(v, utils.ClampInt(min, v, max))
	}
}

func TestClampDoubleClampsValueBelowMin(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Float64()*200000 - 100000
		max := min + rng.Float64()*100000
		v := min - rng.Float64()*500
		assert.InDelta(min, utils.ClampFloat64(min, v, max), delta)
	}
}

func TestClampDoubleClampsValueAboveMax(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Float64()*200000 - 100000
		max := min + rng.Float64()*100000
		v := max + rng.Float64()*500
		assert.InDelta(max, utils.ClampFloat64(min, v, max), delta)
	}
}

func TestClampDoublePreservesValuesBetweenMinAndMax(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		min := rng.Float64()*200000 - 100000
		v := min + rng.Float64()*100000
		max := v + rng.Float64()*100000
		assert.InDelta(v, utils.ClampFloat64(min, v, max), delta)
	}
}

func TestDotProductOfAVectorWithItselfIsSquaredNorm(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		x := rng.Float64()*200 - 100
		y := rng.Float64()*200 - 100
		expected := x*x + y*y
		assert.InDelta(expected, utils.DotProduct(x, y, x, y), delta)
	}
}

func TestDotProductOfOrthogonalVectorsIsZero(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		x := rng.Float64()*200 - 100
		y := rng.Float64()*200 - 100
		assert.InDelta(0, utils.DotProduct(x, y, -y, x), delta)
		assert.InDelta(0, utils.DotProduct(x, y, y, -x), delta)
	}
}

func TestSquaredNormIsEqualToSquaredHypot(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		uX := rng.Float64()*2000 - 1000
		uY := rng.Float64()*2000 - 1000
		norm := math.Hypot(uX, uY)
		assert.InDelta(norm*norm, utils.SquaredNorm(uX, uY), delta)
	}
}

func TestNormIsEqualToHypot(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(0))

	for i := 0; i < randomIterations; i++ {
		uX := rng.Float64()*2000 - 1000
		uY := rng.Float64()*2000 - 1000
		norm := math.Hypot(uX, uY)
		assert.InDelta(norm, utils.Norm(uX, uY), delta)
	}
}

func TestProjectionLengthWorksOnKnownValues(t *testing.T) {
	assert := assert.New(t)

	actual1 := utils.ProjectionLength(1, 2, 3, 4, 5, 6)
	expected1 := 5.65685424949238
	assert.InDelta(expected1, actual1, delta)

	actual2 := utils.ProjectionLength(1, 1, 2, 3, 5, 8)
	expected2 := 8.049844718999243
	assert.InDelta(expected2, actual2, delta)
}
