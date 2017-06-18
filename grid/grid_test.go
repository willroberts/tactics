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
	if g.Width() != GridWidth || g.Height() != GridHeight {
		t.Errorf("error: unexpected grid dimensions")
	}
	if g.CellWidth() != CellWidth || g.CellHeight() != CellHeight {
		t.Errorf("error: unexpected cell dimensions")
	}

	// Test retrieval of all cells.
	cols := g.Cells()
	if len(cols) != g.Width() {
		t.Errorf("error: unexpected number of columns")
	}
	for i := range cols {
		if len(cols[i]) != g.Height() {
			t.Errorf("error: unexpected number of rows")
		}
	}

	// Test that a cell's coordinates match the requested cell.
	FiveSeven := g.Cell(5, 7)
	if FiveSeven.GridX() != 5 || FiveSeven.GridY() != 7 {
		t.Errorf("error: failed to retrieve cell coordinates")
	}

	// Tests checkerColor and Checkerboard.
	var color1 uint32 = 0xffff0000
	var color2 uint32 = 0xff0000ff
	if checkerColor(color1, color2, 1, 3) != color1 {
		t.Errorf("error: checkerColor returned unexpected result color2")
	}
	if checkerColor(color1, color2, 1, 4) != color2 {
		t.Errorf("error: checkerColor returned unexpected result color1")
	}
	g.Checkerboard(color1, color2)
	if g.Cell(1, 3).Color() != color1 {
		t.Errorf("error: Checkerboard returned unexpected result color2")
	}
	if g.Cell(1, 4).Color() != color2 {
		t.Errorf("error: Checkerboard returned unexpected result color1")
	}
}
