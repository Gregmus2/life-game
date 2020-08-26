package main

import (
	"github.com/Gregmus2/simple-engine/scenes"
	"time"
)

type Life struct {
	scenes.Base
	factory *ObjectFactory
	m       *Map
}

func NewLife(base scenes.Base, f *ObjectFactory) *Life {
	return &Life{
		Base:    base,
		factory: f,
	}
}

func (l *Life) Init() {
	w, h := l.Window.GetSize()
	l.m = NewMap(l.factory, w, h, 100)

	units := l.m.AddUnitSlice([][]uint8{
		{0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1},
	})

	for _, unit := range units {
		l.DrawObjects.Put(unit)
	}
}

func (l *Life) Update() {
	time.Sleep(100 * time.Millisecond)
	newMap := CopyFromMap(l.m)
	for x := 0; x < l.m.size; x++ {
		for y := 0; y < l.m.size; y++ {
			neighbourhoods := l.m.CalcNeighbourhoods(x, y)

			switch {
			case neighbourhoods < 2 || neighbourhoods > 3:
				l.DrawObjects.Remove(newMap.KillUnit(x, y))
			case neighbourhoods == 3:
				l.DrawObjects.Put(newMap.AddUnit(x, y))
			}
		}
	}
	l.m = newMap
}
