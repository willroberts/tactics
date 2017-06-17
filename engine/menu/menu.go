package menu

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Menu represents a basic main menu with buttons aligned vertically. Other
// layouts may be supported later.
type Menu interface {
	AddButton(string)
	Buttons() []*Button
	Font() *ttf.Font
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

func (m *menu) AddButton(text string) {
	// Create and add the new button.
	b := &Button{Text: text}
	m.buttons = append(m.buttons, b)

	// Update positions for all buttons.
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

func (m *menu) Buttons() []*Button {
	return m.buttons
}

func (m *menu) Font() *ttf.Font {
	return m.font
}

type NewMenuParams struct {
	W        int32
	H        int32
	ButtonW  int32
	ButtonH  int32
	FontFile string
}

func NewMenu(p NewMenuParams) (Menu, error) {
	f, err := InitializeFont(p.FontFile, int(p.ButtonH))
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
