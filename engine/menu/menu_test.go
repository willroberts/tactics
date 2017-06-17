package menu

import (
	"testing"

	"github.com/willroberts/tactics/engine"
)

const (
	testFont string = "testassets/font.ttf"
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
		FontFile: testFont,
	}
	m, err := NewMenu(p)
	if err != nil {
		t.Errorf("error: failed to create menu")
	}
	m.AddButton("Start Game")
	m.AddButton("Settings")
	m.AddButton("Quit")

	// Draw to test the result.
	for _, b := range m.Buttons() {
		err = eng.DrawLabel(b.Text, b.Rect, m.Font())
		if err != nil {
			t.Errorf("error: failed to draw label")
		}
	}
	err = eng.UpdateSurface()
	if err != nil {
		t.Errorf("error: failed to update surface")
	}
	eng.PauseRendering(400)
}
