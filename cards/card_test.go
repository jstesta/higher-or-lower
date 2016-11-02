package cards

import "testing"

func TestBasicCardEquality(t *testing.T) {

	card := AceOfClubs

	if 0 != card.CompareTo(AceOfClubs) {
		t.Error("basic Card equality check failed")
	}
}

func TestCompareToLower(t *testing.T) {

	lowerCard := TwoOfClubs
	higherCard := FiveOfSpades

	if -1 != lowerCard.CompareTo(higherCard) {
		t.Error("basic Card less than check failed")
	}
}

func TestCompareToGreater(t *testing.T) {

	lowerCard := TwoOfClubs
	higherCard := FiveOfSpades

	if 1 != higherCard.CompareTo(lowerCard) {
		t.Error("basic Card greater than check failed")
	}
}

func TestParseCard_2(t *testing.T) {

	input := "AH" // Ace of Hearts

	parsed, err := ParseCard(input)
	if err != nil {
		t.Error("ParseCard error: " + err.Error())
	}

	if *parsed != AceOfHearts {
		t.Error("ParseCard returned wrong Card")
	}
}

func TestParseCard_3(t *testing.T) {

	input := "10D" // Ten of Diamonds

	parsed, err := ParseCard(input)
	if err != nil {
		t.Error("ParseCard Card error: " + err.Error())
	}

	if *parsed != TenOfDiamonds {
		t.Error("ParseCard returned wrong Card")
	}
}

func TestParseCard_All(t *testing.T) {

	deck := NewDeck()

	for _, card := range deck {
		var input string
		input += card.value.shortName
		input += card.suit.ShortName

		parsed, err := ParseCard(input)
		if err != nil {
			t.Error("ParseCard Card error: " + err.Error())
		}

		if *parsed != card {
			t.Error("ParseCard returned wrong Card: " + card.String())
		}
	}
}
