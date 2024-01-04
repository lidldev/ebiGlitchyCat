package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type char struct {
	x  int
	y  int
	vx int
	vy int

	s      *ebiten.Image
	camera camera
}

const (
	groundY = 290
	unit    = 10
)

func (c *char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *char) update() {
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

func (c *char) draw(screen *ebiten.Image) {
	c.s = assets.MainSprite

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.2, 1.2)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	c.camera.draw(c.s, op)
	c.camera.render(screen)
}

type Player struct {
	player *char
	camera camera
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = -5 * unit // p.camera.vx changing has been cut
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = 5 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.player.tryJump()
	}
	p.player.update()
	p.camera.movement()
	p.camera.x = p.player.x // This is added to make the camera focus on the player's position
	p.camera.y = p.player.y
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.player.draw(screen)
}
