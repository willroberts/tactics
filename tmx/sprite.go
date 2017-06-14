package tmx

import (
	"image"
	"image/png"
	"os"

	"github.com/oliamb/cutter"
	"github.com/veandco/go-sdl2/sdl"
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

func (s *spritesheet) ImgToTexture() *sdl.Texture {
	return &sdl.Texture{}
	/* Borrowed example code from go-sdl2/img.
	// Loads a PNG file, returning *sdl.Surface.
	// Then converts *sdl.Surface to *sdl.Texture.
	// How can I create *sdl.Texture from image.Image?
	// Draws *sdl.Texture with renderer.Copy().
	image, err := img.Load(imageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return 3
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	src = sdl.Rect{0, 0, 512, 512}
	dst = sdl.Rect{100, 50, 512, 512}

	renderer.Clear()
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.Delay(2000)
	*/
}
