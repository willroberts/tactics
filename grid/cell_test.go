package grid

import "testing"

func TestCell(t *testing.T) {
	// Test that creating a new cell produces the desired coordinates.
	c := NewCell(1, 2, 20, 30)
	if c.X() != 1 {
		t.FailNow()
	}
	if c.Y() != 2 {
		t.FailNow()
	}
	if c.Width() != 20 {
		t.FailNow()
	}
	if c.Height() != 30 {
		t.FailNow()
	}

	// Test reassignment of coordinates.
	c.SetX(9)
	if c.X() != 9 {
		t.FailNow()
	}
	c.SetY(5)
	if c.Y() != 5 {
		t.FailNow()
	}

	// Test assignment of color.
	var colorRed uint32 = 0xffff0000
	c.SetColor(colorRed)
	if c.Color() != colorRed {
		t.FailNow()
	}
}
