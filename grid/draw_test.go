package grid

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	ColorWhite uint32 = 0xffffffff
	ColorGreen uint32 = 0xff00ff00
	ColorRed   uint32 = 0xffff0000

	WindowDuration uint32 = 200 // Milliseconds.
)

func TestDrawing(t *testing.T) {
	g := NewGrid(9, 5, 50, 50)
	if err := drawGrid(g); err != nil {
		t.Errorf("error: failed to draw grid")
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
			d := g.Cell(x, y).Dimensions()
			outer := &sdl.Rect{
				X: int32(d.X),
				Y: int32(d.Y),
				W: int32(d.W),
				H: int32(d.H),
			}
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
			color := checkerColor(0xffff0000, 0xff0000ff, x, y)
			err = surface.FillRect(inner, color)
			if err != nil {
				return err
			}
		}
	}

	if err = window.UpdateSurface(); err != nil {
		return err
	}

	sdl.Delay(WindowDuration)

	return nil
}
