package grid

import "github.com/veandco/go-sdl2/sdl"

// Cell is a container at a specific position inside a Grid.
type Cell interface {
	// Grid coordinates.
	X() int
	Y() int

	// Spatial coordinates and properties within the window.
	PosX() int32
	PosY() int32
	Width() int
	Height() int
	Elevation() int
	Rect() *sdl.Rect

	// Logical properties.
	Contents() Occupier
	SetContents(Occupier)
	ClearContents()
	IsOccupied() bool

	// Methods for retrieving and setting Texture.
}

type cell struct {
	x         int
	y         int
	width     int
	height    int
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
	return int32(0 + (c.x * c.width))
}

func (c *cell) PosY() int32 {
	return int32(0 + (c.y * c.height))
}

func (c *cell) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: c.PosX(),
		Y: c.PosY(),
		W: int32(c.width),
		H: int32(c.height),
	}
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
