package main

import (
	"log"

	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/grid"
)

const (
	gridWidth    int = 9
	gridHeight   int = 5
	cellWidth    int = 50
	cellHeight   int = 50
	windowWidth  int = gridWidth * cellWidth
	windowHeight int = gridHeight * cellHeight

	framesPerSecond uint32 = 60
	frameTime       uint32 = 1000 / framesPerSecond // Milliseconds.

	colorLightBlue uint32 = 0xff6495ed
	colorDarkBlue  uint32 = 0xff4682b4
	//colorDarkBlue uint32 = 0xff4169e1
)

func main() {
	// FIXME: Move some things to init(), leaving only the loop in main().
	g := grid.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight)
	g.Checkerboard(colorLightBlue, colorDarkBlue)

	// TODO: Load textures from a TMX file.

	// TODO: Add units.

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
				if err = eng.DrawRect(cell.Rect(), cell.Color()); err != nil {
					log.Fatalln("error drawing cell:", err)
				}
			}
		}

		if err = eng.UpdateSurface(); err != nil {
			log.Fatalln("error updating surface:", err)
		}

		eng.PauseRendering(frameTime)
	}
}
