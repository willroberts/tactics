package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/grid"
	"github.com/willroberts/tactics/unit"
)

const (
	gridWidth  int = 9
	gridHeight int = 5

	// Using multiples of four so we can cut a cell into quadrants easily.
	cellWidth    int = 100
	cellHeight   int = 100
	windowWidth  int = gridWidth * cellWidth
	windowHeight int = gridHeight * cellHeight

	framesPerSecond uint32 = 30
	frameTime       uint32 = 1000 / framesPerSecond // Milliseconds.

	colorLightBlue uint32 = 0xff6495ed
	colorDarkBlue  uint32 = 0xff4682b4
	colorWhite     uint32 = 0xffffffff
)

func main() {
	// FIXME: Move some things to init(), leaving only the loop in main().
	g := grid.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight)

	// TODO: Load textures from a TMX file.
	g.Checkerboard(colorLightBlue, colorDarkBlue)

	// Adds units.
	// Positions chosen for testing. Replace with something legitimate.
	g.Cell(0, 2).SetContents(unit.NewUnit())
	g.Cell(8, 2).SetContents(unit.NewUnit())
	g.Cell(4, 0).SetContents(unit.NewUnit())
	g.Cell(4, 4).SetContents(unit.NewUnit())
	g.Cell(5, 2).SetContents(unit.NewUnit())

	eng, err := engine.NewSDLEngine("tactics", windowWidth, windowHeight)
	if err != nil {
		log.Fatalln("failed to initalize sdl engine:", err)
	}

	for {
		if err = eng.ClearScreen(); err != nil {
			log.Fatalln("error clearing screen:", err)
		}

		// FIXME: Use a channel for iteration.
		for _, col := range g.Cells() {
			for _, cell := range col {
				// Draw cells.
				if err = eng.DrawRect(cell.Rect(), cell.Color()); err != nil {
					log.Fatalln("error drawing cell:", err)
				}

				// Draw units and objects.
				if cell.IsOccupied() {
					unitRect := &sdl.Rect{
						W: int32(cellWidth / 2),
						H: int32(cellHeight / 2),
						X: int32((cell.X() * cellWidth) + cellWidth/4),
						Y: int32((cell.Y() * cellHeight) + cellWidth/4),
					}
					if err = eng.DrawRect(unitRect, colorWhite); err != nil {
						log.Fatalln("error drawing unit:", err)
					}
				}
			}
		}

		if err = eng.UpdateSurface(); err != nil {
			log.Fatalln("error updating surface:", err)
		}

		eng.PauseRendering(frameTime)
	}
}
