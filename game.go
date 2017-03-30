package main

import (
	"math/rand"
	"fmt"
)

type Game struct {
	gridSize int
	startTiles int
	score int
	over bool
	won bool
	grid *Grid
	drawer *Drawer
}

type Vector struct{
	x int
	y int
}

type PositionTraversal struct{
	x []int
	y []int
}

func (g *Game)setup(){
	g.score = 0
	g.startTiles = 2
	g.grid = &Grid{size: g.gridSize}
	g.grid.setup()
	g.addStartTiles()

	fmt.Println("g.gridSize" , g.gridSize)
	add("up", func(message *Message){
		g.move(0)
	})
	add("right", func(message *Message){
		g.move(1)
	})
	add("down", func(message *Message){
		g.move(2)
	})
	add("left", func(message *Message){
		g.move(3)
	})
}

func (g *Game) addStartTiles(){
	for i := 0; i < g.startTiles; i++{
		g.addRandomTile()
	}
}

func (g *Game) addRandomTile(){
	if g.grid.cellsAvailable() {
		value := 2
		if rand.Float32() < 0.9 {
			value = 4
		}
		tile := g.grid.randomAvailableCell()
		newTile := Tile{x: tile.x, y: tile.y, value: value, isEmpty: false}

		g.grid.insertTile(newTile)
	}
}

func (g *Game) GetVector(direction int) Vector{
	res := make(map[int]Vector)

	res[0] = Vector{x: 0, y: -1} // Up
	res[1] = Vector{x: 1, y: 0} // Right
	res[2] = Vector{x: 0, y: 1} // Down
	res[3] = Vector{x: -1, y: 0} // Left

	return res[direction]
}

func (g *Game) moveTile(tile *Tile, cell *Tile){
	g.grid.removeTile(tile)
	g.grid.cells[cell.x][cell.y] = Tile{x: tile.x, y: tile.y, value: tile.value, mergedFrom: tile.mergedFrom}
	tile.updatePosition(cell)
}

func (g *Game) IsGameTerminated() bool{
	return false
}

func (g *Game) BuildTraversals(vec Vector) PositionTraversal{
	traversals := PositionTraversal{x: make([]int, g.gridSize), y: make([]int, g.gridSize)}

	for i := 0; i < g.gridSize; i++ {
		traversals.x[i] = i
		traversals.y[i] = i
	}

	if(vec.x == 1){
		traversals.x = ReverseList(traversals.x)
	}

	if(vec.y == 1){
		traversals.y = ReverseList(traversals.y)
	}

	return traversals
}

func (g *Game) FindFarthestPosition(cell Tile, vector Vector) (*Tile, *Tile){
	var previous = cell

	for g.grid.WithinBounds(&cell) && g.grid.CellAvailable() {
		previous = cell
		//fmt.Println("")
		//fmt.Println("==vector=====", vector)
		//fmt.Println("==cell=====", cell)
		//fmt.Println("==previous=====", previous)
		//fmt.Println("")
		cell = Tile{x: previous.x + vector.x, y: previous.y + vector.y}
	}

	return &previous, &cell
}

func (g *Game) positionsEqual(first *Tile, second *Tile) bool {
	return first.x == second.x && first.y == second.y
}

//Todo, I have to implement later...
func (g *Game) tileMatchesAvailable() bool{
	return true
}

func (g *Game) movesAvailable() bool {
	return g.grid.cellsAvailable() || g.tileMatchesAvailable()

}

func (g *Game) move(direction int){


	if(g.IsGameTerminated()){
		return
	}
	//
	//moved := false
	//fmt.Println("=======================vector===========================", direction)
	vector := g.GetVector(direction)
	//fmt.Println("=======================vector===========================", vector)
	traversals := g.BuildTraversals(vector)
	//fmt.Println("=======================moved===========================", traversals)
	//fmt.Println("==========================g.gridSize==============================", g.gridSize)
	//fmt.Println("moved", moved)

	for i := 0; i < len(traversals.x); i++ {
		for j := 0; j < len(traversals.y); j++{
			cell := Tile{x: traversals.x[i], y: traversals.y[j]}
			tile := g.grid.CellContent(&cell)
			if(!tile.isEmpty){
				farPos, nextPos := g.FindFarthestPosition(cell, vector)
				//fmt.Println("")
				//fmt.Println("=======================tile===========================", tile)
				//fmt.Println("=======================farPos===========================", farPos)
				//fmt.Println("=======================nextPos===========================", nextPos)
				next := g.grid.CellContent(nextPos)
				//fmt.Println("=======================farPos===========================", farPos)
				//fmt.Println("")
				if( next != nil && next.value == tile.value /*&& !next.mergedFrom*/){
	//				merged := Tile{x: nextPos.x, y: nextPos.y, value: tile.value * 2}
	//				tiles := make([]*Tile, 2)
	//				tiles[0] = tile
	//				tiles[1] = nextPos
	//				merged.mergedFrom = tiles//[]Tile{tile, nextPos}
	//
	//				g.grid.insertTile(merged)
	//				g.grid.removeTile(tile)
	//
	//				tile.updatePosition(nextPos)
	//
	//				g.score += merged.value
				} else {
					fmt.Println("=======move======")
					g.moveTile(tile, farPos)
				}
				if(!g.positionsEqual(&cell, tile)){
					//moved = true
				}
			}
		}
	}
	//
	//if(moved){
		//g.addRandomTile()
	//
	//	if(!g.movesAvailable()){
	//		g.over = true
	//	}
	//
		g.actuate()
	//}
}

func (g *Game) actuate(){
	g.drawer.redraw(g.grid)
	//fmt.Println("==========torima=================", g.grid)
}
