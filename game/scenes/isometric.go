package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/game/unit"
	"github.com/willroberts/tactics/grid"
)

const (
	gridW int = 12
	gridH int = 12
	cellW int = 40
	cellH int = 40

	gameWidth  int32 = 1280
	gameHeight int32 = 720

	cBlack  uint32 = 0xff000000
	cDkGray uint32 = 0xff333333
	cDkBlue uint32 = 0xff4682b4
	cLtBlue uint32 = 0xff6495ed
	cWhite  uint32 = 0xffffffff
)

type isometricScene struct {
	eng  engine.SDLEngine
	grid grid.Grid
}

func (s *isometricScene) Setup() error {
	s.eng.Window().SetSize(int(gameWidth), int(gameHeight))

	// FIXME: Find a better way to derive these values.
	s.eng.Camera().MoveTo(-gameWidth/2, -gameHeight/5)

	s.grid = grid.NewGrid(gridW, gridH, cellW, cellH)
	s.grid.Checkerboard(cLtBlue, cDkBlue)

	s.grid.Cell(0, 2).SetContents(unit.NewUnit())
	s.grid.Cell(4, 0).SetContents(unit.NewUnit())
	s.grid.Cell(4, 4).SetContents(unit.NewUnit())
	s.grid.Cell(5, 2).SetContents(unit.NewUnit())
	s.grid.Cell(8, 2).SetContents(unit.NewUnit())

	return nil
}

func (s *isometricScene) Main() error {
	err := s.eng.ClearScreen()
	if err != nil {
		return err
	}

	err = s.eng.FillWindow(cDkGray)
	if err != nil {
		return err
	}

	for _, col := range s.grid.Cells() {
		for _, cell := range col {
			d := cell.Dimensions()
			err = s.eng.DrawIsometricRect(&sdl.Rect{
				X: int32(d.X),
				Y: int32(d.Y),
				W: int32(d.W),
				H: int32(d.H),
			}, cell.Color())
			if err != nil {
				return err
			}

			if cell.IsOccupied() {
				err = s.eng.DrawIsometricRect(&sdl.Rect{
					X: int32(d.X + (d.W / 4)),
					Y: int32(d.Y + (d.H / 4)),
					W: int32(d.W / 2),
					H: int32(d.H / 2),
				}, cWhite)
				if err != nil {
					return err
				}
			}
		}
	}

	err = s.eng.UpdateSurface()
	if err != nil {
		return err
	}

	s.eng.PauseRendering(fTime)
	return nil
}

func (s *isometricScene) Teardown() error {
	return nil
}

func NewIsometricScene(e engine.SDLEngine) engine.Scene {
	return &isometricScene{eng: e}
}
