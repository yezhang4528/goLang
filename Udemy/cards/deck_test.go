package main

import (
    "testing"
    "os"
)


func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) != 16 {
        t.Errorf("Expected deck of length 16, got %v", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("Expected first card to be Ace of Spaces, got %v", d[0])
    }

    if d[len(d) - 1] != "Four of Club" {
        t.Errorf("Expected first card to be Four of Club, got %v", d[len(d) - 1])
    }
}

func TestsaveToFileAndgetDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

    deck := newDeck()
    deck.saveToFile("_decktesting")

    loadedDeck := getDeckFromFile("_decktesting")

    if len(loadedDeck) != 16 {
        t.Errorf("Expected loadedDeck length to be 16, got %v", len(loadedDeck))
    }

    os.Remove("_decktesting")
}
