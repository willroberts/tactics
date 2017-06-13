package grid

import "github.com/veandco/go-sdl2/sdl"

// Cell is a container at a specific position inside a Grid.
type Cell interface {
	X() int
	SetX(int)
	Y() int
	SetY(int)

	PosX() int32
	PosY() int32
	Rect() *sdl.Rect

	Width() int
	SetWidth(int)
	Height() int
	SetHeight(int)

	Color() uint32
	SetColor(uint32)

	// Methods for retrieving and setting Rect and Texture.
	// Methods for drawing.

	// Methods for showing contents (player, NPCs, items, etc).
	// Methods for adding new contents.
	// Methods for deleting and clearing contents.
	// Methods for walkability (false for boulders, NPCs, etc).

	// Methods for getting and setting elevation (isometric).
}

type cell struct {
	x      int
	y      int
	width  int
	height int
	color  uint32
}

func (c *cell) X() int {
	return c.x
}

func (c *cell) SetX(i int) {
	c.x = i
}

func (c *cell) Y() int {
	return c.y
}

func (c *cell) SetY(i int) {
	c.y = i
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

func (c *cell) SetWidth(i int) {
	c.width = i
}

func (c *cell) Height() int {
	return c.height
}

func (c *cell) SetHeight(i int) {
	c.height = i
}

func (c *cell) Color() uint32 {
	return c.color
}

func (c *cell) SetColor(i uint32) {
	c.color = i
}

// NewCell initializes and returns a Cell.
func NewCell(x, y, w, h int) Cell {
	return &cell{x: x, y: y, width: w, height: h}
}
