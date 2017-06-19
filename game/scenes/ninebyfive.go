package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/game/unit"
	"github.com/willroberts/tactics/grid"
)

const (
	gridW int = 9
	gridH int = 5
	cellW int = 100
	cellH int = 100

	cLtBlue uint32 = 0xff6495ed
	cDkBlue uint32 = 0xff4682b4
	cWhite  uint32 = 0xffffffff
)

type nineByFiveScene struct {
	eng  engine.SDLEngine
	grid grid.Grid
}

func (s *nineByFiveScene) Setup() error {
	s.eng.Window().SetSize(900, 500)
	s.grid = grid.NewGrid(gridW, gridH, cellW, cellH)
	s.grid.Checkerboard(cLtBlue, cDkBlue)

	s.grid.Cell(0, 2).SetContents(unit.NewUnit())
	s.grid.Cell(4, 0).SetContents(unit.NewUnit())
	s.grid.Cell(4, 4).SetContents(unit.NewUnit())
	s.grid.Cell(5, 2).SetContents(unit.NewUnit())
	s.grid.Cell(8, 2).SetContents(unit.NewUnit())

	return nil
}

func (s *nineByFiveScene) Main() error {
	err := s.eng.ClearScreen()
	if err != nil {
		return err
	}

	for _, col := range s.grid.Cells() {
		for _, cell := range col {
			d := cell.Dimensions()
			err = s.eng.DrawRect(&sdl.Rect{
				X: int32(d.X),
				Y: int32(d.Y),
				W: int32(d.W),
				H: int32(d.H),
			}, cell.Color())
			if err != nil {
				return err
			}

			if cell.IsOccupied() {
				err = s.eng.DrawRect(&sdl.Rect{
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

func (s *nineByFiveScene) Teardown() error {
	return nil
}

func NewNineByFiveScene(e engine.SDLEngine) engine.Scene {
	return &nineByFiveScene{eng: e}
}
