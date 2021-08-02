package internal

import (
	"github.com/beefsack/go-astar"
	"math"
)

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0, 2)
	for _, offset := range [][]int32{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.m.Tile(t.X+offset[0], t.Y+offset[1]); n != nil &&
			n.Kind != Wall {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *Tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	switch value := to.(type) {
	case *Tile:
		return math.Abs(float64(value.X-t.X)) + math.Abs(float64(value.Y-t.Y))
	}

	// todo replace with logger (warning)
	panic("wrong pather type")
}
