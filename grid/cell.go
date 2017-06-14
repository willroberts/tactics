package grid

import "github.com/veandco/go-sdl2/sdl"

// TODO: Consider removing SDL coupling, or add GL coupling.

// Cell is a container at a specific position inside a Grid.
type Cell interface {
	X() int
	Y() int

	PosX() int32
	PosY() int32
	Width() int
	Height() int
	Rect() *sdl.Rect

	Color() uint32
	SetColor(uint32)
	Texture() *sdl.Texture
	SetTexture(*sdl.Texture)

	Elevation() int
	SetElevation(int)

	Contents() Occupier
	SetContents(Occupier)
	ClearContents()
	IsOccupied() bool
}

type cell struct {
	x         int
	y         int
	width     int
	height    int
	color     uint32
	texture   *sdl.Texture
	elevation int
	contents  Occupier
}

func (c *cell) X() int {
	return c.x
}

func (c *cell) Y() int {
	return c.y
}

func (c *cell) PosX() int32 {
	return int32(c.x * c.width)
}

func (c *cell) PosY() int32 {
	return int32(c.y * c.height)
}

func (c *cell) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: c.PosX(),
		Y: c.PosY(),
		W: int32(c.width),
		H: int32(c.height),
	}
}

func (c *cell) Color() uint32 {
	return c.color
}

func (c *cell) SetColor(color uint32) {
	c.color = color
}

func (c *cell) Texture() *sdl.Texture {
	return c.texture
}

func (c *cell) SetTexture(t *sdl.Texture) {
	c.texture = t
}

func (c *cell) Width() int {
	return c.width
}

func (c *cell) Height() int {
	return c.height
}

func (c *cell) Elevation() int {
	return c.elevation
}

func (c *cell) SetElevation(i int) {
	c.elevation = i
}

func (c *cell) Contents() Occupier {
	return c.contents
}

func (c *cell) SetContents(o Occupier) {
	c.contents = o
}

func (c *cell) ClearContents() {
	c.contents = nil
}

func (c *cell) IsOccupied() bool {
	return c.contents != nil
}

func newCell(x, y, w, h int) Cell {
	return &cell{x: x, y: y, width: w, height: h, elevation: 0}
}
