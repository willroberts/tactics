package unit

import (
	"log"
	"testing"
)

func TestUnit(t *testing.T) {
	u := NewRandomizedUnit()
	log.Println("unit:", u)
}
