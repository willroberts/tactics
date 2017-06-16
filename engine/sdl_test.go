package engine

import (
	"image"
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	eng SDLEngine
	err error
)

// TODO: Determine why this only has 78.6% test coverage.
func TestNewSDLEngine(t *testing.T) {
	eng, err = NewSDLEngine("test", 450, 250)
	if err != nil {
		t.FailNow()
	}

	if eng.Window() == nil {
		t.FailNow()
	}
	if eng.Surface() == nil {
		t.FailNow()
	}
	if eng.Renderer() == nil {
		t.FailNow()
	}
}

func TestProcessTextures(t *testing.T) {
	ss := &spritesheet{}
	for i := 0; i < 3; i++ {
		s := image.NewRGBA(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{1, 1},
		})
		ss.AddSprite(s)
	}

	err := eng.ProcessTextures(ss)
	if err != nil {
		log.Println("error: failed to process textures:", err)
		t.FailNow()
	}
	if len(ss.Sprites()) != len(ss.Textures()) {
		log.Println("error: sprite-texture count mismatch")
		t.FailNow()
	}
	for _, im := range ss.Textures() {
		if im == nil {
			log.Println("error: missing textures")
			t.FailNow()
		}
	}

	ss = &spritesheet{}
	badImg := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{-1, -1},
	})
	ss.AddSprite(badImg)

	err = eng.ProcessTextures(ss)
	if err == nil {
		log.Println("no error processing image")
		t.FailNow()
	}
}

func TestClearScreen(t *testing.T) {
	err = eng.ClearScreen()
	if err != nil {
		t.FailNow()
	}
}

func TestDrawRect(t *testing.T) {
	rect := &sdl.Rect{X: 50, Y: 50, W: 50, H: 50}
	var color uint32 = 0xff33ff33
	err = eng.DrawRect(rect, color)
	if err != nil {
		t.FailNow()
	}
}

// FIXME: Update test after implementing function.
func TestDrawTexture(t *testing.T) {
	tex := &sdl.Texture{}
	err := eng.DrawTexture(tex)
	if err == nil {
		t.FailNow()
	}
}

func TestUpdate(t *testing.T) {
	err = eng.UpdateSurface()
	if err != nil {
		t.FailNow()
	}
}

func TestPauseRendering(t *testing.T) {
	var duration uint32 = 200 // Milliseconds.
	eng.PauseRendering(duration)
}

func TestDestroy(t *testing.T) {
	eng.DestroyWindow()
}
