package grid

type Cell interface {
	X() int
	SetX(int)
	Y() int
	SetY(int)

	Width() int
	SetWidth(int)
	Height() int
	SetHeight(int)

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

func NewCell(x, y, w, h int) Cell {
	return &cell{x: x, y: y, width: w, height: h}
}
