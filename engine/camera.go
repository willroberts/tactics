package engine

import "github.com/veandco/go-sdl2/sdl"

type Camera interface {
	MoveTo(int32, int32)
	Position() (int32, int32)
	ShiftRect(*sdl.Rect) *sdl.Rect
	ShiftVectors([]int16, []int16) ([]int16, []int16)
}

type camera struct {
	x int32
	y int32
}

func (c *camera) MoveTo(x, y int32) {
	c.x = x
	c.y = y
}

func (c *camera) Position() (int32, int32) {
	return c.x, c.y
}

func (c *camera) ShiftRect(r *sdl.Rect) *sdl.Rect {
	r.X = r.X - c.x
	r.Y = r.Y - c.y
	return r
}

func (c *camera) ShiftVectors(vx []int16, vy []int16) ([]int16, []int16) {
	for i, x := range vx {
		vx[i] = x - int16(c.x)
	}
	for i, y := range vy {
		vy[i] = y - int16(c.y)
	}
	return vx, vy
}

func NewCamera() Camera {
	return &camera{x: 0, y: 0}
}
