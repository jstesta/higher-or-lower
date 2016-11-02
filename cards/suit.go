package cards

var Hearts = Suit{1, "Hearts", "♥", "H"}
var Spades = Suit{2, "Spades", "♠", "S"}
var Clubs = Suit{3, "Clubs", "♣", "C"}
var Diamonds = Suit{4, "Diamonds", "♦", "D"}

type Suit struct {
	id     int
	Name   string
	Symbol string
	ShortName string
}

func (s Suit) String() string {
	return s.Name
}