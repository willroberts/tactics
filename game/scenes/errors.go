package scenes

import "errors"

var (
	ErrEndScene error = errors.New("ending scene")
	ErrQuitGame error = errors.New("quitting the game")
)
