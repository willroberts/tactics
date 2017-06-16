package tmx

import (
	"os"

	gotmx "github.com/salviati/go-tmx/tmx"
)

// Dimensions stores the width, height, tile width, and tile height of a Tiled
// map.
type Dimensions struct {
	W     int
	H     int
	TileW int
	TileH int
}

// GetDimensions returns the Dimensions struct for a Tiled map.
func GetDimensions(filename string) (*Dimensions, error) {
	m, err := getMap(filename)
	if err != nil {
		return &Dimensions{}, err
	}
	d := &Dimensions{
		// FIXME: Don't hardcore slice index.
		W:     m.Tilesets[0].Image.Width,
		H:     m.Tilesets[0].Image.Height,
		TileW: m.TileWidth,
		TileH: m.TileHeight,
	}
	return d, nil
}

// GetMap returns a decoded Tiled map from the given filename.
func getMap(filename string) (*gotmx.Map, error) {
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
