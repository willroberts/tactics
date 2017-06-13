package grid

import (
	"testing"
)

func TestDrawing(t *testing.T) {
	g := NewGrid(5, 5, 50, 50)
	AssignColors(g)
	//drawGrid(g) // Drawing disabled to speed up tests.
}
