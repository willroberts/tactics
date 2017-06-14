package tmx

import (
	"testing"
)

func TestGetMap(t *testing.T) {
	_, err := GetMap("testassets/grass.tmx")
	if err != nil {
		t.FailNow()
	}
}
