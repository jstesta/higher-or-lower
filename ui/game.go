package ui

import (
	"fmt"
	"github.com/andlabs/ui"
	"github.com/jstesta/higher-or-lower/cards"
	"github.com/jstesta/higher-or-lower/game"
)

func initGame() *ui.Box {

	gameController := &GameController{
		Game:               game.NewGame(),
		PercentHigherLabel: ui.NewLabel("-"),
		PercentDrawLabel:   ui.NewLabel("-"),
		PercentLowerLabel:  ui.NewLabel("-"),
	}

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(makeCardSelectionBox(gameController), false)
	box.Append(makeStatisticsBox(gameController), false)
	return box
}

func makeCardSelectionBox(g *GameController) *ui.Box {

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(makeSuitGroup(g, cards.Spades), false)
	box.Append(makeSuitGroup(g, cards.Clubs), false)
	box.Append(makeSuitGroup(g, cards.Hearts), false)
	box.Append(makeSuitGroup(g, cards.Diamonds), false)
	return box
}

func makeStatisticsBox(g *GameController) *ui.Box {

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	box.Append(g.PercentLowerLabel, true)
	box.Append(g.PercentDrawLabel, true)
	box.Append(g.PercentHigherLabel, true)
	return box
}

func makeCardButton(g *GameController, c cards.Card) *ui.Button {

	button := ui.NewButton(c.ShortName())
	button.OnClicked(func(b *ui.Button) {
		b.Disable()
		g.Game.Draw(c)
		less, equal, greater := g.Game.Chances()
		g.PercentLowerLabel.SetText(fmt.Sprintf("Less: %2.2f%%", less))
		g.PercentDrawLabel.SetText(fmt.Sprintf("Equal: %2.2f%%", equal))
		g.PercentHigherLabel.SetText(fmt.Sprintf("Greater: %2.2f%%", greater))
	})
	return button
}

func makeSuitGroup(g *GameController, s cards.Suit) *ui.Group {

	deck := cards.NewDeck()
	suited := deck.GetBySuit(s)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	for _, c := range suited {
		hbox.Append(makeCardButton(g, c), true)
	}

	group := ui.NewGroup(s.Name)
	group.SetChild(hbox)
	return group
}
