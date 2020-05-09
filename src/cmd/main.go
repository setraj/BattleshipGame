package main

import (
	"battleship"
)

//Start of the game
func main() {
	player1 := battleship.Player{}
	player2 := battleship.Player{}
	game := battleship.NewGame(player1, player2)
	game.Run()
}
