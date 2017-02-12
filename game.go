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
}

func (g *Game) addStartTiles(){
	for i := 0; i < g.startTiles; i++{
		g.addRandomTile()
	}
}

func (g *Game) addRandomTile(){
	value := rand.Float32()

}
