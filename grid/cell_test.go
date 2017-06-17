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
	if c.X() != CellX || c.Y() != CellY {
		t.Errorf("error: unexpected cell coordinates")
	}
	if c.Width() != CellWidth || c.Height() != CellHeight {
		t.Errorf("error: unexpected cell dimensions")
	}

	if c.PosX() != int32(CellX*CellWidth) || c.PosY() != int32(CellY*CellHeight) {
		t.Errorf("error: unexpected cell position")
	}

	if c.Rect().X != c.PosX() || c.Rect().Y != c.PosY() {
		t.Errorf("error: unexpected rect position")
	}
	if c.Rect().W != int32(CellWidth) || c.Rect().H != int32(CellHeight) {
		t.Errorf("error: unexpected rect size")
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
