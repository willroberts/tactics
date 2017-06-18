package grid

import "github.com/veandco/go-sdl2/sdl"

// CellDimensions stores the visual representation data for a cell. X and Y
// values are the number of pixels from (0,0). W and H are the pixel dimensions
// of the cell.
type CellDimensions struct {
	X int
	Y int
	W int
	H int
}

// Cell is a container at a specific position inside a Grid.
type Cell interface {
	GridX() int
	GridY() int

	Dimensions() CellDimensions

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

func (c *cell) GridX() int {
	return c.x
}

func (c *cell) GridY() int {
	return c.y
}

func (c *cell) Dimensions() CellDimensions {
	return CellDimensions{
		X: c.x * c.width,
		Y: c.y * c.height,
		W: c.width,
		H: c.height,
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
