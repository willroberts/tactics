package grid

// Grid is a two-dimensional array containing a game board which can be
// represented orthogonally or isometrically.
type Grid interface {
	Width() int
	Height() int

	CellWidth() int
	CellHeight() int

	Cell(x, y int) Cell
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

func (g *grid) Cell(x, y int) Cell {
	return g.cells[x][y]
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
