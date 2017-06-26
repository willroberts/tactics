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

* Make the ESC key quit the game in isometric scene
* Make the Q key also quit the game
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
