package main

import (
	"math/rand"
	"time"
)

type Grid struct{
	size int
	cells [][]Tile
}

func (g *Grid) setup(){
	g.cells = make([][]Tile, g.size)
	for lx := 0; lx < g.size; lx++{
		g.cells[lx] = make([]Tile, g.size)

		for ly := 0; ly < g.size; ly++{
			g.cells[lx][ly] = Tile{x: lx, y: ly, isEmpty: true}
		}
	}
}

func (g *Grid)randomAvailableCell() Tile{
	cells := g.cells

	rand.Seed(time.Now().UnixNano())
	if len(cells) > 0 {
		return cells[rand.Int31n(int32(g.size))][rand.Int31n(int32(g.size))]
	} else {
		return Tile{isEmpty: true}
	}
}

func (g *Grid)insertTile(tile Tile){
	g.cells[tile.x][tile.y] = tile
}

func (g *Grid)cellsAvailable() bool{
	isAvailable := false
	for _, row := range g.cells {
		for _, item := range row {
			if item.isEmpty {
				isAvailable = item.isEmpty
			}
		}
	}

	return isAvailable
}
