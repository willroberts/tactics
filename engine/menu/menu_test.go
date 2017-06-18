package menu

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
)

const (
	testFont    string = "testassets/font.ttf"
	missingFont string = "testassets/missing.ttf"
)

func TestMenu(t *testing.T) {
	// Create an engine.
	eng, err := engine.NewSDLEngine("menu_test", 640, 480)
	if err != nil {
		t.Errorf("error: failed to get engine")
	}
	w, h := eng.Window().GetSize()

	// Create a menu.
	p := NewMenuParams{
		W:        int32(w),
		H:        int32(h),
		ButtonW:  int32(w / 2),
		ButtonH:  48,
		FontFile: missingFont,
	}
	m, err := NewMenu(p)
	if err == nil {
		t.Errorf("error: failed to detect invalid font")
	}
	p.FontFile = testFont
	m, err = NewMenu(p)
	if err != nil {
		t.Errorf("error: failed to create menu")
	}

	// Create buttons.
	m.AddButton("Start Game")
	m.AddButton("Settings")
	m.AddButton("Quit")

	// Draw to test the result.
	err = eng.DrawRect(&sdl.Rect{X: 0, Y: 0, W: p.W, H: p.H}, 0xff333333)
	if err != nil {
		t.Errorf("error: failed to draw grey background")
	}
	for _, b := range m.Buttons() {
		err = eng.DrawLabel(b.Text, b.Rect, m.Font())
		if err != nil {
			t.Errorf("error: failed to draw label")
		}
	}
	cursorRect := m.CursorPos()
	err = eng.DrawRect(cursorRect, 0xffff0000)
	if err != nil {
		t.Errorf("error: failed to draw cursor")
	}

	//
	err = eng.UpdateSurface()
	if err != nil {
		t.Errorf("error: failed to update surface")
	}
	eng.PauseRendering(200)
}
