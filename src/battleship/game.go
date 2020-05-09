package battleship

type Game struct {
	ships  []*Ship
	player Player
	board  Board
	id     int
}

func (g *Game) Fire(c Cell) int {
	if g.board[c.X][c.Y] == SHIP {
		g.board[c.X][c.Y] = HIT
		return HIT
	} else {
		g.board[c.X][c.Y] = MISS
		return MISS
	}
}
func (g *Game) IsGameOver() bool {
	var shipSectionCount int
	for _, c := range g.board {
		for _, r := range c {
			if r == SHIP {
				shipSectionCount++
			}
		}
	}
	if shipSectionCount == 0 {
		return true
	}
	return false
}
func (g *Game) SetupShips() {
	var ships []*Ship
	ships = append(ships,
		&Ship{
			Type:   CARRIER,
			Length: 5,
			Hits:   make([]int, 5),
		},
		&Ship{
			Type:   BATTLESHIP,
			Length: 4,
			Hits:   make([]int, 4),
		},
		&Ship{
			Type:   CRUISER,
			Length: 3,
			Hits:   make([]int, 3),
		},
		&Ship{
			Type:   SUBMARINE,
			Length: 3,
			Hits:   make([]int, 3),
		},
		&Ship{
			Type:   DESTROYER,
			Length: 2,
			Hits:   make([]int, 2),
		},
	)

	g.player.ParkShips(ships)

	for _, ship := range ships {
		g.board.AddShip(ship)
	}
	g.ships = ships

}
func IsValidShips(ships []*Ship) bool {

	shipLengths := []int{5, 4, 3, 3, 2}

	board := NewBoard()

	for i, ship := range ships {

		if ship.Length != shipLengths[i] {
			return false
		}

		if ship.Start.X > ship.End.X || ship.Start.Y > ship.End.Y {
			return false
		}

		if ship.Start.X < 0 || ship.Start.Y < 0 {
			return false
		}
		if ship.End.X >= 10 || ship.End.Y >= 10 {
			return false
		}

		board.AddShip(ship)
	}

	var shipSectionCount int
	//sum of length of all ships should be 17
	for _, c := range board {
		for _, r := range c {
			if r == SHIP {
				shipSectionCount++
			}
		}
	}
	if shipSectionCount != 17 {
		return false
	}

	return true
}
