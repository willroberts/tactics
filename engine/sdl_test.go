package engine

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	engine SDLEngine
	err    error
)

// TODO: Determine why this only has 78.6% test coverage.
func TestNewSDLEngine(t *testing.T) {
	engine, err = NewSDLEngine("test", 450, 250)
	if err != nil {
		t.FailNow()
	}

	if engine.Window() == nil {
		t.FailNow()
	}
	if engine.Surface() == nil {
		t.FailNow()
	}
	if engine.Renderer() == nil {
		t.FailNow()
	}
}

func TestClearScreen(t *testing.T) {
	err = engine.ClearScreen()
	if err != nil {
		t.FailNow()
	}
}

func TestDrawRect(t *testing.T) {
	rect := &sdl.Rect{X: 50, Y: 50, W: 50, H: 50}
	var color uint32 = 0xff33ff33
	err = engine.DrawRect(rect, color)
	if err != nil {
		t.FailNow()
	}
}

func TestUpdate(t *testing.T) {
	err = engine.UpdateSurface()
	if err != nil {
		t.FailNow()
	}
}

func TestPauseRendering(t *testing.T) {
	var duration uint32 = 200 // Milliseconds.
	engine.PauseRendering(duration)
}

func TestDestroy(t *testing.T) {
	engine.DestroyWindow()
}
