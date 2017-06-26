package grid

// Grid is a two-dimensional array containing a game board which can be
// represented orthogonally or isometrically.
type Grid interface {
	Width() int
	Height() int

	CellWidth() int
	CellHeight() int

	Cells() [][]Cell
	Cell(x, y int) Cell

	Checkerboard(color1, color2 uint32)
}

type grid struct {
	width  int
	height int

	cells      [][]Cell
	cellWidth  int
	cellHeight int
}

func (g *grid) Width() int {
	return g.width
}

func (g *grid) Height() int {
	return g.height
}

func (g *grid) CellWidth() int {
	return g.cellWidth
}

func (g *grid) CellHeight() int {
	return g.cellHeight
}

func (g *grid) Cells() [][]Cell {
	return g.cells
}

func (g *grid) Cell(x, y int) Cell {
	return g.cells[x][y]
}

// Checkerboard fills a grid's cells to create a checkerboard
// pattern.
func (g *grid) Checkerboard(color1, color2 uint32) {
	for x, col := range g.cells {
		for y := range col {
			g.cells[x][y].SetColor(checkerColor(color1, color2, x, y))
		}
	}
}

// checkerColor is a helper function for determining which color
// a cell should be, based on its coordinates.
func checkerColor(color1, color2 uint32, x, y int) uint32 {
	if x%2 == 0 {
		if y%2 == 0 {
			return color1
		}
		return color2
	}
	if y%2 == 0 {
		return color2
	}
	return color1
}

// NewGrid initializes and returns a Grid. `width` and `height` specify the
// dimensions of the grid itself. `cellWidth` and `cellHeight` specify the
// dimensions of cells within the grid.
func NewGrid(width, height, cellWidth, cellHeight int) Grid {
	cells := make([][]Cell, width)
	for i := 0; i < width; i++ {
		row := make([]Cell, height)
		for j := range row {
			row[j] = newCell(i, j, cellWidth, cellHeight)
		}
		cells[i] = row
	}

	return &grid{
		width:      width,
		height:     height,
		cells:      cells,
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
	}
}
