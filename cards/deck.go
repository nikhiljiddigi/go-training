package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

//create a new type of deck whih is a slice of strings
type deck []string

// newDeck functions returns a slice of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

//print method will log or print the list of cards in the given deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal method will pick the hand size of cards in the given deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//toString method is used convert slice of string to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//saveToFile method is used to save the deck of cards to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

//newDeckFromFile method reads a deck of card from a file and return the deck
func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return deck(strings.Split(string(content), ","))
}

// shuffle method is used to shuffle the deck of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
