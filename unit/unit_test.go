package unit

import (
	"testing"
)

func TestNewUnit(t *testing.T) {
	u := NewUnit()
	if u.CanOccupy() != true {
		t.FailNow()
	}
	if u.Name() != "unit" {
		t.FailNow()
	}
}
