package main

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type Unit struct {
	x, y  int
	w, h  float32
	prog  uint32
	shape *graphics.ShapeHelper
}

func (f *ObjectFactory) NewUnit(x, y int, w, h float32) *Unit {
	return &Unit{
		x: x, y: y, h: h, w: w,
		prog:  f.Prog.GetByColor(graphics.White()),
		shape: f.Shape,
	}
}

func (u *Unit) Draw(scale float32) error {
	gl.UseProgram(u.prog)
	u.shape.Box(float32(u.x)*scale-u.w*scale/2, float32(u.y)*scale+u.h*scale/2, u.w*scale, u.h*scale)

	return nil
}
