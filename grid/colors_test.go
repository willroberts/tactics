package grid

import "testing"

func TestColorAssignment(t *testing.T) {
	g := NewGrid(3, 3, 25, 25)
	AssignColors(g)
	if g.Cell(0, 0).Color() != colorGreen {
		t.FailNow()
	}
	if g.Cell(0, 1).Color() != colorWhite {
		t.FailNow()
	}
	if g.Cell(0, 2).Color() != colorGreen {
		t.FailNow()
	}
	if g.Cell(1, 0).Color() != colorWhite {
		t.FailNow()
	}
	if g.Cell(1, 1).Color() != colorGreen {
		t.FailNow()
	}
	if g.Cell(1, 2).Color() != colorWhite {
		t.FailNow()
	}
	if g.Cell(2, 0).Color() != colorGreen {
		t.FailNow()
	}
	if g.Cell(2, 1).Color() != colorWhite {
		t.FailNow()
	}
	if g.Cell(2, 2).Color() != colorGreen {
		t.FailNow()
	}
}
