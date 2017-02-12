package main

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type Tile struct{
	x int
	y int
	value int
	isEmpty bool
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

func initDraw() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	gameWidth := 100
	gameHeight := 50
	const ROW_SIZE = 4
	const COLUMN_SIZE = 4

	cellWidth := gameWidth / COLUMN_SIZE
	cellHeight := gameHeight / ROW_SIZE
	for ly := 0; ly < ROW_SIZE; ly++ {
		for lx := 0; lx < COLUMN_SIZE; lx++ {
			drawSell(lx * cellWidth, ly * cellHeight, cellWidth, cellHeight)
		}
	}

	termbox.Flush()
}

func drawSell(left int, top int , cellWidth int, cellHeight int){
	fill(left, top, cellWidth, 1, termbox.Cell{Ch: '─'})
	fill(left, top, 1, cellHeight, termbox.Cell{Ch: '|'})
	fill(left, top + cellHeight, cellWidth, 1, termbox.Cell{Ch: '─'})
	fill(left + cellWidth, top, 1, cellHeight, termbox.Cell{Ch: '|'})
}

func drawMessage(msg string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, msg)

	termbox.Flush()
}

func handleKeyEvent() {
	initDraw()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowDown:
				drawMessage("you pushed down")
				break
			case termbox.KeyArrowLeft:
				drawMessage("you pushed left")
				break
			case termbox.KeyArrowRight:
				drawMessage("you pushed right")
				break
			case termbox.KeyArrowUp:
				drawMessage("you pushed up")
				break
			default:
				initDraw()
			}
		default:
			initDraw()
		}
	}
}

//ランダムな初期値が欲しい
//できたらデータとViewを繋げる
//できたらTileあたりを肉付け
//できたらGameの肉付け
func main() {
	//err := termbox.Init()
	////Error
	//if err != nil {
	//	panic(err)
	//}

	gameState := Game{gridSize: 4}
	gameState.setup()

	fmt.Print("OK\n")
	fmt.Print(gameState.grid, "\n")
	fmt.Print("Confirm")


	//defer termbox.Close()

	//handleKeyEvent()
}
