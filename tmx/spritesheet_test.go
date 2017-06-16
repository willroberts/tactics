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

func TestPopulateDimensions(t *testing.T) {
	err := s.PopulateDimensions("testassets/grass.tmx")
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestPopulateSprites(t *testing.T) {
	s.PopulateSprites()
}
