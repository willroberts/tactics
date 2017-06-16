package engine

import (
	"image"
	"log"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	pngFile    string = "testassets/spritesheet.png"
	badPngFile string = "testassets/bad.png"

	tiledFile       string = "testassets/grass.tmx"
	noTilesetsFile  string = "testassets/zero.tmx"
	twoTilesetsFile string = "testassets/two.tmx"

	invalidFile string = "testassets/nonexistent.png"
)

var (
	s *spritesheet
)

func TestLoadImage(t *testing.T) {
	s = &spritesheet{}
	err := s.LoadImage(badPngFile)
	if err == nil {
		log.Println("error: failed to detect bad png")
		t.FailNow()
	}
	err = s.LoadImage(invalidFile)
	if err == nil {
		log.Println("error: failed to detect invalid file")
		t.FailNow()
	}
	err = s.LoadImage(pngFile)
	if err != nil {
		log.Println("error:", err)
		t.FailNow()
	}
}

func TestImage(t *testing.T) {
	if s.Image() == nil {
		log.Println("image: image is nil")
		t.FailNow()
	}
}

func TestPopulateDimensions(t *testing.T) {
	err := s.PopulateDimensions(noTilesetsFile)
	if err == nil {
		t.Errorf("failed to detect bad tiled file")
	}
	err = s.PopulateDimensions(twoTilesetsFile)
	if err == nil {
		t.Errorf("failed to detect bad tiled file")
	}
	err = s.PopulateDimensions(tiledFile)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestWidth(t *testing.T) {
	if s.Width() != 400 {
		log.Println("error: unexpected width")
		t.FailNow()
	}
}

func TestHeight(t *testing.T) {
	if s.Height() != 260 {
		log.Println("error: unexpected height")
		t.FailNow()
	}
}

func TestSpriteWidth(t *testing.T) {
	if s.SpriteWidth() != 20 {
		log.Println("error: unexpected sprite width")
		t.FailNow()
	}
}

func TestSpriteHeight(t *testing.T) {
	if s.SpriteHeight() != 20 {
		log.Println("error: unexpected sprite height")
		t.FailNow()
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
		t.FailNow()
	}
}

func TestTextures(t *testing.T) {
	_ = s.textures
}

// FIXME: Determine how to test failures of img.Load() on temp file.
// FIXME: Determine how to test failures of r.CreateTextureFromSurface().
func TestCreateTexture(t *testing.T) {
	eng, err := NewSDLEngine("test", 400, 400)
	if err != nil {
		t.Errorf("failed to create engine: %v", err)
	}

	i := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{20, 20},
	})

	_, err = s.CreateTexture(i, eng.Renderer())
	if err != nil {
		t.Errorf("failed to create texture: %v", err)
	}
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
