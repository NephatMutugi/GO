package main

import "fmt"

func main() {

	// slice declaration
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of spades")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
