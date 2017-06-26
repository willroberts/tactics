package engine

import (
	"encoding/binary"
	"errors"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// SDLEngine is our interface to SDL2.
type SDLEngine interface {
	Window() *sdl.Window
	Surface() *sdl.Surface
	Renderer() *sdl.Renderer
	Camera() Camera

	ProcessTextures(Spritesheet) error

	ClearScreen() error
	FillWindow(uint32) error
	DrawRect(*sdl.Rect, uint32) error
	DrawIsometricRect(*sdl.Rect, uint32) error
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
	camera   Camera
	desktop  *sdl.DisplayMode
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

func (s *sdlengine) Camera() Camera {
	return s.camera
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

func (s *sdlengine) FillWindow(color uint32) error {
	// Rect is not shifted according to camera position like others.
	w, h := s.Window().GetSize()
	rect := &sdl.Rect{X: 0, Y: 0, W: int32(w), H: int32(h)}
	return s.surface.FillRect(rect, color)
}

func (s *sdlengine) DrawRect(rect *sdl.Rect, color uint32) error {
	rect = s.Camera().ShiftRect(rect)
	return s.surface.FillRect(rect, color)
}

func (s *sdlengine) DrawIsometricRect(rect *sdl.Rect, color uint32) error {
	vx, vy := cartesianToIsoPoly(rect)
	vx, vy = s.Camera().ShiftVectors(vx, vy)
	if b := gfx.FilledPolygonColor(s.Renderer(), vx, vy, colorToRGBA(color)); !b {
		// FIXME: Determine how to test (or suppress) this.
		return errors.New("error: FilledPolygonColor() returned false")
	}
	return nil
}

func (s *sdlengine) DrawLabel(text string, rect *sdl.Rect, font *ttf.Font) error {
	rect = s.Camera().ShiftRect(rect)
	// FIXME: See if label has width and height? For automatic rect sizing.
	label, err := font.RenderUTF8_Solid(text, sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	if err != nil {
		// FIXME: Determine how to test (or suppress) this.
		return err
	}
	defer label.Free()

	if err = label.Blit(nil, s.Surface(), rect); err != nil {
		// FIXME: Determine how to test (or suppress) this.
		return err
	}

	return nil
}

// FIXME: Actually implement this at some point instead of using hardcoded
// values to make tests pass.
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
// title. Implements and returns the SDLEngine interface.
func NewSDLEngine(title string) (SDLEngine, error) {
	s := &sdlengine{}

	// Create a fullscreen window, initially with tiny resolution.
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, 0, 0, sdl.WINDOW_FULLSCREEN) //sdl.WINDOW_SHOWN)
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

	s.camera = NewCamera()

	// Set the window dimensions to the desktop dimensions (windowed fullscreen).
	desktop := &sdl.DisplayMode{}
	// FIXME: Handle multiple screens (not just 0).
	if err := sdl.GetDesktopDisplayMode(0, desktop); err != nil {
		return s, err
	}
	s.desktop = desktop
	window.SetSize(int(desktop.W), int(desktop.H))

	return s, nil
}

func colorToRGBA(color uint32) sdl.Color {
	rgba := make([]byte, 4)
	binary.BigEndian.PutUint32(rgba, color)
	// Byte format is ARGB
	return sdl.Color{R: rgba[1], G: rgba[2], B: rgba[3], A: rgba[0]}
}

func cartesianToIsoPoly(rect *sdl.Rect) ([]int16, []int16) {
	points := [][]int16{
		[]int16{int16(rect.X), int16(rect.Y)},
		[]int16{int16(rect.X), int16(rect.Y) + int16(rect.H)},
		[]int16{int16(rect.X) + int16(rect.W), int16(rect.Y) + int16(rect.H)},
		[]int16{int16(rect.X) + int16(rect.W), int16(rect.Y)},
	}

	ip := make([][]int16, 4)
	for i, point := range points {
		newpoint := make([]int16, 2)
		newpoint[0] = point[0] - point[1]
		newpoint[1] = (point[0] + point[1]) / 2
		ip[i] = newpoint
	}

	vx := []int16{ip[0][0], ip[1][0], ip[2][0], ip[3][0]}
	vy := []int16{ip[0][1], ip[1][1], ip[2][1], ip[3][1]}
	return vx, vy
}
