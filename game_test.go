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

func TestGameSetup(t *testing.T) {
	game := Game{gridSize: 4}
	preTileState := [][][]int{[][]int{[]int{0, 0, 3}, []int{0, 1, 2}}}
	gInfo := GameInfo{HighScore: 10, CurrentScore: 5, TileState: preTileState}
	game.setup(gInfo)

	if game.score != 5 {
		t.Errorf("score %v expected %v", game.score, 5)
	}

	if game.highScore != 10 {
		t.Errorf("high score %v expected %v", game.score, 5)
	}

	for x := 0; x < len(game.grid.cells); x++ {
		for y := 0; y < len(game.grid.cells[x]); y++ {
			if game.grid.cells[x][y].value != preTileState[x][y][2] {
				t.Errorf("the value of game cells in %v column %v row is wrong", game.score, 5)
			}
		}
	}
}

func TestGameSetupNotOvertakePre(t *testing.T) {
	game := Game{gridSize: 4}
	gInfo := GameInfo{}
	game.setup(gInfo)

	if game.score != 0{
		t.Errorf("score %v expected %v", game.score, 0)
	}

	if game.highScore != 0 {
		t.Errorf("high score %v expected %v", game.score, 0)
	}

	notEmptyCount := 0

	for x := 0; x < len(game.grid.cells); x++ {
		for y := 0; y < len(game.grid.cells[x]); y++ {
			if game.grid.cells[x][y].isEmpty {
				if game.grid.cells[x][y].value != 0 {
					t.Errorf("when cells is empty, %v should be equal to zero", game.grid.cells[x][y].value)
				}
			} else {
				if game.grid.cells[x][y].value == 0 {
					t.Errorf("when cells is empty, %v mustn't be equals to zero", game.grid.cells[x][y].value)
				}

				notEmptyCount++
			}
		}
	}

	if notEmptyCount != 2 {
		t.Errorf("Not empty should be %v", 2)
	}
}
