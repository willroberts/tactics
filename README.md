# tactics

A 2D tactics game written in go.

## Packages

* engine: sdl2-based game engine
* engine/audio: music and sound effects
* engine/input: keyboard and mouse input
* engine/menu: menu logic and rendering
* engine/sprite: png decoding into image.Image
* game: logic code specific to the game we're making
* game/unit: code for representing units (players, NPCs, objects)
* grid: standalone code for representing game boards
* tmx: standalone code for interacting with tiled files

## To Do

### Priority 1 (Core Features)

* Sprite (Texture) Rendering
* Input Handling (input)
* Audio (audio)

### Priority 2 (Design and Consistency)

* Remove all SDL-specific code from grid package
	* Textures stored in cells
	* Colors stored in cells
	* Checkerboard & CheckerColor funcs for color operations
	* OR: move grid to engine/grid (?)
* Move tmx to engine/tmx (?)
* Use int32 in engine and int elsewhere (?)
* Add OpenGL Context for 3D rendering (engine)
