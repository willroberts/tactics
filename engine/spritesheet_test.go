package engine

import (
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
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

func TestImage(t *testing.T) {
	if s.Image() == nil {
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

func TestWidth(t *testing.T) {
	if s.Width() != 400 {
		t.FailNow()
	}
}

func TestHeight(t *testing.T) {
	if s.Height() != 260 {
		t.FailNow()
	}
}

func TestSpriteWidth(t *testing.T) {
	if s.SpriteWidth() != 20 {
		t.FailNow()
	}
}

func TestSpriteHeight(t *testing.T) {
	if s.SpriteHeight() != 20 {
		t.FailNow()
	}
}

func TestSprites(t *testing.T) {
}

func TestAddSprite(t *testing.T) {
}

func TestPopulateSprites(t *testing.T) {
	s.PopulateSprites()
}

func TestTextures(t *testing.T) {
	_ = s.textures
}

func TestCreateTexture(t *testing.T) {

}

func TestAddTexture(t *testing.T) {
	tex := &sdl.Texture{}
	count := len(s.Textures())
	s.AddTexture(tex)
	if len(s.Textures()) != count+1 {
		t.FailNow()
	}
}

func TestDestroyTextures(t *testing.T) {
	s.DestroyTextures()
}
