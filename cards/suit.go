package cards

import (
	"errors"
	"strings"
)

var Hearts = Suit{1, "Hearts", "♥", "H"}
var Spades = Suit{2, "Spades", "♠", "S"}
var Clubs = Suit{3, "Clubs", "♣", "C"}
var Diamonds = Suit{4, "Diamonds", "♦", "D"}

var suitMap = map[string]*Suit{
	Hearts.ShortName:   &Hearts,
	Spades.ShortName:   &Spades,
	Clubs.ShortName:    &Clubs,
	Diamonds.ShortName: &Diamonds,
}

type Suit struct {
	id        int
	Name      string
	Symbol    string
	ShortName string
}

func (s Suit) String() string {
	return s.Name
}

func ParseSuit(input string) (*Suit, error) {

	if len(input) != 1 {
		return nil, errors.New("unable to parse Suit input string")
	}

	input = strings.ToUpper(input)

	if suitMap[input] != nil {
		return suitMap[input], nil
	}

	return nil, errors.New("unable to parse Suit input string")
}
