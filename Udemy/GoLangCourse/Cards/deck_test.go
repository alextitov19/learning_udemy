package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %s", d[1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(ld))
	}

	os.Remove("_decktesting")
}