package grid

import "testing"

func TestCanOccupy(t *testing.T) {
	o := &occupier{}
	if !o.CanOccupy() {
		t.FailNow()
	}
}
