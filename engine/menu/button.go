package menu

import "github.com/veandco/go-sdl2/sdl"

// Button stores the state and logic for a menu button.
type Button struct {
	Text string
	Rect *sdl.Rect

	// Handler is executed when the user presses a submit key with this Button
	// selected.
	Handler func() error
}
