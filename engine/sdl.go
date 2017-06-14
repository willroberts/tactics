package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

// SDLEngine is our interface to SDL2.
type SDLEngine interface {
	Window() *sdl.Window
	Surface() *sdl.Surface
	Renderer() *sdl.Renderer
	ClearScreen() error
	DrawRect(*sdl.Rect, uint32) error
	UpdateSurface() error
	PauseRendering(uint32)
	DestroyWindow()
}

type sdlengine struct {
	window   *sdl.Window
	surface  *sdl.Surface
	renderer *sdl.Renderer
}

func (s *sdlengine) Window() *sdl.Window {
	return s.window
}

func (s *sdlengine) Surface() *sdl.Surface {
	return s.surface
}

func (s *sdlengine) Renderer() *sdl.Renderer {
	return s.renderer
}

func (s *sdlengine) ClearScreen() error {
	return s.renderer.Clear()
}

func (s *sdlengine) DrawRect(rect *sdl.Rect, color uint32) error {
	return s.surface.FillRect(rect, color)
}

func (s *sdlengine) UpdateSurface() error {
	return s.window.UpdateSurface()
}

func (s *sdlengine) PauseRendering(t uint32) {
	sdl.Delay(t)
}

func (s *sdlengine) DestroyWindow() {
	s.window.Destroy()
}

// NewSDLEngine creates an SDL window, surface, and renderer with the given
// properties. Implements and returns the SDLEngine interface.
func NewSDLEngine(title string, width int, height int) (SDLEngine, error) {
	s := &sdlengine{}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return s, err
	}
	s.window = window

	surface, err := window.GetSurface()
	if err != nil {
		return s, err
	}
	s.surface = surface

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		return s, err
	}
	s.renderer = renderer

	return s, nil
}
