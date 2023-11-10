package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Receivers by convention are one or two letters that match the type of the receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	// Convert deck to slice of strings
	// Join the slice of strings into a single string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// Convert deck to slice of strings
	byteslice := []byte(d.toString())

	// Write to file with 0666 permissions (read/write)
	return ioutil.WriteFile(filename, byteslice, fs.FileMode(0666))
}

func newDeckFromFile(filename string) deck {
	// Read from file
	byteslice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// return newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1) // 1 indicates error
	}

	// Convert byte slice to string
	s := strings.Split(string(byteslice), ",")

	// Convert string slice to deck
	return deck(s)
}

func (d deck) shuffle() {
	// Create a new source or seed for the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap the positions of the cards
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
