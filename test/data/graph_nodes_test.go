package data_test

import (
	"BikeCH/src/data"
	"BikeCH/test"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphNodesWorksOnGivenExample(t *testing.T) {
	assert := assert.New(t)

	b := []int32{
		2_600_000 << 4,
		1_200_000 << 4,
		0x2_000_1234,
	}
	ns := data.NewGraphNodes(b)
	assert.Equal(1, ns.Count())
	assert.Equal(2_600_000.0, ns.NodeE(0))
	assert.Equal(1_200_000.0, ns.NodeN(0))
	assert.Equal(2, ns.OutDegree(0))
	assert.Equal(0x1234, ns.EdgeID(0, 0))
	assert.Equal(0x1235, ns.EdgeID(0, 1))
}

func TestGraphNodesCountWorksFrom0To99(t *testing.T) {
	assert := assert.New(t)

	for count := 0; count < 100; count++ {
		buffer := make([]int32, 3*count)
		graphNodes := data.NewGraphNodes(buffer)
		assert.Equal(count, graphNodes.Count())
	}
}

func TestGraphNodesENWorkOnRandomCoordinates(t *testing.T) {
	assert := assert.New(t)

	nodesCount := 10_000
	buffer := make([]int32, 3*nodesCount)
	rng := test.NewRandom()

	for i := 0; i < test.RANDOM_ITERATIONS; i++ {
		e := 2_600_000 + int32(50_000*rng.Float64())
		n := 1_200_000 + int32(50_000*rng.Float64())
		e_q28_4 := int32(float64(e) * math.Pow(2, 4))
		n_q28_4 := int32(float64(n) * math.Pow(2, 4))
		e = e_q28_4 / 16
		n = n_q28_4 / 16
		nodeId := rng.Intn(nodesCount)
		buffer[3*nodeId] = e_q28_4
		buffer[3*nodeId+1] = n_q28_4
		graphNodes := data.NewGraphNodes(buffer)
		assert.InDelta(e, graphNodes.NodeE(nodeId), test.DELTA)
		assert.InDelta(n, graphNodes.NodeN(nodeId), test.DELTA)
	}
}

func TestGraphNodesOutDegreeWorks(t *testing.T) {
	assert := assert.New(t)

	nodesCount := 10_000
	buffer := make([]int32, 3*nodesCount)
	rng := test.NewRandom()

	for outDegree := 0; outDegree < 16; outDegree++ {
		firstEdgeId := rng.Intn(1 << 28)
		nodeId := rng.Intn(nodesCount)
		buffer[3*nodeId+2] = int32(outDegree<<28 | firstEdgeId)
		graphNodes := data.NewGraphNodes(buffer)
		assert.Equal(outDegree, graphNodes.OutDegree(nodeId))
	}
}

func TestGraphNodesEdgeIdWorksOnRandomValues(t *testing.T) {
	assert := assert.New(t)

	nodesCount := 10_000
	buffer := make([]int32, 3*nodesCount)
	rng := test.NewRandom()

	for outDegree := 0; outDegree < 16; outDegree++ {
		firstEdgeId := rng.Intn(1 << 28)
		nodeId := rng.Intn(nodesCount)
		buffer[3*nodeId+2] = int32(outDegree<<28 | firstEdgeId)
		graphNodes := data.NewGraphNodes(buffer)
		for i := 0; i < outDegree; i++ {
			expectedEdgeId := firstEdgeId + i
			assert.Equal(expectedEdgeId, graphNodes.EdgeID(nodeId, i))
		}
	}
}
