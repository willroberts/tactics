package tmx

import (
	"testing"

	gotmx "github.com/salviati/go-tmx/tmx"
)

const (
	tiledFile   string = "testassets/grass.tmx"
	badFile     string = "README.md"
	missingFile string = "missing.file"
)

var (
	tiledMap *gotmx.Map
)

func TestGetMap(t *testing.T) {
	m, err := GetMap(tiledFile)
	if err != nil {
		t.Errorf("error: failed to get map: %v", err)
	}
	tiledMap = m

	_, err = GetMap(badFile)
	if err == nil {
		t.Errorf("error: failed to detect bad map")
	}

	_, err = GetMap(missingFile)
	if err == nil {
		t.Errorf("error: failed to detect missing map")
	}
}

func TestGetDimensions(t *testing.T) {
	d, err := GetDimensions(tiledMap)
	if err != nil {
		t.Errorf("error: failed to get map dimensions: %v", err)
	}

	// Fixture tests.
	if d.W != 400 || d.H != 260 || d.TileW != 20 || d.TileH != 20 {
		t.Errorf("error: unexpected map dimensions")
	}

	badMap := tiledMap
	badMap.Tilesets = append(badMap.Tilesets, badMap.Tilesets[0])
	if _, err = GetDimensions(badMap); err == nil {
		t.Errorf("error: failed to return error on map with >1 tileset")
	}

	badMap.Tilesets = []gotmx.Tileset{}
	if _, err = GetDimensions(badMap); err == nil {
		t.Errorf("error: failed to return error on map with no tileset")
	}
}
