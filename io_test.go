package main

import (
	"os"
	"testing"
)

func TestIoOperations(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	d := newDeck()

	err := ioOperation.save(filename, d)
	if err != nil {
		t.Errorf("Could not save the file %v", err)
	}

	loadedDeck := ioOperation.load(filename)
	loadedDeck = deck(loadedDeck)

	os.Remove("card_files/" + filename)

	if loadedDeck[0] != d[0] {
		t.Errorf("Testing normal save and load, expected value %v1, got %v2", d[0], loadedDeck[0])
	}
}
