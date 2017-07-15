package item

type Weapon interface {
	Name() string
	Damage() int
	Accuracy() float64
	Reach() int
	Strikes() int
	Area() int
}

type weapon struct {
	name     string
	damage   int
	accuracy float64
	reach    int // Wanted to name this 'range' but it's a reserved word.
	strikes  int
	area     int
}

func (w *weapon) Name() string {
	return w.name
}

func (w *weapon) Damage() int {
	return w.damage
}

func (w *weapon) Accuracy() float64 {
	return w.accuracy
}

func (w *weapon) Reach() int {
	return w.reach
}

func (w *weapon) Strikes() int {
	return w.strikes
}

func (w *weapon) Area() int {
	return w.area
}

type NewWeaponParams struct {
	Name     string
	Damage   int
	Accuracy float64
	Reach    int
	Strikes  int
	Area     int
}

func NewWeapon(p NewWeaponParams) Weapon {
	return &weapon{
		name:     p.Name,
		damage:   p.Damage,
		accuracy: p.Accuracy,
		reach:    p.Reach,
		strikes:  p.Strikes,
		area:     p.Area,
	}
}

func InitializeWeapons() map[string]Weapon {
	resp := make(map[string]Weapon)

	params := NewWeaponParams{
		Name:     "Bebut",
		Damage:   60,
		Accuracy: 1.00,
		Reach:    1,
		Strikes:  1,
		Area:     1,
	}
	resp["bebut"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Shashka",
		Damage:   100,
		Accuracy: 0.75,
		Reach:    1,
		Strikes:  1,
		Area:     1,
	}
	resp["shashka"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Revolver",
		Damage:   60,
		Accuracy: 0.80,
		Reach:    4,
		Strikes:  1,
		Area:     1,
	}
	resp["revolver"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Grenade",
		Damage:   80,
		Accuracy: 1.00,
		Reach:    5,
		Strikes:  1,
		Area:     5,
	}
	resp["grenade"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Rifle",
		Damage:   80,
		Accuracy: 0.80,
		Reach:    6,
		Strikes:  1,
		Area:     1,
	}
	resp["rifle"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Repeater",
		Damage:   50,
		Accuracy: 0.50,
		Reach:    5,
		Strikes:  2,
		Area:     1,
	}
	resp["repeater"] = NewWeapon(params)

	params = NewWeaponParams{
		Name:     "Machine Gun",
		Damage:   20,
		Accuracy: 0.20,
		Reach:    5,
		Strikes:  8,
		Area:     1,
	}
	resp["machinegun"] = NewWeapon(params)

	return resp
}
