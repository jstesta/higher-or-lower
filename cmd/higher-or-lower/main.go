package main

import (
	//"flag"
	"github.com/jstesta/higher-or-lower/cards"
	"log"
	"os"
)

func main() {
	//var (
	//	midiFile = flag.String("input", "", "Filesystem location of MIDI file to parse")
	//)
	//flag.Parse()

	//f, err := os.Create("higher-or-lower.log")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//logger := log.New(f, "higher-or-lower ", log.LUTC|log.LstdFlags)
	logger := log.New(os.Stdout, "higher-or-lower ", log.LUTC|log.LstdFlags)
	logger.Print("begin")

	deck := cards.NewDeck()

	for _, c := range deck {
		logger.Println(c.Name())
		logger.Println(c.ShortName())
	}

	logger.Print("end")
}
