package main

import (
	"os"
	"testing"
)

// TestNewDeck will test if we got the expected number of cards and correct cards
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Hearts" {
		t.Errorf("Expected last card of Five of Hearts, but got %v", d[len(d)-1])
	}
}

// TestSaveToDeckAndNewDeckFromFile will test to save the deck into file and load from file
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
