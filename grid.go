package main

type Grid struct{
	size int
	cells [][]Tile
}

func (g *Grid) setup(){
	g.cells = make([][]Tile, g.size)
	for lx := 0; lx < g.size; lx++{
		g.cells[lx] = make([]Tile, g.size)

		for ly := 0; ly < g.size; ly++{
			g.cells[lx][ly] = Tile{isEmpty: true}
		}
	}
}

