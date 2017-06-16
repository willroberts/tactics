package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

// SDLEngine is our interface to SDL2.
type SDLEngine interface {
	Window() *sdl.Window
	Surface() *sdl.Surface
	Renderer() *sdl.Renderer

	ProcessTextures(Spritesheet) error

	ClearScreen() error
	DrawRect(*sdl.Rect, uint32) error
	DrawTexture(*sdl.Texture) error
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

func (s *sdlengine) DrawTexture(tex *sdl.Texture) error {
	src := &sdl.Rect{X: 0, Y: 0, W: 20, H: 20}
	dst := &sdl.Rect{X: 20, Y: 20, W: 20, H: 20}
	err := s.renderer.Copy(tex, src, dst)
	if err != nil {
		return err
	}
	/* HOW TO DRAW A TEX:
	   src = sdl.Rect{0, 0, 512, 512}
	   dst = sdl.Rect{100, 50, 512, 512}
	   renderer.Clear()
	   renderer.SetDrawColor(255, 0, 0, 255)
	   renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
	   renderer.Copy(texture, &src, &dst)
	   renderer.Present()
	   sdl.Delay(2000)
	*/
	return nil
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
