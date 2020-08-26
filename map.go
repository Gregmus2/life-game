package main

import "math"

type Map struct {
	factory                        *ObjectFactory
	units                          [][]*Unit
	size                           int
	xOffset, yOffset, halfUnitSize int
	unitSize                       float32
	centerX, centerY               int
}

func CopyFromMap(m *Map) *Map {
	units := make([][]*Unit, m.size)
	for i := range units {
		units[i] = make([]*Unit, m.size)
		copy(units[i], m.units[i])
	}

	return &Map{
		factory:      m.factory,
		size:         m.size,
		units:        units,
		xOffset:      m.xOffset,
		yOffset:      m.yOffset,
		halfUnitSize: m.halfUnitSize,
		unitSize:     m.unitSize,
	}
}

func NewMap(factory *ObjectFactory, w, h, size int) *Map {
	units := make([][]*Unit, size)
	for i := 0; i < size; i++ {
		units[i] = make([]*Unit, size)
	}

	min := w
	if h < w {
		min = h
	}
	unitSize := float32(math.Floor(float64(min / size)))
	xOffset := (w - int(unitSize)*size) / 2
	yOffset := (h - int(unitSize)*size) / 2
	halfUnitSize := int(unitSize / 2)

	return &Map{
		factory:      factory,
		units:        units,
		size:         size,
		xOffset:      xOffset,
		yOffset:      yOffset,
		halfUnitSize: halfUnitSize,
		unitSize:     unitSize,
		centerX:      size / 2,
		centerY:      size / 2,
	}
}

func (m *Map) normalizeCoords(x, y int) (int, int) {
	/*if x < 0 || y < 0 {
		originalSize := m.size
		m.size *= 2
		units := make([][]*Unit, m.size)
		for i := 0; i < originalSize; i++ {
			units[i] = make([]*Unit, m.size)
		}
		for i := originalSize; i < m.size; i++ {
			units[i] = make([]*Unit, m.size)
			copy(units[i][originalSize:], m.units[i])
		}
		m.units = units
	}
	if x > m.size || y > m.size {
		originalSize := m.size
		m.size *= 2
		units := make([][]*Unit, m.size)
		for i := 0; i < originalSize; i++ {
			units[i] = make([]*Unit, m.size)
			copy(units[i][:originalSize], m.units[i])
		}
		for i := originalSize; i < m.size; i++ {
			units[i] = make([]*Unit, m.size)
		}
		m.units = units
	}*/

	if x < 0 {
		x += m.size
	}
	if y < 0 {
		y += m.size
	}
	if x >= m.size {
		x -= m.size
	}
	if y >= m.size {
		y -= m.size
	}

	return x, y
}

func (m *Map) AddUnitSlice(slice [][]uint8) []*Unit {
	units := make([]*Unit, 0)
	sliceLen := len(slice)
	for x, ySlice := range slice {
		ySliceLen := len(ySlice)
		for y, flag := range ySlice {
			if flag == 1 {
				units = append(units, m.AddUnit(sliceLen-x+m.size/2, ySliceLen-y+m.size/2))
			}
		}
	}

	return units
}

func (m *Map) AddUnit(x, y int) *Unit {
	if m.units[y][x] != nil {
		return m.units[y][x]
	}

	m.units[y][x] = m.factory.NewUnit(
		m.yOffset+int(m.unitSize)*y+m.halfUnitSize,
		m.xOffset+int(m.unitSize)*x+m.halfUnitSize,
		m.unitSize-2, m.unitSize-2,
	)

	return m.units[y][x]
}

func (m *Map) KillUnit(x, y int) *Unit {
	unit := m.units[y][x]
	m.units[y][x] = nil

	return unit
}

func (m *Map) CalcNeighbourhoods(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			x1, y1 := m.normalizeCoords(x+i, y+j)
			if m.units[y1][x1] != nil {
				count++
			}
		}
	}

	return count
}
