package menu

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Menu represents a basic main menu with buttons aligned vertically. Other
// layouts may be supported later.
type Menu interface {
	Font() *ttf.Font
	Buttons() []*Button
	CursorPos() int
	ResetCursor()

	CursorUp()
	CursorDown()
	CursorRect() (*sdl.Rect, error)

	AddButton(string, func() error)
	ClearButtons()
}

type menu struct {
	width        int32
	height       int32
	buttonWidth  int32
	buttonHeight int32
	font         *ttf.Font
	buttons      []*Button
	cursorPos    int
}

func (m *menu) Font() *ttf.Font {
	return m.font
}

func (m *menu) Buttons() []*Button {
	return m.buttons
}

func (m *menu) CursorPos() int {
	return m.cursorPos
}

func (m *menu) ResetCursor() {
	m.cursorPos = 0
}

func (m *menu) CursorUp() {
	if len(m.buttons) == 0 {
		return
	}
	if m.cursorPos == 0 {
		// can't move up, wrap around
		m.cursorPos = len(m.buttons) - 1
		return
	}
	m.cursorPos--
}

func (m *menu) CursorDown() {
	if len(m.buttons) == 0 {
		return
	}
	if m.cursorPos == len(m.buttons)-1 {
		// can't move down, wrap around
		m.cursorPos = 0
		return
	}
	m.cursorPos++
}

func (m *menu) CursorRect() (*sdl.Rect, error) {
	if len(m.buttons) == 0 {
		return &sdl.Rect{}, errors.New("error: can't create cursor without buttons")
	}
	adj := m.buttons[m.cursorPos].Rect
	size := adj.H
	margin := size / 4
	rect := &sdl.Rect{
		X: adj.X - size + margin,
		Y: adj.Y + margin,
		W: size / 2,
		H: size / 2,
	}
	return rect, nil
}

func (m *menu) AddButton(text string, handler func() error) {
	b := &Button{Text: text, Handler: handler}
	m.buttons = append(m.buttons, b)

	// Recalculate positions for all buttons automatically.
	l := len(m.buttons)
	for i := 0; i < l; i++ {
		startingPos := (m.height / 2) - (int32(l) * m.buttonHeight / 2)
		m.buttons[i].Rect = &sdl.Rect{
			W: m.buttonWidth,
			H: m.buttonHeight,
			X: m.width/2 - m.buttonWidth/2,
			Y: startingPos + (int32(i) * m.buttonHeight),
		}
	}
}

func (m *menu) ClearButtons() {
	m.buttons = []*Button{}
}

type NewMenuParams struct {
	W        int32
	H        int32
	ButtonW  int32
	ButtonH  int32
	FontFile string
}

func NewMenu(p NewMenuParams) (Menu, error) {
	err := ttf.Init()
	if err != nil {
		return &menu{}, err
	}

	f, err := ttf.OpenFont(p.FontFile, int(p.ButtonH))
	if err != nil {
		return &menu{}, err
	}
	m := &menu{
		width:        p.W,
		height:       p.H,
		buttonWidth:  p.ButtonW,
		buttonHeight: p.ButtonH,
		font:         f,
		cursorPos:    0,
	}
	return m, nil
}
