package grid

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	CellX      int = 1
	CellY      int = 2
	CellWidth  int = 20
	CellHeight int = 30
)

func TestCell(t *testing.T) {
	c := newCell(CellX, CellY, CellWidth, CellHeight)
	if c.GridX() != CellX || c.GridY() != CellY {
		t.Errorf("error: unexpected cell coordinates")
	}

	d := c.Dimensions()
	if d.W != CellWidth || d.H != CellHeight {
		t.Errorf("error: unexpected cell dimensions")
	}

	if d.X != CellX*CellWidth || d.Y != CellY*CellHeight {
		t.Errorf("error: unexpected cell position")
	}

	c.SetColor(0xff336699)
	if c.Color() != 0xff336699 {
		t.Errorf("error: failed to set color")
	}

	if c.Elevation() != 0 {
		t.Errorf("error: initial elevation was non-zero")
	}
	c.SetElevation(10)
	if c.Elevation() != 10 {
		t.Errorf("error: failed to set elevation")
	}

	o := c.Contents()
	if o != nil || c.IsOccupied() {
		t.Errorf("error: initial cell contents were not nil")
	}

	occ := &occupier{}
	c.SetContents(occ)
	o = c.Contents()
	if o != occ || !c.IsOccupied() {
		t.Errorf("error: failed to set cell contents")
	}

	c.ClearContents()
	o = c.Contents()
	if o != nil || c.IsOccupied() {
		t.Errorf("error: failed to clear cell contents")
	}

	tex := &sdl.Texture{}
	c.SetTexture(tex)
	if c.Texture() != tex {
		t.Errorf("error: failed to set cell texture")
	}
}
