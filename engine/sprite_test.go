package engine

import (
	"image"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	pngFile     string = "../assets/sprites/tilecrusader/tiles.png"
	badPngFile  string = "../assets/sprites/bad.png"
	invalidFile string = "../assets/sprites/missing.png"

	tiledFile       string = "../assets/maps/map.tmx"
	noTilesetsFile  string = "../assets/maps/notilesets.tmx"
	twoTilesetsFile string = "../assets/maps/twotilesets.tmx"
)

var (
	s *spritesheet
)

func TestLoadImage(t *testing.T) {
	s = &spritesheet{}
	err := s.LoadImage(badPngFile)
	if err == nil {
		t.Errorf("error: failed to detect bad png")
	}
	err = s.LoadImage(invalidFile)
	if err == nil {
		t.Errorf("error: failed to detect invalid file")
	}
	err = s.LoadImage(pngFile)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestImage(t *testing.T) {
	if s.Image() == nil {
		t.Errorf("error: image is nil")
	}
}

func TestPopulateDimensions(t *testing.T) {
	err := s.PopulateDimensions(noTilesetsFile)
	if err == nil {
		t.Errorf("error: failed to detect bad tiled file")
	}
	err = s.PopulateDimensions(twoTilesetsFile)
	if err == nil {
		t.Errorf("error: failed to detect bad tiled file")
	}
	err = s.PopulateDimensions(tiledFile)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestDimensions(t *testing.T) {
	if s.Width() != 400 || s.Height() != 260 {
		t.Errorf("error: unexpected image dimensions")
	}
}

func TestSpriteDimensions(t *testing.T) {
	if s.SpriteWidth() != 20 || s.SpriteHeight() != 20 {
		t.Errorf("error: unexpected sprite dimensions")
	}
}

func TestPopulateSprites(t *testing.T) {
	s.PopulateSprites()
}

func TestSprites(t *testing.T) {
	if len(s.Sprites()) != 260 {
		t.Errorf("error: unexpected number of sprites")
	}
}

func TestAddSprite(t *testing.T) {
	spr := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{1, 1},
	})
	count := len(s.Sprites())
	s.AddSprite(spr)
	if len(s.Sprites()) != count+1 {
		t.Errorf("error: unexpected number of sprites")
	}
}

func TestTextures(t *testing.T) {
	_ = s.textures
}

func TestCreateTexture(t *testing.T) {
	eng, err := NewSDLEngine("test")
	if err != nil {
		t.Errorf("error: failed to create engine: %v", err)
	}

	i := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{20, 20},
	})

	_, err = s.CreateTexture(i, eng.Renderer())
	if err != nil {
		t.Errorf("error: failed to create texture: %v", err)
	}
}

func TestAddTexture(t *testing.T) {
	tex := &sdl.Texture{}
	count := len(s.Textures())
	s.AddTexture(tex)
	if len(s.Textures()) != count+1 {
		t.Errorf("error: unexpected number of textures")
	}
}

func TestDestroyTextures(t *testing.T) {
	s.DestroyTextures()
}
