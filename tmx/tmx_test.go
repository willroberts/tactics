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
		t.FailNow()
	}
	tiledMap = m

	_, err = GetMap(badFile)
	if err == nil {
		t.FailNow()
	}

	_, err = GetMap(missingFile)
	if err == nil {
		t.FailNow()
	}
}

func TestGetDimensions(t *testing.T) {
	d, err := GetDimensions(tiledMap)
	if err != nil {
		t.FailNow()
	}

	// Fixture tests.
	if d.W != 400 || d.H != 260 || d.TileW != 20 || d.TileH != 20 {
		t.FailNow()
	}

	badMap := tiledMap
	badMap.Tilesets = append(badMap.Tilesets, badMap.Tilesets[0])
	if _, err = GetDimensions(badMap); err == nil {
		t.FailNow()
	}

	badMap.Tilesets = []gotmx.Tileset{}
	if _, err = GetDimensions(badMap); err == nil {
		t.FailNow()
	}
}
