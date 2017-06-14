package tmx

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

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
	defer f.Close()
	i, _, err := image.Decode(f)
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

func (s *spritesheet) Cut() error {
	// Read original image.
	bounds := s.image.Bounds()
	rgba := image.NewRGBA(bounds)
	for x := 0; x < s.width; x += s.spriteWidth {
		for y := 0; y < s.height; y += s.spriteHeight {
			subrect := image.Rectangle{
				Min: image.Point{x, y},
				Max: image.Point{x + s.spriteWidth, y + s.spriteHeight},
			}
			subimage := rgba.SubImage(subrect)
			s.sprites = append(s.sprites, subimage)
		}
	}

	// Write new images.
	for i, s := range s.sprites {
		f, err := os.Create(fmt.Sprintf("testassets/%d.png", i))
		if err != nil {
			return err
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		err = png.Encode(w, s)
		if err != nil {
			return err
		}
		w.Flush()
	}

	return nil
}
