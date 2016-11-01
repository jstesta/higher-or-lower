package cards

var Hearts = Suit{1, "Hearts", "♥"}
var Spades = Suit{2, "Spades", "♠"}
var Clubs = Suit{3, "Clubs", "♣"}
var Diamonds = Suit{4, "Diamonds", "♦"}

type Suit struct {
	id     int
	Name   string
	Symbol string
}
