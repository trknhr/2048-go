package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

const GAME_WIDTH = 80
const GAME_HEIGHT = 40

type Tile struct {
	x          int
	y          int
	value      int
	isEmpty    bool
	mergedFrom []Tile
}

func (t *Tile) updatePosition(pos *Tile) {
	t.x = pos.x
	t.y = pos.y
}

func drawMessage(x, y int, str string, color termbox.Attribute){
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}
func drawLine(x, y int, str string) {
	drawMessage(x, y, str, termbox.ColorDefault)
}

func defaultColorFill(x, y, w, h int, cell termbox.Cell) {
	fill(x,y,w,h,cell, termbox.ColorDefault)
}

func fill(x, y, w, h int, cell termbox.Cell, color termbox.Attribute) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, color, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func drawSell(tile Tile, left int, top int, cellWidth int, cellHeight int) {
	const coldef = termbox.ColorDefault

	defaultColorFill(left, top, cellWidth, 1, termbox.Cell{Ch: '─'})
	defaultColorFill(left, top, 1, cellHeight, termbox.Cell{Ch: '|'})
	defaultColorFill(left, top+cellHeight, cellWidth, 1, termbox.Cell{Ch: '─'})
	defaultColorFill(left+cellWidth, top, 1, cellHeight, termbox.Cell{Ch: '|'})
	if !tile.isEmpty {
		tbprint(left+cellWidth/2, top+cellHeight/2, coldef, coldef, strconv.Itoa(tile.value))
	}
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
			}
		default:
		}
	}
}

type Drawer struct{}

func (d *Drawer) redraw(grid *Grid, score int, isOver bool) {
	if debugRun {
		dumpCell(grid, score, isOver)
	} else {
		gridDraw(grid, score, isOver)
	}
}
const GAME_TOP_OFFSET = 1

func gridDraw(grid *Grid, score int, isOver bool) {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	drawCellNumber(grid)

	//draw score
	drawScore(score)

	if isOver {
		drawOver()
	}

	termbox.Flush()
}
func drawOver(){
	gameOver := "Game Over"
	gameOverLen := len(gameOver)

	lastMessage := "If you quit it, please press ESC"

	top := GAME_HEIGHT / 2
	width := (GAME_WIDTH - gameOverLen) / 2
	color := termbox.ColorRed

	fill(0, top, width, 1, termbox.Cell{Ch: '='}, color)
	drawMessage(width, top, gameOver, color)
	fill(width + gameOverLen, top, width, 1, termbox.Cell{Ch: '='}, color)

	drawMessage((GAME_WIDTH - len(lastMessage) )/ 2, top  + 1, lastMessage, color)
}

func drawScore(score int){
	drawLine(0, 0, fmt.Sprintf("Score: %d", score))
}

func drawCellNumber(grid *Grid) {
	cellWidth := GAME_WIDTH / grid.size
	cellHeight := GAME_HEIGHT / grid.size

	for ly := 0; ly < grid.size; ly++ {
		for lx := 0; lx < grid.size; lx++ {
			tile := grid.cells[lx][ly]
			drawSell(tile, lx*cellWidth, GAME_TOP_OFFSET + ly*cellHeight, cellWidth, cellHeight)
		}
	}

}

func dumpCell(grid *Grid, score int, isOver bool) {
	fmt.Println("==========================================")
	sumValue := 0
	countIsNotEmpty := 0
	for ly := 0; ly < grid.size; ly++ {
		for lx := 0; lx < grid.size; lx++ {
			if !grid.cells[lx][ly].isEmpty {
				sumValue += grid.cells[lx][ly].value
				countIsNotEmpty += 1
			}
		}
	}
	fmt.Println("==================sumValue================", sumValue)
	fmt.Println("================countIsNotEmpty===========", (16 - countIsNotEmpty))
}

var (
	debugRun bool
)

func main() {
	flag.BoolVar(&debugRun, "debug", false, "debugRun flag")
	flag.Parse()

	drawer := Drawer{}
	gameState := Game{gridSize: 2, drawer: &drawer}
	gameState.setup()

	if debugRun {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			switch scanner.Text() {
			case "d":
				dispatch("down", &Message{data: "push left"})
				break
			case "l":
				dispatch("left", &Message{data: "push left"})
				break
			case "r":
				dispatch("right", &Message{data: "push left"})
				break
			case "u":
				dispatch("up", &Message{data: "push left"})
				break
			default:
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	} else {
		err := termbox.Init()
		//Error
		if err != nil {
			panic(err)
		}

		gridDraw(gameState.grid, 0, false)

		defer termbox.Close()

		handleKeyEvent()
	}
}
