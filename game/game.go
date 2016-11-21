package game

import (
	"fmt"
	"github.com/jstesta/higher-or-lower/cards"
)

type Game struct {
	available        cards.Deck
	used             cards.Deck
	lastDrawnCard    *cards.Card
	remainingLess    cards.Deck
	remainingEqual   cards.Deck
	remainingGreater cards.Deck
}

func NewGame() *Game {

	a := cards.NewDeck()
	u := cards.EmptyDeck()
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

	g.remainingLess, g.remainingEqual, g.remainingGreater = g.available.Split(c)
}

func (g *Game) String() string {
	return fmt.Sprintf("Available=%s, Used=%s", g.available, g.used)
}

func (g *Game) Chances() (float32, float32, float32) {

	total := len(g.available)

	if total == 0 {
		return 0, 0, 0
	}

	totalf := float32(total)

	less := float32(len(g.remainingLess))
	equal := float32(len(g.remainingEqual))
	greater := float32(len(g.remainingGreater))
	return 100 * less / totalf, 100 * equal / totalf, 100 * greater / totalf
}
