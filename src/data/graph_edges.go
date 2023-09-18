package data

import (
	"BikeCH/src/utils"
)

const (
	BYTE_SIZE     = 8
	SHORT_SIZE    = 16
	SHORT_BYTES   = 2
	INTEGER_BYTES = 4
)

const (
	OFFSET_DIRECTION_AND_TARGET = 0
	OFFSET_LENGTH               = OFFSET_DIRECTION_AND_TARGET + INTEGER_BYTES // 4
	OFFSET_ELEVATION            = OFFSET_LENGTH + SHORT_BYTES                 // 6
	OFFSET_OSM                  = OFFSET_ELEVATION + SHORT_BYTES              // 8
	EDGES_BYTES                 = OFFSET_OSM + SHORT_BYTES                    // 10
	PROFILE_INTS                = 1
	TYPE1_BITS                  = SHORT_SIZE
	TYPE2_BITS                  = BYTE_SIZE
	TYPE3_BITS                  = 4
)

// GraphEdges represents the array of all edges of the graph.
type GraphEdges struct {
	edgesBuffer []int8
	profileIds  []int32
	elevations  []int16
}

// NewGraphEdges creates a new GraphEdges instance with the given buffer.
func NewGraphEdges(edgesBuffer []int8, profileIds []int32, elevations []int16) *GraphEdges {
	return &GraphEdges{edgesBuffer: edgesBuffer, profileIds: profileIds, elevations: elevations}
}

// IsInverted reports whether the given edge is inverted.
func (g *GraphEdges) IsInverted(edgeId int) bool {
	index := edgeId*EDGES_BYTES + OFFSET_DIRECTION_AND_TARGET
	direction := g.edgesBuffer[index] // 1st bit = direction --> check sign

	return direction < 0
}

// TargetNodeId returns the identity of the target node of the given edge.
func (g *GraphEdges) TargetNodeID(edgeId int) int {
	index := edgeId*EDGES_BYTES + OFFSET_DIRECTION_AND_TARGET
	targetNodeId := int(g.edgesBuffer[index])<<24 | int(g.edgesBuffer[index+1])<<16 | int(g.edgesBuffer[index+2])<<8 | int(g.edgesBuffer[index+3])

	if g.IsInverted(edgeId) {
		return ^targetNodeId
	}

	return targetNodeId
}

// Length returns the length of the given edge.
func (g *GraphEdges) Length(edgeId int) float64 {
	lengthQ28_4 := g.LengthQ28_4(edgeId)
	return utils.AsDouble(lengthQ28_4)
}

// ElevationGain returns the elevation gain of the given edge.
func (g *GraphEdges) ElevationGain(edgeId int) float64 {
	index := edgeId*EDGES_BYTES + OFFSET_ELEVATION
	elevationQ28_4 := g.GetValueQ28_4(index)

	return utils.AsDouble(elevationQ28_4)
}

// HasProfile reports whether the given edge has a profile.
func (g *GraphEdges) HasProfile(edgeId int) bool {
	return g.ProfileType(edgeId) != 0
}

// ProfileSamples returns the profile samples of the given edge.
func (g *GraphEdges) ProfileSamples(edgeId int) []float32 {
	if !g.HasProfile(edgeId) {
		return []float32{}
	}

	numberSamples := 1 + utils.CeilDiv(int(g.LengthQ28_4(edgeId)), int(utils.OfInt(2)))
	firstSampleId := g.FirstSampleId(edgeId)

	firstSampleQ28_4 := uint32(g.elevations[firstSampleId])
	firstSample := utils.AsFloat(int32(firstSampleQ28_4))

	profileSamples := make([]float32, numberSamples)
	profileSamples[0] = firstSample

	var sampleBitsNumber int
	switch g.ProfileType(edgeId) {
	case 1:
		sampleBitsNumber = TYPE1_BITS // Uncompressed profile | 16 bits | UQ12.4
	case 2:
		sampleBitsNumber = TYPE2_BITS // Compressed Q4.4 profile | 8 bits | Q4.4
	case 3:
		sampleBitsNumber = TYPE3_BITS // Compressed Q0.4 profile | 4 bits | Q0.4
	default:
		panic("Unexpected profile type")
	}

	samplesByShort := SHORT_SIZE / sampleBitsNumber

	for i := 1; i < numberSamples; i++ {
		index := 1 + (i-1)/samplesByShort
		samplesInShort := g.elevations[firstSampleId+index]

		sampleIndex := (i-1)%samplesByShort + 1
		start := SHORT_SIZE - sampleIndex*sampleBitsNumber

		var q28_4 int32
		if sampleBitsNumber == TYPE1_BITS {
			q28_4 = int32(utils.ExtractUnsigned(int32(samplesInShort), start, sampleBitsNumber))
		} else {
			q28_4 = int32(utils.ExtractSigned(int32(samplesInShort), start, sampleBitsNumber))
		}

		sample := utils.AsFloat(q28_4)
		profileSamples[i] = sample
		if sampleBitsNumber != TYPE1_BITS {
			profileSamples[i] += profileSamples[i-1]
		}
	}

	if g.IsInverted(edgeId) {
		reverseArray(profileSamples)
	}
	return profileSamples
}

// AttributesIndex returns the index of the attributes of the given edge.
func (g *GraphEdges) AttributesIndex(edgeId int) int {
	index := edgeId*EDGES_BYTES + OFFSET_OSM
	return int(g.GetValueQ28_4(index))
}

// GetValueQ28_4 returns the value in Q28_4 format at the given index.
func (g *GraphEdges) GetValueQ28_4(index int) int32 {
	value := int32(g.edgesBuffer[index])<<8 | int32(g.edgesBuffer[index+1])

	return value
}

// LengthQ28_4 returns the length in Q28_4 format of the given edge.
func (g *GraphEdges) LengthQ28_4(edgeId int) int32 {
	index := edgeId*EDGES_BYTES + OFFSET_LENGTH
	return g.GetValueQ28_4(index)
}

// ProfileType returns the profile type of the given edge.
func (g *GraphEdges) ProfileType(edgeId int) int {
	profile := g.profileIds[edgeId*PROFILE_INTS]
	return int(utils.ExtractUnsigned(profile, 30, 2))
}

// FirstSampleId returns the first sample id of the given edge.
func (g *GraphEdges) FirstSampleId(edgeId int) int {
	profile := g.profileIds[edgeId*PROFILE_INTS]
	return int(utils.ExtractUnsigned(profile, 0, 30))
}

// reverseArray reverses the given array.
func reverseArray(array []float32) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
