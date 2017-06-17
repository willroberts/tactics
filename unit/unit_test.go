package unit

import (
	"testing"
)

func TestNewUnit(t *testing.T) {
	u := NewUnit()
	if u.CanOccupy() != true {
		t.Errorf("error: unit type does not satisfy Occupier interface")
	}
	if u.Name() != "unit" {
		t.Errorf("error: unexpected unit name")
	}
}
