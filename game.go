package main

import (
	"math/rand"
)

type Game struct {
	gridSize int
	startTiles int
	score int
	over bool
	won bool
	grid *Grid
}


func (g *Game)setup(){
	g.score = 0
	g.startTiles = 2
	g.grid = &Grid{size: g.gridSize}
	g.grid.setup()
	g.addStartTiles()
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

