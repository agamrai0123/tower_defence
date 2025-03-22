package main

import (
	"log"

	"github.com/agamrai0123/tower_defence/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Dragon Defence")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
