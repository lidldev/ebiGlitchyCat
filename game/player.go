package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
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
	groundY = 100
	unit    = 10
)

// func (c *char) tryJump() {
// 	if c.y == groundY*unit {
// 		c.vy = -10 * unit
// 	}
// }

// func (c *char) update() {
// 	c.x += c.vx
// 	c.y += c.vy

// 	if c.y > groundY*unit {
// 		c.y = groundY * unit
// 	}
// 	if c.vx > 0 {
// 		c.vx -= 2
// 	} else if c.vx < 0 {
// 		c.vx += 2
// 	}
// 	if c.vy < 20*unit {
// 		c.vy += 8
// 	}
// }

func (c *char) draw(screen *ebiten.Image) {
	c.s = assets.MainSprite

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.2, 1.2)
	op.GeoM.Translate(float64(c.camera.x)/unit, float64(c.camera.y)/unit)
	c.camera.draw(c.s, op)
}

type Player struct {
	player *char
	camera camera
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &char{x: 50 * unit, y: groundY * unit}
	}

	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
	// 	p.camera.x += 10
	// 	p.player.vx = -5 * unit
	// } else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	// 	p.camera.y -= 10
	// 	p.player.vx = 5 * unit
	// }
	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	// 	p.player.tryJump()
	// }
	// p.player.update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.player.draw(screen)
}
