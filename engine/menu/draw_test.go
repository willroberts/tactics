package menu

import (
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/engine/input"
)

const (
	testFont    string = "testassets/font.ttf"
	missingFont string = "testassets/missing.ttf"
	frameTime   uint32 = 1000 / 30
	width       int    = 640
	height      int    = 480
)

var (
	m Menu
)

func TestMenu(t *testing.T) {
	p := NewMenuParams{
		W:        int32(width),
		H:        int32(height),
		ButtonW:  int32(width / 2),
		ButtonH:  48,
		FontFile: missingFont,
	}
	var err error
	m, err = NewMenu(p)
	if err == nil {
		t.Errorf("error: failed to detect invalid font")
	}
	p.FontFile = testFont
	m, err = NewMenu(p)
	if err != nil {
		t.Errorf("error: failed to create menu")
	}
	m.AddButton("Start Game")
	m.AddButton("Settings")
	m.AddButton("Quit")
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
			} else if res == input.ActionQuit {
				quit = true
			} else if res == input.ActionUp {
				m.CursorUp()
			} else if res == input.ActionDown {
				m.CursorDown()
			} else if res == input.ActionRelease {
				// Ignore.
			} else {
				log.Println("unknown action:", res)
			}
		}

		err = eng.DrawRect(m.CursorRect(), 0xffff0000)
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
