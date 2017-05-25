package main

import "testing"

func TestIsGameTerminated(t *testing.T) {
	game := Game{gridSize: 4}
	actual := game.IsGameTerminated()
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
