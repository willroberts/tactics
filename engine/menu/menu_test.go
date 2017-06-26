package menu

import (
	"log"
	"testing"
)

const (
	menuWidth    int32 = 640
	menuHeight   int32 = 480
	buttonWidth  int32 = menuWidth / 2
	buttonHeight int32 = 48

	testFontFile string = "../../assets/fonts/pixelated.ttf"
	badFontFile  string = "../../assets/fonts/missing.ttf"
)

var (
	testMenu Menu
)

func TestNewMenu(t *testing.T) {
	params := NewMenuParams{
		W:        menuWidth,
		H:        menuHeight,
		ButtonW:  buttonWidth,
		ButtonH:  buttonHeight,
		FontFile: testFontFile,
	}
	var err error
	testMenu, err = NewMenu(params)
	if err != nil {
		t.Errorf("error: failed to create menu: %v", err)
	}
}

func TestBadFontFile(t *testing.T) {
	params := NewMenuParams{
		W:        menuWidth,
		H:        menuHeight,
		ButtonW:  buttonWidth,
		ButtonH:  buttonHeight,
		FontFile: badFontFile,
	}
	var err error
	_, err = NewMenu(params)
	if err == nil {
		t.Errorf("error: failed to detect bad font file")
	}
}

func TestFont(t *testing.T) {
	if testMenu.Font() == nil {
		t.Errorf("error: failed to set font")
	}
}

func TestAddButton(t *testing.T) {
	handlerFunc := func() error {
		log.Println("button was clicked!")
		return nil
	}
	testMenu.AddButton("button", handlerFunc)
	if len(testMenu.Buttons()) != 1 {
		t.Errorf("error: failed to create button")
	}
}

func TestClearButtons(t *testing.T) {
	testMenu.ClearButtons()
	if len(testMenu.Buttons()) != 0 {
		t.Errorf("error: failed to clear buttons")
	}
}

func TestCursorPos(t *testing.T) {
	if testMenu.CursorPos() != 0 {
		t.Errorf("error: failed to set cursor position")
	}
}

func TestCursorRect(t *testing.T) {
	// Test without buttons (should fail).
	_, err := testMenu.CursorRect()
	if err == nil {
		t.Errorf("error: attempted to create cursor rect with no buttons")
	}

	// Test with a button.
	testMenu.AddButton("button", func() error { return nil })
	_, _ = testMenu.CursorRect()
}

func TestCursorUp(t *testing.T) {
	// Test multiple numbers of buttons, including once with no buttons.
	testMenu.ClearButtons()
	for i := 0; i < 5; i++ {
		p1 := testMenu.CursorPos()
		testMenu.CursorUp()
		p2 := testMenu.CursorPos()

		if len(testMenu.Buttons()) == 0 {
			if p1 != 0 || p2 != 0 {
				t.Errorf("error: moved cursor with no buttons present")
			}
		} else if p1 == 0 {
			// If p1 is the top of the list, p2 must be the bottom after wrapping.
			if p2 != len(testMenu.Buttons())-1 {
				t.Errorf("error: failed to wrap cursor from top")
			}
		} else if p1 != p2+1 {
			// Otherwise, p2 should be one index lower (higher button).
			log.Println("p1:", p1, " p2:", p2)
			t.Errorf("error: failed to move cursor up")
		}

		buttonFunc := func() error { return nil }
		testMenu.AddButton("testbutton", buttonFunc)
	}
}

func TestCursorDown(t *testing.T) {
	// Test multiple numbers of buttons, including once with no buttons.
	testMenu.ClearButtons()
	testMenu.ResetCursor()
	for i := 0; i < 5; i++ {
		p1 := testMenu.CursorPos()
		testMenu.CursorDown()
		p2 := testMenu.CursorPos()
		if len(testMenu.Buttons()) == 0 {
			if p1 != 0 || p2 != 0 {
				t.Errorf("error: moved cursor with no buttons present")
			}
		} else if p1 == len(testMenu.Buttons())-1 {
			// If p1 is at the bottom of the list, p2 must be at the top after wrapping.
			if p2 != 0 {
				t.Errorf("error: failed to wrap cursor from bottom")
			}
		} else if p1 != p2-1 {
			// Otherwise, p2 should be one index higher (lower button).
			t.Errorf("error: failed to move cursor down")
		}

		buttonFunc := func() error { return nil }
		testMenu.AddButton("testbutton", buttonFunc)
	}
}
