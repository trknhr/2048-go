package main

import (
	"testing"
)

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

	if game.score != 0 {
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

func TestGetVector(t *testing.T) {
	var game Game
	vec := game.getVector(0)
	if vec.x != 0 || vec.y != -1 {
		t.Error("Vector to up value is expected x = %v, y = %v", 0, -1)
	}

	vec = game.getVector(1)
	if vec.x != 1 || vec.y != 0 {
		t.Error("Vector to right value is expected x = %v, y = %v", 0, -1)
	}

	vec = game.getVector(2)
	if vec.x != 0 || vec.y != 1 {
		t.Error("Vector to down value is expected x = %v, y = %v", 0, -1)
	}

	vec = game.getVector(3)
	if vec.x != -1 || vec.y != 0 {
		t.Error("Vector to left value is expected x = %v, y = %v", 0, -1)
	}
}

func TestMoveTile(t *testing.T) {
	pos := [][]int{[]int{0,0}, []int{0, 1}}
	game := Game{gridSize: 4}
	game.setup(GameInfo{TileState: createTileState(4, pos) })

	game.moveTile(&Tile{x: 0, y: 0, value: 4}, &Tile{x: 3, y: 3})

	if game.grid.cells[3][3].value != 4 {
		t.Error("Moved cell should be equal to 4");
	}
}

func createTileState(gridSize int, pos [][]int) [][][]int{
	p := [][][]int{}

	for ly := 0; ly < gridSize; ly++ {
		r := [][]int{}
		for lx := 0; lx < gridSize; lx++ {
			var value int = 0
			for pi := 0; pi < len(pos); pi++ {
				if pos[pi][0] == lx && pos[pi][1] == ly {
					value = 4
				}
			}

			r = append(r, []int{lx, ly, value})
		}
		p = append(p, r)
	}

	return p
}
