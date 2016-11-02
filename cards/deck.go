package cards

import "fmt"

type Deck []Card

func NewDeck() Deck {
	d := make(Deck, 0, 52)
	d = append(d, TwoOfHearts)
	d = append(d, ThreeOfHearts)
	d = append(d, FourOfHearts)
	d = append(d, FiveOfHearts)
	d = append(d, SixOfHearts)
	d = append(d, SevenOfHearts)
	d = append(d, EightOfHearts)
	d = append(d, NineOfHearts)
	d = append(d, TenOfHearts)
	d = append(d, JackOfHearts)
	d = append(d, QueenOfHearts)
	d = append(d, KingOfHearts)
	d = append(d, AceOfHearts)
	d = append(d, TwoOfSpades)
	d = append(d, ThreeOfSpades)
	d = append(d, FourOfSpades)
	d = append(d, FiveOfSpades)
	d = append(d, SixOfSpades)
	d = append(d, SevenOfSpades)
	d = append(d, EightOfSpades)
	d = append(d, NineOfSpades)
	d = append(d, TenOfSpades)
	d = append(d, JackOfSpades)
	d = append(d, QueenOfSpades)
	d = append(d, KingOfSpades)
	d = append(d, AceOfSpades)
	d = append(d, TwoOfClubs)
	d = append(d, ThreeOfClubs)
	d = append(d, FourOfClubs)
	d = append(d, FiveOfClubs)
	d = append(d, SixOfClubs)
	d = append(d, SevenOfClubs)
	d = append(d, EightOfClubs)
	d = append(d, NineOfClubs)
	d = append(d, TenOfClubs)
	d = append(d, JackOfClubs)
	d = append(d, QueenOfClubs)
	d = append(d, KingOfClubs)
	d = append(d, AceOfClubs)
	d = append(d, TwoOfDiamonds)
	d = append(d, ThreeOfDiamonds)
	d = append(d, FourOfDiamonds)
	d = append(d, FiveOfDiamonds)
	d = append(d, SixOfDiamonds)
	d = append(d, SevenOfDiamonds)
	d = append(d, EightOfDiamonds)
	d = append(d, NineOfDiamonds)
	d = append(d, TenOfDiamonds)
	d = append(d, JackOfDiamonds)
	d = append(d, QueenOfDiamonds)
	d = append(d, KingOfDiamonds)
	d = append(d, AceOfDiamonds)
	return d
}

func (d Deck) Contains(c Card) bool {
	for _, a := range d {
		if a == c {
			return true
		}
	}
	return false
}

// Returns a slice from this Deck without the specified Card
func (d Deck) Remove(c Card) Deck {
	for idx, a := range d {
		if a == c {
			d[len(d)-1], d[idx] = d[idx], d[len(d)-1]
			return d[:len(d)-1]
		}
	}
	return d
}

func (d Deck) Add(c Card) Deck {
	return append(d, c)
}

// Returns a new Deck with only cards from this Deck of the specified Suit
func (d Deck) GetBySuit(s Suit) Deck {
	bySuit := make(Deck, 0, 13)
	for _, c := range d {
		if c.suit.id == s.id {
			bySuit = append(bySuit, c)
		}
	}
	return bySuit
}

// Returns a new Deck with only cards from this Deck of the specified Value
func (d Deck) GetByValue(v Value) Deck {
	byValue := make(Deck, 0, 4)
	for _, c := range d {
		if c.value.faceValue == v.faceValue {
			byValue = append(byValue, c)
		}
	}
	return byValue
}

func (d Deck) String() (s string) {
	s = "Deck["
	prefixLength := len(s)
	for _, c := range d {
		s += fmt.Sprintf("%s, ", c)
	}
	if len(s) > prefixLength {
		s = s[:len(s)-2]
	}
	s += "]"
	return
}
