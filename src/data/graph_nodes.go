package data

import (
	"BikeCH/src/utils"
)

// GraphNodes represents the array of all nodes of the graph.
type GraphNodes struct {
	buffer []int32
}

const (
	offsetE        = 0
	offsetN        = offsetE + 1
	offsetOutEdges = offsetN + 1
	nodeInts       = offsetOutEdges + 1
	outEdgesStart  = 28
	outEdgesLength = 4
	idEdgesStart   = 0
	idEdgesLength  = 28
)

// NewGraphNodes creates a new GraphNodes instance with the given buffer.
func NewGraphNodes(buffer []int32) *GraphNodes {
	return &GraphNodes{buffer: buffer}
}

// Count returns the total number of nodes.
func (gn *GraphNodes) Count() int {
	return len(gn.buffer) / nodeInts
}

// NodeE returns the East coordinate of the given node (nodeId).
func (gn *GraphNodes) NodeE(nodeId int) float64 {
	index := nodeId*nodeInts + offsetE
	q28_4 := gn.buffer[index]
	return utils.AsDouble(q28_4)
}

// NodeN returns the North coordinate of a given node (nodeId).
func (gn *GraphNodes) NodeN(nodeId int) float64 {
	index := nodeId*nodeInts + offsetN
	q28_4 := gn.buffer[index]
	return utils.AsDouble(q28_4)
}

// OutDegree returns the number of outgoing edges given a particular node (nodeId).
func (gn *GraphNodes) OutDegree(nodeId int) int {
	index := nodeId*nodeInts + offsetOutEdges
	bit := gn.buffer[index]
	return int(utils.ExtractUnsigned(bit, outEdgesStart, outEdgesLength))
}

// EdgeID returns the identity of the edgeIndex-th outgoing edge given a particular node (nodeId).
func (gn *GraphNodes) EdgeID(nodeId, edgeIndex int) int {
	if edgeIndex < 0 || edgeIndex >= gn.OutDegree(nodeId) {
		panic("Invalid edge index")
	}

	index := nodeId*nodeInts + offsetOutEdges
	bit := gn.buffer[index]
	firstEdgeID := int(utils.ExtractUnsigned(bit, idEdgesStart, idEdgesLength))

	return firstEdgeID + edgeIndex
}
