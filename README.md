# tactics

A 2D tactics game written in go.

## Design

Currently floating somewhere between Tactical RPGs (Tactics Ogre) and Tactical
CCGs (Duelyst).

## To Do

* 100% test coverage for `engine` package
* Replace log.Println+t.FailNow with t.Errorf
* Input Handling (package `input`)
* Fonts (package `text` or `label`)
* Menus (package `menu`)
* Audio (package `audio`)
* Game Logic (package `logic` or `game`)
  * Populate the `unit` package as needed
* Add OpenGL Context for 3D rendering to the `engine` package
