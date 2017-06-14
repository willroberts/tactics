package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/gfx"
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
)

func main() {
	// Build a grid.
	g := grid.NewGrid(9, 5, 50, 50)
	_ = g

	// TODO: Load textures from a TMX file.

	// Add units.

	// Render with SDL.
	// TODO: Try with OpenGL.
	eng, err := engine.NewSDLEngine("tactics", windowWidth, windowHeight)
	if err != nil {
		log.Fatalln("failed to initalize sdl engine:", err)
	}

	// TODO: Create game objects somewhere.
	var gameObjects []*sdl.Rect

	// Main loop.
	// TODO: Move everything above this into init()?
	for {
		eng.ClearScreen()
		//game.ProcessFrame() // TODO: Implement package `game`.
		for _, o := range gameObjects {
			eng.DrawRect(o, 0xff33aa33)
		}
		eng.UpdateSurface()
		eng.PauseRendering(frameTime)
	}
}
