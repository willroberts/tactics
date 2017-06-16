package tmx

import (
	"testing"
)

const (
	tiledFile   string = "testassets/grass.tmx"
	badFile     string = "README.md"
	missingFile string = "missing.file"
)

func TestGetDimensions(t *testing.T) {
	d, err := GetDimensions(tiledFile)
	if err != nil {
		t.FailNow()
	}

	// Fixture tests.
	if d.W != 400 || d.H != 260 || d.TileW != 20 || d.TileH != 20 {
		t.FailNow()
	}

	d, err = GetDimensions(badFile)
	if err == nil {
		t.FailNow()
	}
}

func TestGetMap(t *testing.T) {
	_, err := getMap(tiledFile)
	if err != nil {
		t.FailNow()
	}

	_, err = getMap(badFile)
	if err == nil {
		t.FailNow()
	}

	_, err = getMap(missingFile)
	if err == nil {
		t.FailNow()
	}
}
