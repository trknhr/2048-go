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

var up Vector = Vector{x: 0, y: -1}
var right Vector = Vector{x: 1, y: 0}
var down Vector = Vector{x: 0, y: 1}
var left Vector = Vector{x: -1, y: 0}

func TestBuildTraversals(t *testing.T){
	game := Game{gridSize: 4}

	var expected PositionTraversal

	//up
	expected = PositionTraversal{x: []int{0, 1, 2, 3}, y: []int{0, 1, 2, 3}}
	checkPositionTraversal(game.BuildTraversals(up), expected, t)

	//left
	expected = PositionTraversal{x: []int{0, 1, 2, 3}, y: []int{0, 1, 2, 3}}
	checkPositionTraversal(game.BuildTraversals(left), expected, t)

	//down
	expected = PositionTraversal{x: []int{0, 1, 2, 3}, y: []int{3, 2, 1, 0}}
	checkPositionTraversal(game.BuildTraversals(down), expected, t)

	//right
	expected = PositionTraversal{x: []int{3, 2, 1, 0}, y: []int{0, 1, 2, 3}}
	checkPositionTraversal(game.BuildTraversals(right), expected, t)
}

func TestFindFarthestPosition(t *testing.T){
	game := Game{gridSize: 4}
	game.setup(GameInfo{TileState: createTileState(4, [][]int{[]int{3, 3}, []int{0, 3}}) })

	actualNext, actualFar := game.FindFarthestPosition(Tile{x: 0, y: 0}, down)

	if actualNext.x != 0  || actualNext.y != 3 {
		t.Errorf("Next cell should be 0, 3", actualNext)
	}

	if actualFar.x != 0  || actualFar.y != 4 {
		t.Errorf("Next cell should be 0, 4", actualFar)
	}

	actualNext, actualFar = game.FindFarthestPosition(Tile{x: 0, y: 0}, left)

	if actualNext.x != 0  || actualNext.y != 0 {
		t.Errorf("Next cell should be 0, 0", actualNext)
	}

	if actualFar.x != -1  || actualFar.y != 0 {
		t.Errorf("Next cell should be -1, 0", actualFar)
	}

	actualNext, actualFar = game.FindFarthestPosition(Tile{x: 0, y: 0}, right)
	if actualNext.x != 2  || actualNext.y != 0 {
		t.Errorf("Next cell should be 2, 0", actualNext)
	}

	if actualFar.x != 3 || actualFar.y != 0 {
		t.Errorf("Next cell should be 3, 0", actualFar)
	}
}

func TestPositionsEqual(t *testing.T){
	game := Game{gridSize: 4}

	if game.positionsEqual(&Tile{x: 1, y: 1}, &Tile{x: 1, y: 1}) != true {
		t.Error("positionsEqual should be true")
	}

	if game.positionsEqual(&Tile{x: 0, y: 1}, &Tile{x: 1, y: 1}) != false {
		t.Error("positionsEqual should be false")
	}
}

func TestTileMatchesAvailable(t *testing.T){
	game := Game{gridSize: 4}
	gameInfo := GameInfo{TileState: createTileState(4, [][]int{[]int{3, 3}, []int{0, 3}}) }
	game.setup(gameInfo)

	if game.tileMatchesAvailable() != false {
		t.Error("The game still has empty cells. So it should return false")
	}
	gameInfo = GameInfo{TileState: createFullTileState(4, []int{2, 2}) }
	game.setup(gameInfo)
	if game.tileMatchesAvailable() != true {
		t.Error("The game still has empty cells. So it should return true")
	}

	gameInfo = GameInfo{
			TileState: createFullTileState(4,
			[]int{
				2, 4, 2, 4,
				4, 2, 4, 2,
				2, 4, 2, 4,
				4, 2, 4, 2,
		})}
	game.setup(gameInfo)
	if game.tileMatchesAvailable() != false {
		t.Error("The game still has empty cells. So it should return true")
	}
}

func TestMovesAvailable(t *testing.T){
	game := Game{gridSize: 4}
	gameInfo := GameInfo{TileState: createTileState(4, [][]int{[]int{3, 3}, []int{0, 3}}) }
	game.setup(gameInfo)

	if game.movesAvailable() != true {
		t.Error("The game still has empty cells. So it should return false")
	}

	gameInfo = GameInfo{TileState: createFullTileState(4, []int{2, 2}) }
	game.setup(gameInfo)
	if game.movesAvailable() != true {
		t.Error("The game still has empty cells. So it should return true")
	}

	gameInfo = GameInfo{
		TileState: createFullTileState(4,
			[]int{
				2, 4, 2, 4,
				4, 2, 4, 2,
				2, 4, 2, 4,
				4, 2, 4, 2,
			})}
	game.setup(gameInfo)
	if game.movesAvailable() != false {
		t.Error("The game still has empty cells. So it should return true")
	}
}

