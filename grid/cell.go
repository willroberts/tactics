package grid

import "github.com/veandco/go-sdl2/sdl"

type Cell interface {
	// Phase 1 (Display):

	Rect() *sdl.Rect
	Texture() *sdl.Texture
	SetTexture(t *sdl.Texture)
	Draw()

	// Phase 2 (Logic):

	// Methods for showing contents (player, NPCs, items, etc).
	// Methods for adding new contents.
	// Methods for deleting and clearing contents.
	// Methods for walkability (false for boulders, NPCs, etc).

	// Phase 3 (Isometric):

	// Methods for getting and setting elevation.
}

type cell struct {
	rect    *sdl.Rect
	texture *sdl.Texture
}

func (c *cell) Rect() *sdl.Rect {
	return c.rect
}

func (c *cell) Texture() *sdl.Texture {
	return c.texture
}

func (c *cell) SetTexture(t *sdl.Texture) {
	c.texture = t
}

func (c *cell) Draw() {
	return
}

func NewCell() Cell {
	return &cell{}
}
