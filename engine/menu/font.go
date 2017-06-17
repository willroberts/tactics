package menu

import "github.com/veandco/go-sdl2/ttf"

func InitializeFont(filename string, fontSize int) (*ttf.Font, error) {
	if err := ttf.Init(); err != nil {
		return &ttf.Font{}, err
	}

	font, err := ttf.OpenFont(filename, fontSize)
	if err != nil {
		return &ttf.Font{}, err
	}

	return font, nil
}
