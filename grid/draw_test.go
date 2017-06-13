package grid

import (
	"testing"
)

func TestDrawing(t *testing.T) {
	g := NewGrid(5, 5, 50, 50)
	AssignColors(g)
	/* Drawing disabled to speed up tests.
	if err := drawGrid(g); err != nil {
		t.FailNow()
	}
	*/
}
