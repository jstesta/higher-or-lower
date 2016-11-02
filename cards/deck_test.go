package cards

import "testing"

func TestDeck_Contains(t *testing.T) {

	deck := NewDeck()
	card := AceOfClubs

	if !deck.Contains(card) {
		t.Error("basic Deck Contains check failed")
	}
}

func TestDeck_Remove(t *testing.T) {

	deck := NewDeck()
	card := AceOfClubs

	deckWithoutCard := deck.Remove(card)

	for _, c := range deckWithoutCard {
		if card == c {
			t.Error("basic Deck Remove check failed")
		}
	}
}

func TestDeck_NotContainsRemovedCard(t *testing.T) {

	deck := NewDeck()
	card := AceOfClubs

	deckWithoutCard := deck.Remove(card)

	if deckWithoutCard.Contains(card) {
		t.Error("Deck remove, then Contains check failed")
	}
}

func TestDeck_GetBySuit(t *testing.T) {

	deck := NewDeck()
	suit := Hearts

	onlyHearts := deck.GetBySuit(suit)

	if len(onlyHearts) != 13 {
		t.Error("GetBySuit check failed: not 13 cards")
	}

	remainingSuits := []Suit{Clubs, Spades, Diamonds}

	for _, c := range onlyHearts {
		for _, s := range remainingSuits {
			if c.suit == s {
				t.Error("GetBySuit check failed: invalid Suit found")
			}
		}
	}
}

func TestDeck_GetByValue(t *testing.T) {

	deck := NewDeck()
	value := Three

	onlyThrees := deck.GetByValue(value)

	if len(onlyThrees) != 4 {
		t.Error("GetByValue check failed: not 4 cards")
	}

	remainingValues := []Value{Two, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, c := range onlyThrees {
		for _, s := range remainingValues {
			if c.value.faceValue == s.faceValue {
				t.Error("GetByValue check failed: invalid Value found")
			}
		}
	}
}

func TestDeck_Add(t *testing.T) {

	deck := make(Deck, 0)
	card := AceOfDiamonds

	deck = deck.Add(card)

	if len(deck) != 1 {
		t.Error("Add check failed: deck does not have 1 card")
	}

	if !deck.Contains(card) {
		t.Error("Add check failed: deck does not contain the added card")
	}
}
