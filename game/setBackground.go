package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct{}

func (b *Background) intialBackground(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
}
