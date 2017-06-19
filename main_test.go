// +build integration

// Ported from menu package. Add a menu to main.go.

package main

import (
	"errors"
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/engine/input"
	"github.com/willroberts/tactics/engine/menu"
)

const (
	testFont    string = "engine/menu/testassets/font.ttf"
	missingFont string = "engine/menu/testassets/missing.ttf"
	width       int    = 640
	height      int    = 480
)

var (
	m       menu.Menu
	QuitNow error = errors.New("quitting")
)

func TestMenu(t *testing.T) {
	p := menu.NewMenuParams{
		W:        int32(width),
		H:        int32(height),
		ButtonW:  int32(width / 2),
		ButtonH:  48,
		FontFile: missingFont,
	}
	var err error
	m, err = menu.NewMenu(p)
	if err == nil {
		t.Errorf("error: failed to detect invalid font")
	}
	p.FontFile = testFont
	m, err = menu.NewMenu(p)
	if err != nil {
		t.Errorf("error: failed to create menu")
	}

	startFunc := func() error {
		log.Println("starting the game")
		return nil
	}
	m.AddButton("Start Game", startFunc)

	quitFunc := func() error {
		log.Println("quitting")
		return QuitNow
	}
	m.AddButton("Quit", quitFunc)
}

func TestRendering(t *testing.T) {
	// Test rendering.
	eng, err := engine.NewSDLEngine("menu_test", 640, 480)
	if err != nil {
		t.Errorf("error: failed to get engine")
	}

	quit := false
	for !quit {
		_ = eng.ClearScreen()
		_ = eng.DrawRect(&sdl.Rect{
			X: 0,
			Y: 0,
			W: int32(width),
			H: int32(height),
		}, 0xff333333)
		for _, b := range m.Buttons() {
			err = eng.DrawLabel(b.Text, b.Rect, m.Font())
			if err != nil {
				t.Errorf("error: failed to draw label")
			}
		}

		// Process Input (integration test).
		for _, e := range eng.Events() {
			res := input.HandleInput(e)
			if res == input.ActionSubmit {
				// FIXME: Implement selection.
				err = m.Buttons()[m.CursorPos()].Handler()
				if err == QuitNow {
					quit = true
				}
				if err != nil && err != QuitNow {
					log.Println("error:", err)
				}
			} else if res == input.ActionQuit {
				quit = true
			} else if res == input.ActionUp {
				m.CursorUp()
			} else if res == input.ActionDown {
				m.CursorDown()
			} else if res == input.ActionRelease {
				// Ignore.
			} else {
				log.Println("error: unknown action:", res)
			}
		}

		cr, err := m.CursorRect()
		if err != nil {
			t.Errorf("error: failed to create cursor: %v", err)
		}
		err = eng.DrawRect(cr, 0xffff0000)
		if err != nil {
			t.Errorf("error: failed to draw cursor")
		}

		err = eng.UpdateSurface()
		if err != nil {
			t.Errorf("error: failed to update surface")
		}

		eng.PauseRendering(frameTime)
	}
}
