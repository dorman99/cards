package main

import (
	"os"
	"testing"
)

// check the len of decks

func TestNewDeck(t *testing.T) {
	d := newdeck()
	startingCard := "ACE of SPADES"
	lastCard := "KING of CLUB"
	deckLen := 52
	if len(d) != deckLen {
		t.Errorf("ERROR SIZE, Expect %d, Got : %v", deckLen, len(d))
	}

	if d[0] != startingCard {
		t.Errorf("The Starting Card Must be %s not %s", startingCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("The Last Card Not Equal to %v, instead: %v", lastCard, d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newdeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Len of loaded deck not equal to 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
