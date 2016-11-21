package ui

import "github.com/andlabs/ui"

func Start() {
	err := ui.Main(func() {
		window := ui.NewWindow("Higher-or-Lower Helper", 800, 400, false)
		window.SetMargined(true)
		window.SetChild(makeMainBox())
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func makeMainBox() *ui.Box {

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(makeWelcomeGroup(box), false)
	box.Append(ui.NewLabel("Press \"New Game\" to start!"), false)
	return box
}

func makeWelcomeGroup(parent *ui.Box) *ui.Group {

	newGameButton := ui.NewButton("New Game")
	newGameButton.OnClicked(func(b *ui.Button) {
		parent.Delete(1) // deletes existing game or initial label
		newGameBox := initGame()
		parent.Append(newGameBox, false)
	})

	quitButton := ui.NewButton("Quit")
	quitButton.OnClicked(func(b *ui.Button) {
		ui.Quit()
	})

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	box.Append(newGameButton, true)
	box.Append(quitButton, true)

	group := ui.NewGroup("Welcome")
	group.SetChild(box)
	return group
}
