package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected dec length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Daimonds" {
		t.Errorf("Expected first card ace of diamonds but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card Four of Clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()

	d.saveToFile("_decktesting.txt")

	loadeddeck := newDeckFromFile("_decktesting.txt")

	if len(loadeddeck) != 16 {
		t.Errorf("Expected dec length of 20 but got %v", len(loadeddeck))
	}

	os.Remove("_decktesting.txt")

}
