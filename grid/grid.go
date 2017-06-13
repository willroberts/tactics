package grid

type Grid interface {
	Width() int
	Height() int

	Cell(x, y int) Cell
	SetCell(x, y int, c Cell)
}

type grid struct {
	width  int
	height int
	cells  [][]Cell
}

func (g *grid) Width() int {
	return g.width
}

func (g *grid) Height() int {
	return g.height
}

func (g *grid) Cell(x, y int) Cell {
	return g.cells[x][y]
}

func (g *grid) SetCell(x, y int, c Cell) {
	g.cells[x][y] = c
}

func NewGrid(x, y int) Grid {
	cells := make([][]Cell, x)
	for i := 0; i < x; i++ {
		row := make([]Cell, y)
		for j := range row {
			row[j] = NewCell(i, j)
		}
		cells[i] = row
	}

	return &grid{
		width:  x,
		height: y,
		cells:  cells,
	}
}
