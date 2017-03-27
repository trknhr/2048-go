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

func (g *Grid) eachCell(callback func(x int, y int, tile Tile)){
	for x := 0; x < g.size; x++{
		for y:= 0; y < g.size; y++{
			callback(x, y, g.cells[x][y])
		}
	}
}

func (g *Grid) CellAvailable() bool{
	isEmpty := false
	g.eachCell(func(x int, y int, tile Tile){
		isEmpty = tile.isEmpty
	})

	return isEmpty
}

func (g * Grid) WithinBounds(position *Tile) bool{
	return position.x >= 0 && position.x < g.size &&
	position.y >= 0 && position.y < g.size
}

//ここから実装。。。
func (g *Grid) AvailableCells() [][]Tile{
	cells := make([][]Tile, g.size)

	g.eachCell(func(x int, y int, tile Tile){
		cells[x][y] = tile
	})

	return cells
}

func (g *Grid) removeTile(tile *Tile){
	g.cells[tile.x][tile.y].isEmpty = true
}

func (g *Grid)CellContent(cell *Tile) *Tile {
	if(g.WithinBounds(cell)){
		return &g.cells[cell.x][cell.y]
	} else {
		return nil
	}

}
