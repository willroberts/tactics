package menu

import (
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
)

const (
	testFont string = "testassets/font.ttf"
)

func TestMenu(t *testing.T) {
	eng, err := engine.NewSDLEngine("menu_test", 640, 480)
	if err != nil {
		t.Errorf("error: failed to get engine")
	}

	w, h := eng.Window().GetSize()

	p := NewMenuParams{
		W:        int32(w),
		H:        int32(h),
		ButtonW:  320,
		ButtonH:  64,
		FontFile: testFont,
		FontSize: 48,
	}
	m, err := NewMenu(p)
	if err != nil {
		t.Errorf("error: failed to create menu")
	}
	m.AddButton("Start Game")
	m.AddButton("Settings")
	m.AddButton("Quit")

	err = eng.DrawRect(&sdl.Rect{W: 640, H: 1, X: 0, Y: 240}, 0xffff0000)
	if err != nil {
		t.Errorf("error: failed to draw separator line")
	}

	for _, b := range m.Buttons() {
		log.Println("button y value:", b.Rect.Y)
		err = eng.DrawLabel(b.Text, b.Rect, m.Font())
		if err != nil {
			t.Errorf("error: failed to draw label")
		}
	}

	err = eng.UpdateSurface()
	if err != nil {
		t.Errorf("error: failed to update surface")
	}
	eng.PauseRendering(2000)
}
