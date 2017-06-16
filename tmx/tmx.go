package tmx

import (
	"os"

	gotmx "github.com/salviati/go-tmx/tmx"
)

type dimensions struct {
	W     int
	H     int
	TileW int
	TileH int
}

func GetDimensions(filename string) (*dimensions, error) {
	m, err := getMap(filename)
	if err != nil {
		return &dimensions{}, err
	}
	d := &dimensions{
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
