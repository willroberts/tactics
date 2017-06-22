package engine

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	c Camera
)

func TestCamera(t *testing.T) {
	c = NewCamera()
	c.MoveTo(10, 20)
	cx, cy := c.Position()
	if cx != 10 {
		t.Errorf("error: failed to move camera x")
	}
	if cy != 20 {
		t.Errorf("error: failed to move camera y")
	}
}

func TestShiftRect(t *testing.T) {
	r := &sdl.Rect{X: 20, Y: 30, W: 10, H: 10}
	r = c.ShiftRect(r)
	if r.X != 10 {
		t.Errorf("error: failed to shift rect x")
	}
	if r.Y != 10 {
		t.Errorf("error: failed to shift rect y")
	}
}

func TestShiftVectors(t *testing.T) {
	vx := []int16{100, 110, 120, 130}
	vy := []int16{200, 210, 220, 230}
	vx, vy = c.ShiftVectors(vx, vy)

	expectedX := []int16{90, 100, 110, 120}
	for i, x := range vx {
		if x != expectedX[i] {
			t.Errorf("failed to shift vector x")
		}
	}

	expectedY := []int16{180, 190, 200, 210}
	for i, y := range vy {
		if y != expectedY[i] {
			t.Errorf("failed to shift vector y")
		}
	}
}
