package main

import (
	//"flag"
	"github.com/jstesta/higher-or-lower/cards"
	"log"
	"os"
	"github.com/jstesta/higher-or-lower/game"
	"fmt"
)

func main() {
	logger := log.New(os.Stdout, "higher-or-lower ", log.LUTC|log.LstdFlags)
	logger.Print("begin")

	game := game.NewGame()

	logger.Println(fmt.Sprintf("%s", game))
	game.Use(cards.AceOfClubs)
	game.Use(cards.FiveOfDiamonds)
	logger.Println(fmt.Sprintf("%s", game))

	logger.Print("end")
}
