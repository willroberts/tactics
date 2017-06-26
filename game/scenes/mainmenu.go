package scenes

import (
	"errors"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/engine/input"
	"github.com/willroberts/tactics/engine/menu"
)

const (
	fontFile string = "assets/fonts/pixelated.ttf"
	fTime    uint32 = 1000 / 30
	cGray    uint32 = 0xff333333
	cRed     uint32 = 0xffff0000
)

var (
	ErrStartingGame error = errors.New("starting game")
	ErrQuitting     error = errors.New("quitting")
)

type mainMenuScene struct {
	eng engine.SDLEngine
	m   menu.Menu
}

func (s *mainMenuScene) Setup() error {
	w, h := s.eng.Window().GetSize()
	p := menu.NewMenuParams{
		W:        int32(w),
		H:        int32(h),
		ButtonW:  int32(w / 2),
		ButtonH:  48,
		FontFile: fontFile,
	}

	m, err := menu.NewMenu(p)
	if err != nil {
		return err
	}
	s.m = m

	m.AddButton("Start Game", func() error {
		return ErrStartingGame
	})

	m.AddButton("Quit", func() error {
		return ErrQuitting
	})

	return nil
}

func (s *mainMenuScene) Main() error {
	w, h := s.eng.Window().GetSize()

	err := s.eng.ClearScreen()
	if err != nil {
		return err
	}

	err = s.eng.DrawRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(w),
		H: int32(h),
	}, cGray)
	if err != nil {
		return err
	}

	for _, b := range s.m.Buttons() {
		err = s.eng.DrawLabel(b.Text, b.Rect, s.m.Font())
		if err != nil {
			return err
		}
	}

	for _, e := range s.eng.Events() {
		res := input.HandleInput(e)
		if res == input.ActionSubmit {
			return s.m.Buttons()[s.m.CursorPos()].Handler()
		} else if res == input.ActionQuit {
			return ErrQuitting
		} else if res == input.ActionUp {
			s.m.CursorUp()
		} else if res == input.ActionDown {
			s.m.CursorDown()
		} else if res == input.ActionNotImplemented {
			// Ignore.
		} else {
			log.Println("unhandled action:", res)
		}
	}

	cr, err := s.m.CursorRect()
	if err != nil {
		return err
	}

	if err = s.eng.DrawRect(cr, cRed); err != nil {
		return err
	}

	if err = s.eng.UpdateSurface(); err != nil {
		return err
	}

	s.eng.PauseRendering(fTime)
	return nil
}

func (s *mainMenuScene) Teardown() error {
	return nil
}

func NewMainMenuScene(e engine.SDLEngine) engine.Scene {
	return &mainMenuScene{eng: e}
}
