package tmx

import (
	"image"
	"image/png"
	"os"

	"github.com/oliamb/cutter"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

type Spritesheet interface {
}

type spritesheet struct {
	image  image.Image
	width  int
	height int

	sprites      []image.Image
	spriteWidth  int
	spriteHeight int
}

func (s *spritesheet) Image() image.Image {
	return s.image
}

func (s *spritesheet) Width() int {
	return s.width
}

func (s *spritesheet) Height() int {
	return s.height
}

func (s *spritesheet) SpriteWidth() int {
	return s.spriteWidth
}

func (s *spritesheet) SpriteHeight() int {
	return s.spriteHeight
}

func (s *spritesheet) LoadImage(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	i, err := png.Decode(f)
	if err != nil {
		return err
	}
	s.image = i
	return nil
}

func (s *spritesheet) FromTileset(filename string) error {
	m, err := GetMap(filename)
	if err != nil {
		return err
	}
	// FIXME: Don't hardcode slice index.
	s.width = m.Tilesets[0].Image.Width
	s.height = m.Tilesets[0].Image.Height
	s.spriteWidth = m.TileWidth
	s.spriteHeight = m.TileHeight
	return nil
}

func (s *spritesheet) Cut() {
	for x := 0; x < s.width; x += s.spriteWidth {
		for y := 0; y < s.height; y += s.spriteHeight {
			sub, err := cutter.Crop(s.image, cutter.Config{
				Width:   s.spriteWidth,
				Height:  s.spriteHeight,
				Anchor:  image.Point{x, y},
				Options: cutter.Copy,
			})
			s.sprites = append(s.sprites, sub)
		}
	}
}
