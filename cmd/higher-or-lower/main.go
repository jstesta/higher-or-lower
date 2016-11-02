package main

import (
	//"flag"
	"fmt"
	"github.com/jstesta/higher-or-lower/cards"
	"github.com/jstesta/higher-or-lower/game"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "higher-or-lower ", log.LUTC|log.LstdFlags)
	logger.Print("begin")

	game := game.NewGame()

	// TODO input for cards

	logger.Println(fmt.Sprintf("%s", game))
	game.Draw(cards.AceOfClubs)
	game.Draw(cards.FiveOfDiamonds)
	logger.Println(fmt.Sprintf("%s", game))

	logger.Print("end")
}
