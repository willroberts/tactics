package tmx

import (
	"log"
	"testing"
)

var (
	s *spritesheet
)

func TestLoadImage(t *testing.T) {
	s = &spritesheet{}
	err := s.LoadImage("testassets/floor-tiles-20x20.png")
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestFromTileset(t *testing.T) {
	err := s.FromTileset("testassets/grass.tmx")
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestCut(t *testing.T) {
	s.Cut()
}
