package grid

import (
	"testing"
)

const (
	GridWidth  int = 20
	GridHeight int = 10
	CellWidth  int = 30
	CellHeight int = 40
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
