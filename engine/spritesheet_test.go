package engine

import (
	"log"
	"testing"
)

const (
	spritesheetFile string = "testassets/spritesheet.png"
	tiledFile       string = "testassets/grass.tmx"
)

var (
	s *spritesheet
)

func TestLoadImage(t *testing.T) {
	s = &spritesheet{}
	err := s.LoadImage(spritesheetFile)
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestPopulateDimensions(t *testing.T) {
	err := s.PopulateDimensions(tiledFile)
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestPopulateSprites(t *testing.T) {
	s.PopulateSprites()
}
