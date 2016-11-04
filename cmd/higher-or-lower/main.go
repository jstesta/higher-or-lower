package main

import (
	"bufio"
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter 'q' to quit")
	for {
		fmt.Print("Enter card: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1] // trim off trailing newline

		if text == "q" {
			fmt.Println("Quit")
			break
		}

		card, err := cards.ParseCard(text)
		if err != nil {
			fmt.Println("Bad input: " + err.Error())
			continue
		}
		fmt.Println("Input Card: " + card.String())
		game.Draw(*card)

		chanceGreater, chanceEqual, chanceLess := game.Chances()
		fmt.Println(fmt.Sprintf(
			"Less: %2.2f%%, Equal: %2.2f%%, Greater: %2.2f%%",
			chanceLess, chanceEqual, chanceGreater))
	}

	logger.Print("end")
}
