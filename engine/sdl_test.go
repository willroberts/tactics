package engine

import (
	"image"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	eng SDLEngine
	ss  Spritesheet
	err error
)

// TODO: Determine why this only has 78.6% test coverage.
func TestNewSDLEngine(t *testing.T) {
	eng, err = NewSDLEngine("test")
	if err != nil {
		t.Errorf("error: failed to create sdl engine: %v", err)
	}

	if eng.Window() == nil {
		t.Errorf("error: failed to create window")
	}
	if eng.Surface() == nil {
		t.Errorf("error: failed to create surface")
	}
	if eng.Renderer() == nil {
		t.Errorf("error: failed to create renderer")
	}
}

func TestProcessTextures(t *testing.T) {
	// Test bad sprites.
	ss = &spritesheet{}
	badImg := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{-1, -1},
	})
	ss.AddSprite(badImg)

	err = eng.ProcessTextures(ss)
	if err == nil {
		t.Errorf("error: did not detect processing failure")
	}

	// Test valid sprites.
	ss = &spritesheet{}
	for i := 0; i < 3; i++ {
		s := image.NewRGBA(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{1, 1},
		})
		ss.AddSprite(s)
	}

	err := eng.ProcessTextures(ss)
	if err != nil {
		t.Errorf("error: failed to process textures: %v", err)
	}
	if len(ss.Sprites()) != len(ss.Textures()) {
		t.Errorf("error: sprite/texture count mismatch")
	}
	for _, im := range ss.Textures() {
		if im == nil {
			t.Errorf("error: missing textures")
		}
	}
}

func TestFillWindow(t *testing.T) {
	err = eng.FillWindow(0xff000080)
	if err != nil {
		t.Errorf("error: failed to fill window with color")
	}
}

func TestDrawIsometricRect(t *testing.T) {
	rect := &sdl.Rect{X: 20, Y: 20, W: 20, H: 20}
	err = eng.DrawIsometricRect(rect, 0xff008000)
	if err != nil {
		t.Errorf("error: failed to draw isometric rect")
	}
}

func TestDrawLabel(t *testing.T) {
	// Initialize a font for testing. Most font features are located in the menu
	// package currently.
	_ = ttf.Init()
	font, _ := ttf.OpenFont("../assets/fonts/pixelated.ttf", 20)

	text := "testing"
	rect := &sdl.Rect{X: 20, Y: 50, W: 50, H: 20}
	err = eng.DrawLabel(text, rect, font)
	if err != nil {
		t.Errorf("error: failed to draw label")
	}
}

func TestColorToRGBA(t *testing.T) {
	rgba := colorToRGBA(0xff336699)
	if rgba.R != 51 {
		t.Errorf("error: bad red value %d for color to rgba", rgba.R)
	}
	if rgba.G != 102 {
		t.Errorf("error: bad green value %d for color to rgba", rgba.G)
	}
	if rgba.B != 153 {
		t.Errorf("error: bad blue value %d for color to rgba", rgba.B)
	}
	if rgba.A != 255 {
		t.Errorf("error: bad alpha value %d for color to rgba", rgba.A)
	}
}

func TestCartesianToIsoPoly(t *testing.T) {
	rect := &sdl.Rect{X: 32, Y: 32, W: 32, H: 32}
	vx, vy := cartesianToIsoPoly(rect)

	expectedX := []int16{0, -32, 0, 32}
	for i, x := range vx {
		if x != expectedX[i] {
			t.Errorf("error: unexpected value %d->%d from cartesianToIsoPoly", i, x)
		}
	}

	expectedY := []int16{32, 48, 64, 48}
	for i, y := range vy {
		if y != expectedY[i] {
			t.Errorf("error: unexpected value %d->%d from cartesianToIsoPoly", i, y)
		}
	}
}

func TestClearScreen(t *testing.T) {
	err = eng.ClearScreen()
	if err != nil {
		t.Errorf("error: failed to clear the screen: %v", err)
	}
}

func TestDrawRect(t *testing.T) {
	rect := &sdl.Rect{X: 50, Y: 20, W: 20, H: 20}
	var color uint32 = 0xff33ff33
	err = eng.DrawRect(rect, color)
	if err != nil {
		t.Errorf("error: failed to draw rect: %v", err)
	}
}

func TestDrawTexture(t *testing.T) {
	// Test valid texture.
	tex := ss.Textures()[0]
	err := eng.DrawTexture(tex)
	if err != nil {
		t.Errorf("error: failed to draw texture: %v", err)
	}

	// Test invalid texture.
	tex = &sdl.Texture{}
	err = eng.DrawTexture(tex)
	if err == nil {
		t.Errorf("error: failed to detect invalid texture")
	}
}

func TestUpdateSurface(t *testing.T) {
	err = eng.UpdateSurface()
	if err != nil {
		t.Errorf("error: failed to update surface: %v", err)
	}
}

func TestEvents(t *testing.T) {
	_ = eng.Events()
}

func TestPauseRendering(t *testing.T) {
	var duration uint32 = 200 // Milliseconds.
	eng.PauseRendering(duration)
}

func TestDestroy(t *testing.T) {
	eng.DestroyWindow()
}
