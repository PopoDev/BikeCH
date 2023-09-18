package data

import (
	"BikeCH/src/projection"
	"BikeCH/src/utils"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

// Graph represents a graph with nodes, sectors, edges, and attribute sets.
type Graph struct {
	nodes         *GraphNodes
	sectors       *GraphSectors
	edges         *GraphEdges
	attributeSets []AttributeSet
}

// NewGraph constructs a graph with the given nodes, sectors, edges, and attribute sets.
func NewGraph(nodes *GraphNodes, sectors *GraphSectors, edges *GraphEdges, attributeSets []AttributeSet) *Graph {
	return &Graph{
		nodes:         nodes,
		sectors:       sectors,
		edges:         edges,
		attributeSets: append([]AttributeSet{}, attributeSets...), // Create a copy to ensure immutability
	}
}

// LoadFrom loads a Graph from the files located in basePath.
func LoadFrom(basePath string) (*Graph, error) {
	nodesPath := filepath.Join(basePath, "nodes.bin")
	sectorsPath := filepath.Join(basePath, "sectors.bin")
	edgesPath := filepath.Join(basePath, "edges.bin")
	profileIdsPath := filepath.Join(basePath, "profile_Ids.bin")
	elevationsPath := filepath.Join(basePath, "elevations.bin")
	attributesPath := filepath.Join(basePath, "attributes.bin")
	println(nodesPath)

	nodesBuffer := readFileAsInt32Buffer(nodesPath)
	sectorsBuffer := readFileAsByteBuffer(sectorsPath)
	edgesBuffer := readFileAsInt8Buffer(edgesPath)
	profileIdsBuffer := readFileAsInt32Buffer(profileIdsPath)
	elevationsBuffer := readFileAsShortBuffer(elevationsPath)
	attributesBuffer := readFileAsLongBuffer(attributesPath)

	println("profileIdsBuffer: ", len(profileIdsBuffer))

	nodes := NewGraphNodes(nodesBuffer)
	sectors := NewGraphSectors(sectorsBuffer)
	edges := NewGraphEdges(edgesBuffer, profileIdsBuffer, elevationsBuffer)
	attributeSets := getAttributeSets(attributesBuffer)

	return NewGraph(nodes, sectors, edges, attributeSets), nil
}

// NodeCount returns the total number of nodes in the graph.
func (g *Graph) NodeCount() int {
	return g.nodes.Count()
}

// NodePoint returns the position of the given identity node in the Swiss coordinate.
func (g *Graph) NodePoint(nodeID int) projection.ChPoint {
	return projection.NewChPoint(g.nodes.NodeE(nodeID), g.nodes.NodeN(nodeID))
}

// NodeOutDegree returns the number of edges going out from the given identity node.
func (g *Graph) NodeOutDegree(nodeID int) int {
	return g.nodes.OutDegree(nodeID)
}

// NodeOutEdgeID returns the identity of the edgeIndex-th edge outgoing from the given identity node.
func (g *Graph) NodeOutEdgeID(nodeID, edgeIndex int) int {
	return g.nodes.EdgeID(nodeID, edgeIndex)
}

// NodeClosestTo returns the identity of the node closest to the given point, at the given maximum distance (in meters).
func (g *Graph) NodeClosestTo(point projection.ChPoint, searchDistance float64) int {
	closest := -1
	min := searchDistance * searchDistance

	for _, sector := range g.sectors.SectorsInArea(point, searchDistance) {
		startNodeID := sector.StartNodeID
		endNodeID := sector.EndNodeID

		for id := startNodeID; id < endNodeID; id++ {
			squaredDistance := point.SquaredDistanceTo(g.NodePoint(id))
			if squaredDistance <= min {
				closest = id
				min = squaredDistance
			}
		}
	}

	return closest
}

// EdgeTargetNodeID returns the identity of the destination node that belongs to the edge.
func (g *Graph) EdgeTargetNodeID(edgeID int) int {
	return g.edges.TargetNodeID(edgeID)
}

// EdgeIsInverted returns true if the edge goes in the opposite direction to the OSM path.
func (g *Graph) EdgeIsInverted(edgeID int) bool {
	return g.edges.IsInverted(edgeID)
}

// EdgeAttributes returns the set of OSM attributes attached to the given identity edge.
func (g *Graph) EdgeAttributes(edgeID int) AttributeSet {
	return g.attributeSets[g.edges.AttributesIndex(edgeID)]
}

// EdgeLength returns the length, in meters, of the given identity edge.
func (g *Graph) EdgeLength(edgeID int) float64 {
	return g.edges.Length(edgeID)
}

// EdgeElevationGain returns the total positive elevation of the edge with the given identity.
func (g *Graph) EdgeElevationGain(edgeID int) float64 {
	return g.edges.ElevationGain(edgeID)
}

// EdgeProfile returns the profile of the given identity edge as a function.
func (g *Graph) EdgeProfile(edgeID int) utils.DoubleUnaryOperator {
	if !g.edges.HasProfile(edgeID) {
		println("No profile")
		return utils.Constant(math.NaN())
	}

	samples := g.edges.ProfileSamples(edgeID)
	xMax := g.edges.Length(edgeID)
	return utils.Sampled(samples, xMax)
}

// getAttributeSets extracts list of AttributeSet from a buffer.
func getAttributeSets(buffer []uint64) []AttributeSet {
	attributeSets := make([]AttributeSet, len(buffer))
	for i, val := range buffer {
		attributeSets[i] = NewAttributeSet(val)
	}
	return attributeSets
}

// readFileAsByteBuffer reads a binary file and returns its contents as a byte buffer.
func readFileAsByteBuffer(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	utils.CheckArgument(err == nil, fmt.Sprintf("Error while reading file %s", filePath))

	return data
}

// readFileAsShortBuffer reads a binary file and returns its contents as a short buffer.
func readFileAsShortBuffer(filePath string) []int16 {
	data := readFileAsByteBuffer(filePath)
	reader := bytes.NewReader(data)

	size := len(data) / SHORT_BYTES

	var buffer []int16 = make([]int16, size)
	binary.Read(reader, binary.BigEndian, &buffer)

	return buffer
}

// readFileAsIn8Buffer reads a binary file and returns its contents as a int8 buffer.
func readFileAsInt8Buffer(filePath string) []int8 {
	data := readFileAsByteBuffer(filePath)
	reader := bytes.NewReader(data)

	size := len(data)

	var buffer []int8 = make([]int8, size)
	binary.Read(reader, binary.BigEndian, &buffer)

	return buffer
}

// readFileAsInt32Buffer reads a binary file and returns its contents as a int32 buffer.
func readFileAsInt32Buffer(filePath string) []int32 {
	data := readFileAsByteBuffer(filePath)
	reader := bytes.NewReader(data)

	size := len(data) / INTEGER_BYTES

	var buffer []int32 = make([]int32, size)
	binary.Read(reader, binary.BigEndian, &buffer)

	return buffer
}

// readFileAsIntBuffer reads a binary file and returns its contents as an int buffer.
func readFileAsIntBuffer(filePath string) []int {
	data := readFileAsByteBuffer(filePath)
	reader := bytes.NewReader(data)

	size := len(data) / INTEGER_BYTES

	var buffer []int = make([]int, size)
	binary.Read(reader, binary.BigEndian, &buffer)

	return buffer
}

// readFileAsLongBuffer reads a binary file and returns its contents as a long buffer.
func readFileAsLongBuffer(filePath string) []uint64 {
	data := readFileAsByteBuffer(filePath)
	reader := bytes.NewReader(data)

	size := len(data) / 8

	var buffer []uint64 = make([]uint64, size)
	binary.Read(reader, binary.BigEndian, &buffer)

	return buffer
}
