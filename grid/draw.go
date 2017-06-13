package grid

import "github.com/veandco/go-sdl2/sdl"

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
			surface.FillRect(c.Rect(), c.Color())
		}
	}

	if err = window.UpdateSurface(); err != nil {
		return err
	}

	sdl.Delay(5000)

	return nil
}
