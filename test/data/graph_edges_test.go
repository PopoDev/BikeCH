package data_test

import (
	"BikeCH/src/data"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphEdgesWorksOnGivenExample(t *testing.T) {
	assert := assert.New(t)

	edgesBuffer := []int8{
		^0, ^0, ^0, ^12, // Inverted, target node = 12
		0x01, 0x0b, // Length: 0x010.b = 16.6875 [m]
		0x01, 0x00, // Elevation gain: 0x010.0 = 16.0 [m]
		0x04, 0x10, // 1040
	}

	profileIds := []int32{
		// (3 << 30) | 1, // Type 3, first sample index = 1
		-1073741823,
	}

	elevations := []int16{
		0x0000, 0x180c, -257, -2, -4096,
	}

	edges := data.NewGraphEdges(edgesBuffer, profileIds, elevations)
	fmt.Printf("%08b\n", edgesBuffer)
	fmt.Printf("%08b\n", elevations)

	assert.True(edges.IsInverted(0))
	assert.Equal(12, edges.TargetNodeID(0))
	assert.Equal(16.6875, edges.Length(0))
	assert.Equal(16.0, edges.ElevationGain(0))
	assert.True(edges.HasProfile(0))
	assert.Equal(1040, edges.AttributesIndex(0))

	expectedSamples := []float32{
		384.0625, 384.125, 384.25, 384.3125, 384.375,
		384.4375, 384.5, 384.5625, 384.6875, 384.75,
	}

	samples := edges.ProfileSamples(0)
	for i := 0; i < len(samples); i++ {
		fmt.Printf("Sample %d: %f\n", i, samples[i])
	}

	assert.InDeltaSlice(expectedSamples, edges.ProfileSamples(0), 0.0001)

}
