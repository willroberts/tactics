package scenes

import (
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

type mainMenuScene struct {
	eng        engine.SDLEngine
	m          menu.Menu
	controller input.MenuController
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
		return ErrEndScene
	})

	m.AddButton("Quit", func() error {
		return engine.ErrQuitting
	})

	s.controller = input.NewMenuController(m)

	return nil
}

func (s *mainMenuScene) Main() error {
	if err := s.eng.ClearScreen(); err != nil {
		return err
	}

	if err := s.eng.FillWindow(cGray); err != nil {
		return err
	}

	for _, b := range s.m.Buttons() {
		if err := s.eng.DrawLabel(b.Text, b.Rect, s.m.Font()); err != nil {
			return err
		}
	}

	if err := s.controller.ProcessEvents(s.eng.Events()); err != nil {
		return err
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
