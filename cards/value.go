package cards

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

type Value struct {
	name      string
	faceValue int
	shortName string
}

func (v Value) String() string {
	return v.name
}