package game

// import "github.com/hajimehoshi/ebiten/v2"

// type camera struct {
// 	x int // x position of the camera
// 	y int // y position of the camera

// 	drawable *ebiten.Image // the image that the camera will draw
// }

// const (
// 	outsideWidth  = 800
// 	outsideHeight = 600
// )

// func newCamera() *camera {
// 	c := &camera{}

// 	return c
// }

// func (c *camera) setPos(x, y int) {
// 	c.x += x
// 	c.y += y
// }

// func (c *camera) init() {
// 	c.drawable = ebiten.NewImage(outsideWidth, outsideHeight)
// }

// func (camera *camera) render(screen *ebiten.Image) {
// 	op := &ebiten.DrawImageOptions{}

// 	screen.DrawImage(camera.drawable, op)
// }

// func (c *camera) update() {
// 	if ebiten.IsKeyPressed(ebiten.KeyA) {
// 		c.setPos(-1, 0)
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyD) {
// 		c.setPos(1, 0)
// 	}
// }
// func (c *camera) draw(image *ebiten.Image, op *ebiten.DrawImageOptions) {
// 	op.GeoM.Translate(float64(-c.x), float64(-c.y))

// 	c.drawable.DrawImage(image, op)
// }
