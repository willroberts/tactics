package unit

// Unit is the interface for all units in the game. Satisfies the grid.Occupier
// interface.
type Unit interface {
	CanOccupy() bool
	Name() string
}

type unit struct {
	name string
	//rect  *sdl.Rect
	//image *sdl.Texture
}

func (u *unit) CanOccupy() bool {
	return true
}

func (u *unit) Name() string {
	return "unit"
}

// NewUnit creates and returns a Unit which satisfies grid.Occupier.
func NewUnit() Unit {
	return &unit{}
}
