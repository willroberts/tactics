package tmx

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/oliamb/cutter"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

type Spritesheet interface {
	Image() image.Image
	Width() int
	Height() int

	Sprites() []image.Image
	SpriteWidth() int
	SpriteHeight() int

	LoadImage(string) error
	PopulateDimensions(string) error
	PopulateSprites()

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

func (s *spritesheet) Sprites() []image.Image {
	return s.sprites
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

func (s *spritesheet) PopulateDimensions(filename string) error {
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

// NOTE: Due to the SDL2 API requiring a string rather than a buffer or Reader
// for img.Load(), we temporarily write an image to disk. :(
func CreateTexture(i image.Image, r *sdl.Renderer) (*sdl.Texture, error) {
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
