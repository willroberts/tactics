# tactics

A 2D tactics game written in go.

## Packages

* engine: sdl2-based game engine
* engine/audio: music and sound effects
* engine/input: keyboard and mouse input
* engine/menu: menu logic and rendering
* grid: standalone code for representing game boards
* tmx: standalone code for interacting with tiled files
* game: logic code specific to the game we're making
* game/unit: code for representing units (players, NPCs, objects)

## To Do

* 100% test coverage (engine, menu)
* Keep packages as separate as possible and remove any tight coupling
  * engine/spritesheet uses tmx (this might be okay)
		* consider killing tmx package since it wraps go-tmx so closely
	* grid uses sdl rects (sdl should be in engine only!)
		* does it use them or just give an easy way to create them?
		* consider returning the int coords to the engine and conv to rect there
* use int32 in engine and int elsewhere
* Input Handling (input)
* Audio (audio)
* Game Logic (game)
* Add OpenGL Context for 3D rendering (engine)
