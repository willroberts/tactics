package item

import "testing"

func TestInitializeWeapons(t *testing.T) {
	wpns := InitializeWeapons()
	wpnNames := []string{"bebut", "shashka", "revolver", "grenade", "rifle",
		"repeater", "machinegun"}

	for _, n := range wpnNames {
		if _, ok := wpns[n]; !ok {
			t.Errorf("missing weapon: %s", n)
		}

		wpn := wpns[n]
		if wpn.Name() == "" {
			t.Errorf("failed to set weapon name: %v", wpn)
		}
		if wpn.Damage() == 0 {
			t.Errorf("failed to set weapon damage: %v", wpn)
		}
		if wpn.Accuracy() == 0 {
			t.Errorf("failed to set weapon accuracy: %v", wpn)
		}
		if wpn.Reach() == 0 {
			t.Errorf("failed to set weapon reach: %v", wpn)
		}
		if wpn.Strikes() == 0 {
			t.Errorf("failed to set weapon strikes: %v", wpn)
		}
		if wpn.Area() == 0 {
			t.Errorf("failed to set weapon area: %v", wpn)
		}
	}
}
