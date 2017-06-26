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
	ActionNotImplemented
	ActionUnknown
)

var (
	submitMap map[sdl.Keycode]bool = map[sdl.Keycode]bool{
		keyEnter:    true,
		sdl.K_SPACE: true,
	}
	quitMap map[sdl.Keycode]bool = map[sdl.Keycode]bool{
		keyEscape: true,
		sdl.K_q:   true,
	}
)

func HandleInput(e sdl.Event) int {
	switch t := e.(type) {
	case *sdl.QuitEvent:
		return ActionQuit
	case *sdl.TextInputEvent:
		key := textToKey(t)
		return handleKey(key)
	case *sdl.KeyDownEvent:
		key := t.Keysym.Sym
		return handleKey(key)
	case *sdl.KeyUpEvent:
		return ActionNotImplemented
	case *sdl.MouseMotionEvent:
		return ActionNotImplemented
	case *sdl.WindowEvent:
		return ActionNotImplemented
	default:
		log.Println("unknown action was:", t)
		return ActionUnknown
	}
	// This should never be reached, but go complains about missing return at end
	// of function.
	return ActionUnknown
}

// The first byte of t.Text contains the ASCII ID of the key pressed.
// For example, 'a' is 97, 'b' is 98, 'A' is 65, 'B" is 66, etc.
func textToKey(t *sdl.TextInputEvent) sdl.Keycode {
	keycode := t.Text[0]
	switch keycode {
	case 32:
		return sdl.K_SPACE
	case 113:
		return sdl.K_q
	default:
		log.Println("unknown text input:", keycode)
		return sdl.K_UNKNOWN
	}
}

func handleKey(key sdl.Keycode) int {
	if key == 0 {
		return ActionNotImplemented
	}

	log.Println("key pressed:", key)
	if submitMap[key] {
		return ActionSubmit
	} else if quitMap[key] {
		return ActionQuit
	} else if key == keyArrowUp {
		return ActionUp
	} else if key == keyArrowDown {
		return ActionDown
	} else if key == keyArrowLeft {
		return ActionLeft
	} else if key == keyArrowRight {
		return ActionRight
	}

	return ActionNotImplemented
}
