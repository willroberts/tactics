package input

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	ActionSubmit int = iota
	ActionQuit
	ActionUp
	ActionDown
	ActionLeft
	ActionRight
	ActionRelease
	ActionUnknown
)

var (
	submitMap map[sdl.Keycode]bool = map[sdl.Keycode]bool{
		keyEnter:    true,
		keySpacebar: true,
	}
	quitMap map[sdl.Keycode]bool = map[sdl.Keycode]bool{
		keyEscape: true,
		keyQ:      true,
	}
)

func HandleInput(e sdl.Event) int {
	switch t := e.(type) {
	case *sdl.QuitEvent:
		return ActionQuit
	case *sdl.KeyDownEvent:
		keyPressed := t.Keysym.Sym
		log.Println("key pressed:", keyPressed)
		if submitMap[keyPressed] {
			return ActionSubmit
		}
		if quitMap[keyPressed] {
			return ActionQuit
		}
		if keyPressed == keyArrowUp {
			return ActionUp
		}
		if keyPressed == keyArrowDown {
			return ActionDown
		}
	case *sdl.KeyUpEvent:
		// FIXME: Implement.
		keyReleased := t.Keysym.Sym
		_ = keyReleased
		return ActionRelease
	default:
		return ActionUnknown
	}
	return ActionUnknown
}
