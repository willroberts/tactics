package engine

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/oliamb/cutter"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/tmx"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

// Spritesheet is the interface for an image containing a grid of smaller
// images. Its interface allows us to carve up the image, and convert each one
// into an SDL texture.
type Spritesheet interface {
	Image() image.Image

	Width() int
	Height() int
	SpriteWidth() int
	SpriteHeight() int

	Sprites() []image.Image
	AddSprite(image.Image)

	LoadImage(string) error
	PopulateDimensions(string) error
	PopulateSprites()

	Textures() []*sdl.Texture
	CreateTexture(image.Image, *sdl.Renderer) (*sdl.Texture, error)
	AddTexture(*sdl.Texture)
	DestroyTextures()
}

type spritesheet struct {
	image  image.Image
	width  int
	height int

	sprites      []image.Image
	textures     []*sdl.Texture
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

func (s *spritesheet) Sprites() []image.Image {
	return s.sprites
}

func (s *spritesheet) AddSprite(i image.Image) {
	s.sprites = append(s.sprites, i)
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

	_ = f.Close()
	return nil
}

func (s *spritesheet) PopulateDimensions(filename string) error {
	d, err := tmx.GetDimensions(filename)
	if err != nil {
		return err
	}
	s.width = d.W
	s.height = d.H
	s.spriteWidth = d.TileW
	s.spriteHeight = d.TileH
	return nil
}

func (s *spritesheet) PopulateSprites() {
	for x := 0; x < s.width; x += s.spriteWidth {
		for y := 0; y < s.height; y += s.spriteHeight {
			sub, err := cutter.Crop(s.image, cutter.Config{
				Width:   s.spriteWidth,
				Height:  s.spriteHeight,
				Anchor:  image.Point{x, y},
				Options: cutter.Copy,
			})
			if err != nil {
				log.Println("failed to crop image:", err)
				continue
			}
			s.sprites = append(s.sprites, sub)
		}
	}
}

func (s *spritesheet) Textures() []*sdl.Texture {
	return s.textures
}

// FIXME: Move to engine package.
// NOTE: Due to the SDL2 API requiring a string rather than a buffer or Reader
// for img.Load(), we temporarily write an image to disk. :(
func (s *spritesheet) CreateTexture(i image.Image, r *sdl.Renderer) (*sdl.Texture, error) {
	filename := ".t.png"
	f, err := os.Create(filename)
	if err != nil {
		return &sdl.Texture{}, err
	}

	if err = png.Encode(f, i); err != nil {
		return &sdl.Texture{}, err
	}
	err = f.Close()
	if err != nil {
		return &sdl.Texture{}, err
	}

	sur, err := img.Load(filename)
	if err != nil {
		return &sdl.Texture{}, err
	}
	defer sur.Free()

	tex, err := r.CreateTextureFromSurface(sur)
	if err != nil {
		return &sdl.Texture{}, err
	}

	_ = os.Remove(filename)
	return tex, nil
}

func (s *spritesheet) AddTexture(t *sdl.Texture) {
	s.textures = append(s.textures, t)
}

func (s *spritesheet) DestroyTextures() {
	for _, t := range s.textures {
		t.Destroy()
	}
}
