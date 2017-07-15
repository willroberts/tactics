package input

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
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
		} else if action == ActionDown {
			c.menu.CursorDown()
			return nil
		} else if action == ActionSubmit {
			return c.menu.Buttons()[c.menu.CursorPos()].Handler()
		} else if action == ActionQuit {
			return engine.ErrQuitting
		}
	}
	return nil
}

func NewMenuController(m menu.Menu) MenuController {
	return &menuController{menu: m}
}

//

type CameraController interface {
	ProcessEvents([]sdl.Event) error
}

type cameraController struct {
	camera engine.Camera
}

func (c *cameraController) ProcessEvents(events []sdl.Event) error {
	for _, e := range events {
		action := HandleInput(e)
		if action == ActionQuit {
			return engine.ErrQuitting
		}

		var cameraSpeed int32 = 10
		if action == ActionUp {
			x, y := c.camera.Position()
			c.camera.MoveTo(x, y-cameraSpeed)
		}
		if action == ActionDown {
			x, y := c.camera.Position()
			c.camera.MoveTo(x, y+cameraSpeed)
		}
		if action == ActionLeft {
			x, y := c.camera.Position()
			c.camera.MoveTo(x-cameraSpeed, y)
		}
		if action == ActionRight {
			x, y := c.camera.Position()
			c.camera.MoveTo(x+cameraSpeed, y)
		}
	}
	return nil
}

func NewCameraController(c engine.Camera) CameraController {
	return &cameraController{camera: c}
}

//

type CutsceneController interface {
	ProcessEvents([]sdl.Event) error
}

type cutsceneController struct{}

func (c *cutsceneController) ProcessEvents(events []sdl.Event) error {
	return nil
}

func NewCutsceneController() CutsceneController {
	return &cutsceneController{}
}

//

type GameController interface {
	ProcessEvents([]sdl.Event) error
}

type gameController struct{}

func (c *gameController) ProcessEvents(events []sdl.Event) error {
	return nil
}

func NewGameController() GameController {
	return &gameController{}
}
