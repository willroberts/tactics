package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// SDLEngine is our interface to SDL2.
type SDLEngine interface {
	Window() *sdl.Window
	Surface() *sdl.Surface
	Renderer() *sdl.Renderer

	ProcessTextures(Spritesheet) error

	ClearScreen() error
	DrawRect(*sdl.Rect, uint32) error
	DrawLabel(string, *sdl.Rect, *ttf.Font) error
	DrawTexture(*sdl.Texture) error
	UpdateSurface() error

	Events() []sdl.Event

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

func (s *sdlengine) ProcessTextures(ss Spritesheet) error {
	for _, im := range ss.Sprites() {
		tex, err := ss.CreateTexture(im, s.Renderer())
		if err != nil {
			return err
		}
		ss.AddTexture(tex)
	}
	return nil
}

func (s *sdlengine) ClearScreen() error {
	return s.renderer.Clear()
}

func (s *sdlengine) DrawRect(rect *sdl.Rect, color uint32) error {
	return s.surface.FillRect(rect, color)
}

func (s *sdlengine) DrawLabel(text string, rect *sdl.Rect, font *ttf.Font) error {
	// FIXME: See if label has width and height? For automatic rect sizing.
	label, err := font.RenderUTF8_Solid(text, sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	if err != nil {
		return err
	}
	defer label.Free()

	if err = label.Blit(nil, s.Surface(), rect); err != nil {
		return err
	}

	return nil
}

func (s *sdlengine) DrawTexture(tex *sdl.Texture) error {
	src := &sdl.Rect{X: 0, Y: 0, W: 20, H: 20}
	dst := &sdl.Rect{X: 20, Y: 20, W: 20, H: 20}
	err := s.renderer.Copy(tex, src, dst)
	if err != nil {
		return err
	}
	return nil
}

func (s *sdlengine) UpdateSurface() error {
	return s.window.UpdateSurface()
}

// Events reads all pending events from the event loop, and returns them as a
// slice.
// FIXME: Determine the best way to do event processing, so we don't drop any
// events accidentally.
func (s *sdlengine) Events() []sdl.Event {
	events := []sdl.Event{}
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}
		events = append(events, e)
	}
	return events
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
