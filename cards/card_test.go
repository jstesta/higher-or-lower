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
