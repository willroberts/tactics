package grid

import (
	"testing"
)

const (
	GRID_WIDTH  int = 20
	GRID_HEIGHT int = 10
)

func TestGrid(t *testing.T) {
	// Test that creating a grid produces the desired width and height.
	g := NewGrid(GRID_WIDTH, GRID_HEIGHT)
	if g.Width() != GRID_WIDTH {
		t.FailNow()
	}
	if g.Height() != GRID_HEIGHT {
		t.FailNow()
	}

	// Test that a cell's coordinates match the requested cell.
	FiveSeven := g.Cell(5, 7)
	if FiveSeven.X() != 5 {
		t.FailNow()
	}
	if FiveSeven.Y() != 7 {
		t.FailNow()
	}

	// Test that the grid reassigns coordinates when updating a cell.
	g.SetCell(4, 2, FiveSeven)
	FourTwo := g.Cell(4, 2)
	if FourTwo.X() != 4 {
		t.FailNow()
	}
	if FourTwo.Y() != 2 {
		t.FailNow()
	}
}
