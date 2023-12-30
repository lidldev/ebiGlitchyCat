package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowTitle("Glitchy Orange Cat")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
