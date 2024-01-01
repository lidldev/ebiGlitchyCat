package game

import "github.com/hajimehoshi/ebiten/v2"

type camera struct {
	x  int // x position of the camera
	y  int // y position of the camera
	vx int
	vy int

	drawable *ebiten.Image // the image that the camera will draw
}

const (
	outsideWidth  = 800
	outsideHeight = 600
)

func (c *camera) movement() {
	c.x += c.vx

	if c.vx > 0 {
		c.vx -= 10
	} else if c.vx < 0 {
		c.vx += 10
	}
}

func (c *camera) setPos(x, y int) {
	c.x += x
	c.y += y
}

func (c *camera) init() {
	c.drawable = ebiten.NewImage(outsideWidth, outsideHeight)
}

func (camera *camera) render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(camera.drawable, op)
}

func (c *camera) draw(image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(-c.x), float64(-c.y))

	c.drawable.DrawImage(image, op)
}

func (c *camera) clear() {
	c.drawable.Clear()
}
