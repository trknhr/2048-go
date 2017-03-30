package main

import (
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
	"strconv"
	"fmt"
)

type Tile struct{
	x int
	y int
	value int
	isEmpty bool
	mergedFrom []*Tile
}

func (t *Tile) updatePosition(pos *Tile){
	t.x = pos.x
	t.y = pos.y
}

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func drawSell(tile Tile, left int, top int , cellWidth int, cellHeight int){
	const coldef = termbox.ColorDefault

	fill(left, top, cellWidth, 1, termbox.Cell{Ch: '─'})
	fill(left, top, 1, cellHeight, termbox.Cell{Ch: '|'})
	fill(left, top + cellHeight, cellWidth, 1, termbox.Cell{Ch: '─'})
	fill(left + cellWidth, top, 1, cellHeight, termbox.Cell{Ch: '|'})
	if !tile.isEmpty{
		tbprint(left + cellWidth / 2, top + cellHeight / 2, coldef, coldef, strconv.Itoa(tile.value))
	}
}

func drawMessage(msg string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, msg)

	termbox.Flush()
}

func handleKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowDown:
				dispatch("down", &Message{data: "push left"})
				break
			case termbox.KeyArrowLeft:
				dispatch("left", &Message{data: "push left"})
				break
			case termbox.KeyArrowRight:
				dispatch("right", &Message{data: "push left"})
				break
			case termbox.KeyArrowUp:
				dispatch("up", &Message{data: "push left"})
				break
			default:
				//initDraw()
			}
		default:
			//initDraw()
		}
	}
}

type Drawer struct{}

func (d *Drawer)redraw(grid *Grid){
	gridDraw(grid)
}

func gridDraw(grid *Grid){
	//const coldef = termbox.ColorDefault
	//termbox.Clear(coldef, coldef)
	//gameWidth := 100
	//gameHeight := 50
	//
	//cellWidth := gameWidth / grid.size
	//cellHeight := gameHeight / grid.size
	fmt.Println(grid.cells)

	//for ly := 0; ly < grid.size; ly++ {
	//	for lx := 0; lx < grid.size; lx++ {
	//		tile := grid.cells[lx][ly]
	//		drawSell(tile, lx * cellWidth, ly * cellHeight, cellWidth, cellHeight)
	//	}
	//}

	//termbox.Flush()
}

func main() {
	err := termbox.Init()
	//Error
	if err != nil {
		panic(err)
	}

	drawer := Drawer{}
	gameState := Game{gridSize: 4, drawer: &drawer}
	gameState.setup()

	gridDraw(gameState.grid)

	defer termbox.Close()

	handleKeyEvent()
}
