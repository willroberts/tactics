// +build video

package engine

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestGradient(t *testing.T) {
	eng, _ := NewSDLEngine("gradient testing")
	defer eng.Window().Destroy()
	_ = eng.ClearScreen()
	_ = eng.FillWindow(0xff333333)

	r := &sdl.Rect{X: 0, Y: 32, W: 1920, H: 32}
	var c1 uint32 = 0xff000000 // black
	var c2 uint32 = 0xff007f00 // dk green?
	err = eng.Gradient(r, c1, c2)
	if err != nil {
		t.Errorf("error 2: %v", err)
	}

	_ = eng.UpdateSurface()
	sdl.Delay(5000)
}
