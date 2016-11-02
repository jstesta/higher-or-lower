package game

import (
	"fmt"
	"github.com/jstesta/higher-or-lower/cards"
)

type Game struct {
	available     cards.Deck
	used          cards.Deck
	lastDrawnCard *cards.Card
}

func NewGame() *Game {
	a := cards.NewDeck()
	u := make(cards.Deck, 0)
	return &Game{available: a, used: u}
}

func (g *Game) AvailableCards() cards.Deck {
	return g.available
}

func (g *Game) UsedCards() cards.Deck {
	return g.used
}

func (g *Game) LastDrawnCard() *cards.Card {
	return g.lastDrawnCard
}

func (g *Game) Draw(c cards.Card) {
	g.available = g.available.Remove(c)
	g.used = g.used.Add(c)
	g.lastDrawnCard = &c
}

func (g *Game) String() string {
	return fmt.Sprintf("Available=%s, Used=%s", g.available, g.used)
}
