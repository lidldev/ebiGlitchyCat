package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Char struct {
	x  int
	y  int
	vx int
	vy int

	s      *ebiten.Image
	camera camera
}

const (
	groundY = 390
	unit    = 10
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *Char) draw(screen *ebiten.Image) {
	c.s = assets.MainSprite

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.8, 0.8)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	c.camera.draw(c.s, op)
}

type Player struct {
	player *Char
	camera camera
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &Char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 5 * unit // p.camera.vx changing has been cut
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -5 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.player.tryJump()
	}
	p.camera.x += p.player.x
	p.camera.y += p.player.y
	p.player.update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.player.draw(screen)
}
