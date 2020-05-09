package battleship

import (
	"math/rand"
	"time"
)

type Player struct {
	Name string
}

//ParkShips will park the ship on board
func (p *Player) ParkShips(ships []*Ship) {
	rand.Seed(time.Now().Unix())
	//Randomly placing the ships
	for {

		for _, ship := range ships {
			rnd := rand.Int31()
			// fmt.Println(rnd)
			switch rnd % 2 {
			case 1:
				ship.Start.X = rand.Intn(10)
				ship.Start.Y = rand.Intn(10 - int(ship.Length))
				ship.End.X = ship.Start.X
				ship.End.Y = ship.Start.Y + int(ship.Length) - 1
			case 0:
				ship.Start.Y = rand.Intn(10)
				ship.Start.X = rand.Intn(10 - int(ship.Length))
				ship.End.Y = ship.Start.Y
				ship.End.X = ship.Start.X + int(ship.Length) - 1
			}
		}
		//try untill all ships are placed at valid position
		if IsValidShips(ships) {
			break
		}
	}

}
