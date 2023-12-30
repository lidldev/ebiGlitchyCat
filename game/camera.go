package game

import "github.com/hajimehoshi/ebiten/v2"

type camera struct {
	x int // x position of the camera
	y int // y position of the camera

	drawable *ebiten.Image // the image that the camera will draw
}

func newCamera() *camera {
	c := &camera{}

	return c
}

func (c *camera) init() {

}

func (c *camera) update() {

}

func (c *camera) draw() {

}
