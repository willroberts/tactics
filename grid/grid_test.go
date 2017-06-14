package grid

import (
	"testing"
)

const (
	GridWidth  int = 20
	GridHeight int = 10
)

func TestGrid(t *testing.T) {
	// Test that creating a grid produces the desired width and height.
	g := NewGrid(GridWidth, GridHeight, CellWidth, CellHeight)
	if g.Width() != GridWidth {
		t.FailNow()
	}
	if g.Height() != GridHeight {
		t.FailNow()
	}
	if g.CellWidth() != CellWidth {
		t.FailNow()
	}
	if g.CellHeight() != CellHeight {
		t.FailNow()
	}

	// Test retrieval of all cells.
	cols := g.Cells()
	if len(cols) != g.Width() {
		t.FailNow()
	}
	for i := range cols {
		if len(cols[i]) != g.Height() {
			t.FailNow()
		}
	}

	// Test that a cell's coordinates match the requested cell.
	FiveSeven := g.Cell(5, 7)
	if FiveSeven.X() != 5 {
		t.FailNow()
	}
	if FiveSeven.Y() != 7 {
		t.FailNow()
	}

	// Tests checkerColor and Checkerboard.
	var color1 uint32 = 0xffff0000
	var color2 uint32 = 0xff0000ff
	if checkerColor(color1, color2, 1, 3) != color1 {
		t.FailNow()
	}
	if checkerColor(color1, color2, 1, 4) != color2 {
		t.FailNow()
	}
	g.Checkerboard(color1, color2)
	if g.Cell(1, 3).Color() != color1 {
		t.FailNow()
	}
	if g.Cell(1, 4).Color() != color2 {
		t.FailNow()
	}
}
