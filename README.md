# tactics

A 2D tactics game written in go.

## Documentations

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

* Render an isometric grid
* Render an isometric grid with sprite textures
* Read https://godoc.org/github.com/veandco/go-sdl2/sdl
* Allow resolution adjustment in Settings menu
* Clean up input package, add tests
* Add music and sound effects
* Add units
* Add assets from http://flarerpg.org/
* Add movement logic
* Add turn-based combat logic

Design and Organization

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
