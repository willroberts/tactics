package grid

import (
	"log"
	"testing"
)

const (
	GRID_WIDTH  int = 20
	GRID_HEIGHT int = 10
)

func TestGrid(t *testing.T) {
	g := NewGrid(GRID_WIDTH, GRID_HEIGHT)
	log.Println("w:", g.Width())
	log.Println("h:", g.Height())
	log.Println("5,7:", g.Cell(5, 7))
}
