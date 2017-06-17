package grid

import "testing"

func TestCanOccupy(t *testing.T) {
	o := &occupier{}
	if !o.CanOccupy() {
		t.Errorf("error: occupier type does not satisfy Occupier interface")
	}
}
