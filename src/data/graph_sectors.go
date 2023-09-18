package data

import (
	"BikeCH/src/projection"
	"BikeCH/src/utils"
	"encoding/binary"
)

const (
	NbSectorsPerSide     = 128
	SectorMin, SectorMax = 0, NbSectorsPerSide - 1

	SectorWidth  = projection.WIDTH / NbSectorsPerSide  // 2726.5625 [m]
	SectorHeight = projection.HEIGHT / NbSectorsPerSide // 1726.5625 [m]

	Offset1stNode = 0
	OffsetNbNodes = Offset1stNode + 4
	SectorsBytes  = OffsetNbNodes + 2 // 4 + 2 = 6 bytes
)

type GraphSectors struct {
	Buffer []byte
}

func NewGraphSectors(buffer []byte) *GraphSectors {
	return &GraphSectors{Buffer: buffer}
}

func (gs *GraphSectors) SectorsInArea(center projection.ChPoint, distance float64) []*Sector {
	e := center.E
	n := center.N

	// Swiss Bounds
	eMin := utils.ClampFloat64(projection.MIN_E, e-distance, e)
	eMax := utils.ClampFloat64(e, e+distance, projection.MAX_E)
	nMin := utils.ClampFloat64(projection.MIN_N, n-distance, n)
	nMax := utils.ClampFloat64(n, n+distance, projection.MAX_N)

	// Sector Bounds
	xMin := utils.ClampInt(SectorMin, gs.SectorX(eMin), SectorMax)
	xMax := utils.ClampInt(SectorMin, gs.SectorX(eMax), SectorMax)
	yMin := utils.ClampInt(SectorMin, gs.SectorY(nMin), SectorMax)
	yMax := utils.ClampInt(SectorMin, gs.SectorY(nMax), SectorMax)

	sectorsInArea := make([]*Sector, 0)

	for y := yMin; y <= yMax; y++ { // Vertical
		for x := xMin; x <= xMax; x++ { // Horizontal
			sectorID := gs.SectorID(x, y)

			firstNodeID := gs.FirstNodeID(sectorID)
			numberNodes := gs.NumberNodes(sectorID)
			endNodeID := firstNodeID + numberNodes
			sector := NewSector(firstNodeID, endNodeID)

			sectorsInArea = append(sectorsInArea, sector)
		}
	}

	return sectorsInArea
}

func (gs *GraphSectors) FirstNodeID(sectorID int) int {
	firstNodeIndex := sectorID*SectorsBytes + Offset1stNode
	return int(binary.BigEndian.Uint32(gs.Buffer[firstNodeIndex : firstNodeIndex+4]))
}

func (gs *GraphSectors) NumberNodes(sectorID int) int {
	numberNodesIndex := sectorID*SectorsBytes + OffsetNbNodes
	return int(binary.BigEndian.Uint32(gs.Buffer[numberNodesIndex : numberNodesIndex+4]))
}

func (gs *GraphSectors) SectorID(x, y int) int {
	return x + y*NbSectorsPerSide
}

func (gs *GraphSectors) SectorX(e float64) int {
	return int((e - projection.MIN_E) / SectorWidth)
}

func (gs *GraphSectors) SectorY(n float64) int {
	return int((n - projection.MIN_N) / SectorHeight)
}

type Sector struct {
	StartNodeID int
	EndNodeID   int
}

func NewSector(startNodeID, endNodeID int) *Sector {
	return &Sector{StartNodeID: startNodeID, EndNodeID: endNodeID}
}
