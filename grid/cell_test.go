package grid

import (
	"testing"
)

const (
	CellX      int = 1
	CellY      int = 2
	CellWidth  int = 20
	CellHeight int = 30
)

func TestCell(t *testing.T) {
	// Test that creating a new cell produces the desired coordinates.
	c := newCell(CellX, CellY, CellWidth, CellHeight)
	if c.X() != CellX {
		t.FailNow()
	}
	if c.Y() != CellY {
		t.FailNow()
	}
	if c.Width() != CellWidth {
		t.FailNow()
	}
	if c.Height() != CellHeight {
		t.FailNow()
	}

	// Test position calculations.
	if c.PosX() != int32(CellX*CellWidth) {
		t.FailNow()
	}
	if c.PosY() != int32(CellY*CellHeight) {
		t.FailNow()
	}

	// Test Rect.
	if c.Rect().X != c.PosX() || c.Rect().Y != c.PosY() {
		t.FailNow()
	}
	if c.Rect().W != int32(CellWidth) || c.Rect().H != int32(CellHeight) {
		t.FailNow()
	}

	// Tests elevation.
	if c.Elevation() != 0 {
		t.FailNow()
	}
	c.SetElevation(10)
	if c.Elevation() != 10 {
		t.FailNow()
	}

	// Test assignment, retrieval, and clearing of cell contents.
	o := c.Contents()
	if o != nil || c.IsOccupied() {
		t.FailNow()
	}

	occ := &occupier{}
	c.SetContents(occ)
	o = c.Contents()
	if o != occ || !c.IsOccupied() {
		t.FailNow()
	}

	c.ClearContents()
	o = c.Contents()
	if o != nil || c.IsOccupied() {
		t.FailNow()
	}
}
