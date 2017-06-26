package input

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine/menu"
)

type MenuController interface {
	ProcessEvents([]sdl.Event) error
}

type menuController struct {
	menu menu.Menu
}

func (c *menuController) ProcessEvents(events []sdl.Event) error {
	for _, e := range events {
		action := HandleInput(e)
		if action == ActionUp {
			c.menu.CursorUp()
			return nil
		}
		if action == ActionDown {
			c.menu.CursorDown()
			return nil
		}
		if action == ActionSubmit {
			return c.menu.Buttons()[c.menu.CursorPos()].Handler()
		}
		if action == ActionQuit {
			// FIXME.
			return nil
		}
	}
	return nil
}

func NewMenuController(m menu.Menu) MenuController {
	return &menuController{menu: m}
}

type CameraController interface {
	ProcessEvents([]sdl.Event) error
}

type cameraController struct {
}

func (c *cameraController) ProcessEvents(events []sdl.Event) error {
	return nil
}

func NewCameraController() CameraController {
	return &cameraController{}
}
