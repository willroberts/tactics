package tmx

import (
	"errors"
	"os"

	gotmx "github.com/salviati/go-tmx/tmx"
)

type dimensions struct {
	W     int
	H     int
	TileW int
	TileH int
}

// GetDimensions returns the width, height, tile width, and tile height of a
// Tiled map.
func GetDimensions(m *gotmx.Map) (*dimensions, error) {
	if len(m.Tilesets) == 0 {
		return &dimensions{}, errors.New("no tilesets in this map")
	}
	if len(m.Tilesets) > 1 {
		return &dimensions{}, errors.New("more than one tileset")
	}
	d := &dimensions{
		W:     m.Tilesets[0].Image.Width,
		H:     m.Tilesets[0].Image.Height,
		TileW: m.TileWidth,
		TileH: m.TileHeight,
	}
	return d, nil
}

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
