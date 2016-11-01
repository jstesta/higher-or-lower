package cards

var Two = Value{"Two", 2}
var Three = Value{"Three", 3}
var Four = Value{"Four", 4}
var Five = Value{"Five", 5}
var Six = Value{"Six", 6}
var Seven = Value{"Seven", 7}
var Eight = Value{"Eight", 8}
var Nine = Value{"Nine", 9}
var Ten = Value{"Ten", 10}
var Jack = Value{"Jack", 11}
var Queen = Value{"Queen", 12}
var King = Value{"King", 13}
var Ace = Value{"Ace", 14}

type Value struct {
	name      string
	faceValue int
}
