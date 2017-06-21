# tactics

A 2D tactics game written in go.

## Packages

* engine: sdl2-based game engine
* engine/menu: menu logic and rendering
* engine/input: keyboard and mouse handling
* game/scenes: scenes for the game
* game/unit: code for representing units (players, NPCs, objects)
* grid: standalone code for representing game boards
* tmx: standalone code for interacting with tiled files

## To Do

### Priority 1 (Core Features)

* Menu System
	* Settings Menu
	* Resolution Setting (Resize SDLEngine.Window)
	* sdl.WindowEvent?
* Input Handling
	* Don't use "starting game" as an error
	* Clean up interface once required components are in
	* Tests
* Sprite (Texture) Rendering
* Audio

### Priority 2 (Design and Consistency)

* Remove all SDL-specific code from grid package (?)
	* OR: move grid to engine/grid
	* Textures stored in cells
	* Colors stored in cells
	* Checkerboard & CheckerColor funcs for color operations
* Move tmx to engine/tmx
* Evaluate usage of int32 (engine)
* Add OpenGL Context for 3D rendering (engine)

### Priority 3 (Release)

* 100% test coverage
* Comment audit
