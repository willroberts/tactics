package tmx

import (
	"os"

	gotmx "github.com/salviati/go-tmx/tmx"
)

// GetMap returns a decoded Tiled map from the given filename.
func GetMap(filename string) (*gotmx.Map, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &gotmx.Map{}, err
	}
	m, err := gotmx.Read(f)
	if err != nil {
		return &gotmx.Map{}, err
	}
	return m, nil
}
