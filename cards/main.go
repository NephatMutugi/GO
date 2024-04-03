package main

import "fmt"

var deckSize int

func main() {
	card := "Ace of Spades" // Only use the (:=) syntax when creating a new variable
	card = "Five of Diamonds"
	deckSize = 50
	fmt.Println(card)
	fmt.Printf("Deck Size: %d\n", deckSize)
}
