package data_test

import (
	"BikeCH/src/data"
	_ "BikeCH/testing_init"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGraphLoadFromWorksOnLausanneData(t *testing.T) {
	graph, err := data.LoadFrom("test/data/lausanne")
	require.NoError(t, err)

	// Check that nodes.bin was properly loaded
	actual1 := graph.NodeCount()
	expected1 := 212679
	assert.Equal(t, expected1, actual1)

	actual2 := graph.NodeOutEdgeID(2022, 0)
	expected2 := 4095
	assert.Equal(t, expected2, actual2)

	// Check that edges.bin was properly loaded
	actual3 := graph.EdgeLength(2022)
	expected3 := 17.875
	assert.Equal(t, expected3, actual3)

	// Check that profile_ids.bin and elevations.bin was properly loaded
	actual4 := graph.EdgeProfile(2022)(0)
	expected4 := 625.5625
	assert.Equal(t, expected4, actual4)

	// Check that attributes.bin and elevations.bin was properly loaded
	actual5 := int(graph.EdgeAttributes(2022).Bits)
	expected5 := 16
	assert.Equal(t, expected5, actual5)
}
