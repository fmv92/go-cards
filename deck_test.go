package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckLength := len(d)
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[deckLength-1] != "King of Clubs" {
		t.Errorf("Expected King of Diamonds, but got %v", d[deckLength-1])
	}
}
