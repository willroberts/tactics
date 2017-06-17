package engine

import (
	"image"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	eng SDLEngine
	ss  Spritesheet
	err error
)

// TODO: Determine why this only has 78.6% test coverage.
func TestNewSDLEngine(t *testing.T) {
	eng, err = NewSDLEngine("test", 450, 250)
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

func TestClearScreen(t *testing.T) {
	err = eng.ClearScreen()
	if err != nil {
		t.Errorf("error: failed to clear the screen: %v", err)
	}
}

func TestDrawRect(t *testing.T) {
	rect := &sdl.Rect{X: 50, Y: 50, W: 50, H: 50}
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

func TestUpdate(t *testing.T) {
	err = eng.UpdateSurface()
	if err != nil {
		t.Errorf("error: failed to update surface: %v", err)
	}
}

func TestPauseRendering(t *testing.T) {
	var duration uint32 = 200 // Milliseconds.
	eng.PauseRendering(duration)
}

func TestDestroy(t *testing.T) {
	eng.DestroyWindow()
}
