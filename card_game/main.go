package main

import "fmt"

func main() {
	// Slice of type string
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	// convert byte slice to string
	fmt.Println(cards.toString())

	// save to file
	cards.saveToFile("my_cards")

	// read from file
	cards = newDeckFromFile("my_cards")
	cards.print()

	// shuffle
	cards.shuffle()
	cards.print()
}
