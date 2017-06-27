package engine

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func (e *sdlengine) Gradient(r *sdl.Rect, c1 uint32, c2 uint32) error {
	colorStep := (c2 - c1) / uint32(r.H)
	log.Println("cs:", colorStep)
	c := c1

	for y := r.Y; y <= r.Y+r.H; y++ {
		rgba := colorToRGBA(c)
		log.Println("rgba:", rgba)
		err := e.renderer.SetDrawColor(rgba.R, rgba.G, rgba.B, rgba.A)
		if err != nil {
			return err
		}

		err = e.renderer.DrawLine(int(r.X), int(y), int(r.X+r.W), int(y))
		if err != nil {
			return err
		}

		c += colorStep
	}

	return nil
}

// TODO: Pull this functionality from above.
func interpolateColor(c1, c2, pos, max uint32) uint32 {
	colorStep := (c2 - c1) * (pos / max)
	return colorStep
}
