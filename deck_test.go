package main

import (
	"os"
	"testing"
)

func TestNewCard(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Not correct length, expected 16 got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
