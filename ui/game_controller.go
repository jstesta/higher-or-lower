package ui

import (
	"github.com/andlabs/ui"
	"github.com/jstesta/higher-or-lower/game"
)

type GameController struct {
	Game               *game.Game
	PercentHigherLabel *ui.Label
	PercentDrawLabel   *ui.Label
	PercentLowerLabel  *ui.Label
}
