package unit

import (
	"math/rand"
	"time"
)

const (
	defaultUnitName string = "Unit"

	attackMin int = 1
	attackMax int = 5
	lifeMin   int = 1
	lifeMax   int = 8
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

// NewRandomizedUnit creates new unit with randomly-generated attributes.
func NewRandomizedUnit() Unit {
	u := unit{}
	attack := uint(rand.Intn(attackMax) + attackMin)
	life := uint(rand.Intn(lifeMax) + lifeMin)

	u.Name = defaultUnitName
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
