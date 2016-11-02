package cards

import (
	"errors"
	"fmt"
)

var TwoOfHearts = Card{Two, Hearts}
var ThreeOfHearts = Card{Three, Hearts}
var FourOfHearts = Card{Four, Hearts}
var FiveOfHearts = Card{Five, Hearts}
var SixOfHearts = Card{Six, Hearts}
var SevenOfHearts = Card{Seven, Hearts}
var EightOfHearts = Card{Eight, Hearts}
var NineOfHearts = Card{Nine, Hearts}
var TenOfHearts = Card{Ten, Hearts}
var JackOfHearts = Card{Jack, Hearts}
var QueenOfHearts = Card{Queen, Hearts}
var KingOfHearts = Card{King, Hearts}
var AceOfHearts = Card{Ace, Hearts}
var TwoOfSpades = Card{Two, Spades}
var ThreeOfSpades = Card{Three, Spades}
var FourOfSpades = Card{Four, Spades}
var FiveOfSpades = Card{Five, Spades}
var SixOfSpades = Card{Six, Spades}
var SevenOfSpades = Card{Seven, Spades}
var EightOfSpades = Card{Eight, Spades}
var NineOfSpades = Card{Nine, Spades}
var TenOfSpades = Card{Ten, Spades}
var JackOfSpades = Card{Jack, Spades}
var QueenOfSpades = Card{Queen, Spades}
var KingOfSpades = Card{King, Spades}
var AceOfSpades = Card{Ace, Spades}
var TwoOfClubs = Card{Two, Clubs}
var ThreeOfClubs = Card{Three, Clubs}
var FourOfClubs = Card{Four, Clubs}
var FiveOfClubs = Card{Five, Clubs}
var SixOfClubs = Card{Six, Clubs}
var SevenOfClubs = Card{Seven, Clubs}
var EightOfClubs = Card{Eight, Clubs}
var NineOfClubs = Card{Nine, Clubs}
var TenOfClubs = Card{Ten, Clubs}
var JackOfClubs = Card{Jack, Clubs}
var QueenOfClubs = Card{Queen, Clubs}
var KingOfClubs = Card{King, Clubs}
var AceOfClubs = Card{Ace, Clubs}
var TwoOfDiamonds = Card{Two, Diamonds}
var ThreeOfDiamonds = Card{Three, Diamonds}
var FourOfDiamonds = Card{Four, Diamonds}
var FiveOfDiamonds = Card{Five, Diamonds}
var SixOfDiamonds = Card{Six, Diamonds}
var SevenOfDiamonds = Card{Seven, Diamonds}
var EightOfDiamonds = Card{Eight, Diamonds}
var NineOfDiamonds = Card{Nine, Diamonds}
var TenOfDiamonds = Card{Ten, Diamonds}
var JackOfDiamonds = Card{Jack, Diamonds}
var QueenOfDiamonds = Card{Queen, Diamonds}
var KingOfDiamonds = Card{King, Diamonds}
var AceOfDiamonds = Card{Ace, Diamonds}

type Card struct {
	value Value
	suit  Suit
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Name() string {
	return fmt.Sprintf("%s of %s", c.value.name, c.suit.Name)
}

func (c Card) ShortName() string {
	return fmt.Sprintf("%s%s", c.value.shortName, c.suit.Symbol)
}

func (c Card) FaceValue() int {
	return c.value.faceValue
}

func (c Card) CompareTo(o Card) int {

	cValue := c.value.faceValue
	oValue := o.value.faceValue

	if cValue == oValue {
		return 0
	}

	if cValue > oValue {
		return 1
	}

	return -1
}

func (c Card) String() string {
	return c.Name()
}

func ParseCard(input string) (*Card, error) {

	if len(input) < 2 || len(input) > 3 {
		return nil, errors.New("unable to parse card input string")
	}

	var valuePart string
	var suitPart string
	switch len(input) {
	case 2:
		valuePart = input[:1]
		suitPart = input[1:]
		break
	case 3:
		valuePart = input[:2]
		suitPart = input[2:]
		break
	}

	value, err := ParseValue(valuePart)
	if err != nil {
		return nil, err
	}

	suit, err := ParseSuit(suitPart)
	if err != nil {
		return nil, err
	}

	return &Card{value: *value, suit: *suit}, nil
}
