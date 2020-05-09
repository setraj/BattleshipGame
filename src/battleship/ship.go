package battleship

const (
	CARRIER = iota
	BATTLESHIP
	CRUISER
	SUBMARINE
	DESTROYER
)

type Ship struct {
	Type   int
	Sunk   bool
	Start  Cell
	End    Cell
	Hits   []int
	Length int
}

func (s *Ship) HasSunk() bool {
	return s.Sunk
}

func (s *Ship) MarkSunk() {
	s.Sunk = true
}

func (s *Ship) GetLength() int {
	return s.Length
}

func (s *Ship) GetStart() Cell {
	return s.Start
}
func (s *Ship) GetEnd() Cell {
	return s.End
}
func (s *Ship) GetType() int {
	return s.Type
}
