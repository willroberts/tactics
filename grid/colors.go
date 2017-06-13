package grid

const (
	ColorWhite uint32 = 0xffffffff
	ColorGreen uint32 = 0xff00ff00
	ColorRed   uint32 = 0xffff0000
)

// AssignColors creates a checkerboard pattern in the grid, like below:
/*   1  2  3
  1  G  W  G
	2  W  G  W
	3  G  W  G */
func AssignColors(g Grid) {
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			if x%2 == 0 {
				if y%2 == 0 {
					g.Cell(x, y).SetColor(ColorGreen)
				} else {
					g.Cell(x, y).SetColor(ColorWhite)
				}
			} else {
				if y%2 == 0 {
					g.Cell(x, y).SetColor(ColorWhite)
				} else {
					g.Cell(x, y).SetColor(ColorGreen)
				}
			}
		}
	}
}
