# tactics

A 2D tactics game written in Go, using a custom engine on the SDL2 library.

Unfinished; current state includes rendering of the isometric grid but no actors
or game logic/scenes.

## Requirements

* Mingw-w64 and SDL2: https://github.com/veandco/go-sdl2#requirements

## Documentation

See godoc.org/github.com/willroberts/tactics
* engine: sdl2-based game engine
* engine/menu: menu logic and rendering
* engine/input: keyboard and mouse handling
* game/scenes: scenes for the game
* game/unit: code for representing units (players, NPCs, objects)
* grid: standalone code for representing game boards
* tmx: standalone code for interacting with tiled files

## To Do

Core Features

* Allow camera movement in both X and Y directions simultaneously.
* Clean up input package, add tests
* Allow resolution adjustment in Settings menu
* Add units
* Add sprite textures
* Add assets from http://flarerpg.org
* Add movement logic
* Add turn-based combat logic
* Add music and sound effects

Design and Organization

* Read https://godoc.org/github.com/veandco/go-sdl2/sdl
* Remove all SDL-specific code from grid package (?)
	* OR: move grid to engine/grid
	* Textures stored in cells
	* Colors stored in cells
	* Checkerboard & CheckerColor funcs for color operations
* Move tmx to engine/tmx
* Evaluate usage of int32 (engine)

Packaging and Polish

* 100% test coverage
* Comment audit
