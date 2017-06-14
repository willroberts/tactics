package tmx

import (
	"image"
	"image/png"
	"os"
)

type spritesheet struct {
	image  image.Image
	width  int
	height int

	sprites      []image.Image
	spriteWidth  int
	spriteHeight int
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
		}
	}
	s.sprites = []image.Image{}
}
