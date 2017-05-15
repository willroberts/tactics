package unit

import (
	"math/rand"
	"time"
)

// Unit is the interface for all units in the game.
type Unit interface {
}

type unit struct {
	Name string

	BaseAttack    uint
	CurrentAttack uint

	BaseLife    uint
	CurrentLife uint
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRandomizedUnit creates new unit with randomly-generated attributes.
func NewRandomizedUnit() Unit {
	u := unit{}
	attack := uint(rand.Intn(9)) + 1 // 1-9
	life := uint(rand.Intn(9)) + 1   // 1-9

	u.Name = "Unit"
	u.BaseAttack = attack
	u.CurrentAttack = attack
	u.BaseLife = life
	u.CurrentLife = life

	return u
}

// NewUnit creates a new unit with the given attributes.
func NewUnit(name string, attack, life uint) Unit {
	u := unit{}

	u.Name = name
	u.BaseAttack = attack
	u.CurrentAttack = attack
	u.BaseLife = life
	u.CurrentLife = life

	return u
}
