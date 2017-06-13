package grid

// Occupier is the common interface among anything which can occupy a cell
// within a grid. To place objects in a cell, have your objects satisfy this
// interface.
type Occupier interface {
	CanOccupy() bool
}

type occupier struct{}

func (o *occupier) CanOccupy() bool {
	return true
}
