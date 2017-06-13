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
	if assignCheckerboardColor(0, 0) != colorGreen {
		t.FailNow()
	}
	if assignCheckerboardColor(0, 1) != colorWhite {
		t.FailNow()
	}
	if assignCheckerboardColor(0, 2) != colorGreen {
		t.FailNow()
	}
	if assignCheckerboardColor(1, 0) != colorWhite {
		t.FailNow()
	}
	if assignCheckerboardColor(1, 1) != colorGreen {
		t.FailNow()
	}
	if assignCheckerboardColor(1, 2) != colorWhite {
		t.FailNow()
	}
}

func TestDrawing(t *testing.T) {
	g := NewGrid(5, 5, 50, 50)
	if err := drawGrid(g); err != nil {
		t.FailNow()
	}
}

func assignCheckerboardColor(x, y int) uint32 {
	if x%2 == 0 {
		if y%2 == 0 {
			return colorGreen
		} else {
			return colorWhite
		}
	} else {
		if y%2 == 0 {
			return colorWhite
		} else {
			return colorGreen
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
			// Draw a black outline.
			outer := g.Cell(x, y).Rect()
			err = surface.FillRect(outer, 0xff000000)
			if err != nil {
				return err
			}

			// Draw the colored cell.
			inner := &sdl.Rect{
				X: outer.X + 2,
				Y: outer.Y + 2,
				W: outer.W - 4,
				H: outer.H - 4,
			}
			color := assignCheckerboardColor(x, y)
			err = surface.FillRect(inner, color)
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
