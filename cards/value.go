package cards

import (
	"errors"
	"fmt"
	"strings"
)

var Two = Value{"Two", 2, "2"}
var Three = Value{"Three", 3, "3"}
var Four = Value{"Four", 4, "4"}
var Five = Value{"Five", 5, "5"}
var Six = Value{"Six", 6, "6"}
var Seven = Value{"Seven", 7, "7"}
var Eight = Value{"Eight", 8, "8"}
var Nine = Value{"Nine", 9, "9"}
var Ten = Value{"Ten", 10, "10"}
var Jack = Value{"Jack", 11, "J"}
var Queen = Value{"Queen", 12, "Q"}
var King = Value{"King", 13, "K"}
var Ace = Value{"Ace", 14, "A"}

var valueMap = map[string]*Value{
	Two.shortName:   &Two,
	Three.shortName: &Three,
	Four.shortName:  &Four,
	Five.shortName:  &Five,
	Six.shortName:   &Six,
	Seven.shortName: &Seven,
	Eight.shortName: &Eight,
	Nine.shortName:  &Nine,
	Ten.shortName:   &Ten,
	Jack.shortName:  &Jack,
	Queen.shortName: &Queen,
	King.shortName:  &King,
	Ace.shortName:   &Ace,
}

type Value struct {
	name      string
	faceValue int
	shortName string
}

func (v Value) String() string {
	return v.name
}

func ParseValue(input string) (*Value, error) {

	if len(input) < 1 || len(input) > 2 {
		return nil, errors.New(
			fmt.Sprintf("unable to parse Value input string: string length '%s'", input))
	}

	upperInput := strings.ToUpper(input)

	if valueMap[upperInput] == nil {
		return nil, errors.New(
			fmt.Sprintf("unable to parse Value input string: unknown value '%s'", input))
	}

	return valueMap[upperInput], nil
}
