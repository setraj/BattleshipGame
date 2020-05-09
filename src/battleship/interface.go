package battleship

import (
	"fmt"
	"strings"
)

//Interface is a sort of engine which manages the game between Players
type Interface struct {
	g1 *Game
	g2 *Game
}

func NewGame(p1, p2 Player) *Interface {
	return &Interface{
		g1: &Game{
			player: p1,
			board:  NewBoard(),
			id:     1,
		},
		g2: &Game{
			player: p2,
			board:  NewBoard(),
			id:     2,
		},
	}
}
func (i *Interface) Run() {

	fmt.Println("Starting the Game...\n")
	fmt.Println("Player 1 Board\n")
	i.g1.SetupShips()
	i.g1.board.Draw()
	fmt.Println("\nPlayer 2 Board")
	i.g2.SetupShips()
	i.g2.board.Draw()
	cell := ""
start:
	for {
		fmt.Println("Player 1 Enter the position for fire: ")
		fmt.Scan(&cell)
		if len(cell) != 2 {
			fmt.Println("INVALID POSITION.")
			goto start
		}
		cell = strings.ToUpper(cell)
		x := int(cell[0] - 65)
		y := int(cell[1] - 48)
		c := Cell{
			X: x,
			Y: y,
		}
		//Start with Player 1
		resp := i.g2.Fire(c)

		i.g2.board.Draw()
		if resp == HIT {
			//check if Player 2 has lost
			if i.g2.IsGameOver() {
				fmt.Println("Player 1 WON!")
				break
			}
		}
		fmt.Println("Player 2 Enter the position for fire: ")
		fmt.Scan(&cell)
		if len(cell) != 2 {
			fmt.Println("INVALID POSITION.")
			goto start
		}
		cell = strings.ToUpper(cell)
		x = int(cell[0] - 65)
		y = int(cell[1] - 48)
		c = Cell{
			X: x,
			Y: y,
		}
		//Player 2 fires
		resp = i.g1.Fire(c)

		i.g1.board.Draw()
		if resp == HIT {
			//check if Player 1 has lost
			if i.g1.IsGameOver() {
				fmt.Println("Player 2 WON!")
				break
			}
		}
	}

}
