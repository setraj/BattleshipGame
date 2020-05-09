package battleship

import (
	"fmt"
)

const (
	PLAIN = iota
	HIT
	MISS
	SHIP
)

type Cell struct {
	X int
	Y int
}

type Board [][]int

func NewBoard() (b Board) {
	b = make([][]int, 10)
	for i := range b {
		b[i] = make([]int, 10)
	}
	return
}

func (b Board) AddShip(s *Ship) {
	if s.Start.X == s.End.X {
		for i := s.Start.Y; i <= s.End.Y; i++ {
			b[s.Start.X][i] = SHIP
		}
	} else {
		for i := s.Start.X; i <= s.End.X; i++ {
			b[i][s.Start.Y] = SHIP
		}
	}
}

func (b Board) Draw() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v ", b[i][j])
		}
		fmt.Println("\n")
	}
	fmt.Println("===========================================")
}