const moveUp = 0
const moveRight = 1
const moveDown = 2
const moveLeft = 3

func TestUpMoveItDoesntMerge(t *testing.T){
	checkMoveTileDosentMerge(moveUp, [][]int{[]int{0, 0, 2}, []int{3, 3, 4}}, []Tile{Tile{x: 0, y: 0, value:2}, Tile{x: 3, y: 0, value: 4}}, t)
}

func TestRightMoveItDoesntMerge(t *testing.T){
	checkMoveTileDosentMerge(moveRight, [][]int{[]int{0, 0, 2}, []int{3, 3, 4}}, []Tile{Tile{x: 3, y: 0, value:2}, Tile{x: 3, y: 3, value:4}}, t)
}

func TestDownMoveItDoesntMerge(t *testing.T){
	checkMoveTileDosentMerge(moveDown, [][]int{[]int{0, 0, 2}, []int{3, 3, 4}}, []Tile{Tile{x: 0, y: 3, value:2}, Tile{x: 3, y: 3, value:4}}, t)
}

func TestLeftMoveItDoesntMerge(t *testing.T){
	checkMoveTileDosentMerge(moveLeft, [][]int{[]int{0, 0, 2}, []int{3, 3, 4}}, []Tile{Tile{x: 0, y: 0, value:2}, Tile{x: 0, y: 3, value:4}}, t)
}

func TestMoveItMergeCell(t *testing.T){
	game := Game{gridSize: 4}
	gameInfo := GameInfo{TileState: createTileState(4, [][]int{[]int{0, 0, 4}, []int{0, 3, 4}})}

	game.setup(gameInfo)
	game.move(moveDown)

	expect := Tile{x: 0, y: 3, value: 8}
	cell := game.grid.CellContent(&expect)

	if cell.isEmpty != false && cell.value != expect.value && len(expect.mergedFrom) != 2{
		t.Errorf("%v, %v should be full and it has %v, mergeFrom is %v", expect.x, expect.y, expect.value, expect.mergedFrom)
	}
}

func checkMoveTileDosentMerge(move int, d [][]int, expect []Tile, t *testing.T){
	game := Game{gridSize: 4}
	var gameInfo GameInfo
	gameInfo.load()
	gameInfo = GameInfo{TileState: createTileState(4, d)}

	game.setup(gameInfo)
	game.move(move)

	for i := 0; i < len(expect); i++ {
		cell := game.grid.CellContent(&expect[i])

		if cell.isEmpty != false && cell.value != expect[i].value{
			t.Errorf("%v, %v should be full and it has %v", expect[i].x, expect[i].y, expect[i].value)
		}
	}
}

func checkPositionTraversal(actual PositionTraversal, expected PositionTraversal, t *testing.T){
	for i := 0; i < len(actual.x); i++ {
		if actual.x[i] != expected.x[i] {
			t.Errorf("The actual posiion x %v should be equal to %v at index %v", actual.x[i], expected.x[i], i)
		}
	}

	for i := 0; i < len(actual.y); i++ {
		if actual.y[i] != expected.y[i] {
			t.Errorf("The actual position y %v should be equal to %v at index %v", actual.y[i], expected.y[i], i)
		}
	}
}

func createFullTileState(gridSize int, values []int)[][][]int{
	ts := createTileState(gridSize, [][]int{})

	for ly := 0; ly < gridSize; ly++ {
		for lx := 0; lx < gridSize; lx++ {
			if ly * gridSize + lx < len(values) {
				ts[ly][lx][2] = values[ly * gridSize + lx]
			} else {
				ts[ly][lx][2] = 2
			}
		}
	}

	return ts
}

func createTileState(gridSize int, pos [][]int) [][][]int{
	p := [][][]int{}

	for ly := 0; ly < gridSize; ly++ {
		r := [][]int{}
		for lx := 0; lx < gridSize; lx++ {
			var value int = 0
			for pi := 0; pi < len(pos); pi++ {
				if pos[pi][0] == lx && pos[pi][1] == ly {
					if len(pos[pi]) > 2 {
						value = pos[pi][2]
					} else {
						value = 4
					}
				}
			}

			r = append(r, []int{ly, lx, value})
		}
		p = append(p, r)
	}

	return p
}
