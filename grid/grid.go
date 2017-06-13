package grid

// Grid is a two-dimensional array containing a game board. It has no visual
// properties, and can be represented orthogonally or isometrically.
type Grid interface {
	Width() int
	Height() int

	CellWidth() int
	CellHeight() int

	Cell(x, y int) Cell
	SetCell(x, y int, c Cell)
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

func (g *grid) SetCell(x, y int, c Cell) {
	c.SetX(x)
	c.SetY(y)
	g.cells[x][y] = c
}

// NewGrid initializes and returns a Grid.
func NewGrid(x, y, cellWidth, cellHeight int) Grid {
	cells := make([][]Cell, x)
	for i := 0; i < x; i++ {
		row := make([]Cell, y)
		for j := range row {
			row[j] = NewCell(i, j, cellWidth, cellHeight)
		}
		cells[i] = row
	}

	return &grid{
		width:      x,
		height:     y,
		cells:      cells,
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
	}
}
