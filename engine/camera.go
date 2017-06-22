package engine

type Camera interface {
	MoveTo(int32, int32)
	Position() (int32, int32)
}

type camera struct {
	x int32
	y int32
}

func (c *camera) MoveTo(x, y int32) {
	c.x = x
	c.y = y
}

func (c *camera) Position() (int32, int32) {
	return c.x, c.y
}

func NewCamera() Camera {
	return &camera{x: 0, y: 0}
}
