// +build graphics

package grid

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	colorWhite uint32 = 0xffffffff
	colorGreen uint32 = 0xff00ff00
	colorRed   uint32 = 0xffff0000

	windowDuration uint32 = 3000
)

func TestColorAssignment(t *testing.T) {
	g := NewGrid(3, 3, 25, 25)
	assignColors(g)
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

func TestDrawing(t *testing.T) {
	g := NewGrid(5, 5, 50, 50)
	assignColors(g)
	if err := drawGrid(g); err != nil {
		t.FailNow()
	}
}

// assignColors populates a grid's cells with a checkerboard color scheme.
func assignColors(g Grid) {
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			if x%2 == 0 {
				if y%2 == 0 {
					g.Cell(x, y).SetColor(colorGreen)
				} else {
					g.Cell(x, y).SetColor(colorWhite)
				}
			} else {
				if y%2 == 0 {
					g.Cell(x, y).SetColor(colorWhite)
				} else {
					g.Cell(x, y).SetColor(colorGreen)
				}
			}
		}
	}
}

func drawGrid(g Grid) error {
	window, err := sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		g.Width()*g.CellWidth(),
		g.Height()*g.CellHeight(),
		sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		return err
	}

	if err = renderer.Clear(); err != nil {
		return err
	}

	// Draw every cell
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			c := g.Cell(x, y)
			err = surface.FillRect(c.Rect(), c.Color())
			if err != nil {
				return err
			}
		}
	}

	if err = window.UpdateSurface(); err != nil {
		return err
	}

	sdl.Delay(windowDuration)

	return nil
}
