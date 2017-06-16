package tmx

import (
	"testing"
)

func TestGetMap(t *testing.T) {
	_, err := getMap("testassets/grass.tmx")
	if err != nil {
		t.FailNow()
	}
}
