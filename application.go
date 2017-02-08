package main

import (
	"github.com/nsf/termbox-go"
)

//func keyEventLoop(kch chan termbox.Key) {
//	for {
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			kch <- ev.Key
//		default:
//		}
//	}
//}

/*

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}
func main() {
	print("hello world")
	for{
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowDown{
				print("hello world")
				drawLine(0, 1, "--------------------------------------------------------------------------------")
				termbox.Flush()
				break
			}
		default:
		}
	}

	defer termbox.Close()
}
*/

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

func initdraw() {
	//termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//
	//drawLine(0, 0, "Press ESC to exit.")
	//drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
	//
	//termbox.Flush()
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()
	//oneCellWidth :=  20
	midy := h / 2
	midx := (w - 30) / 2

	println(midy)
	//termbox.SetCell(midx, oneCellWidth + 1, '-', coldef, coldef)

	//midx := (w - edit_box_width) / 2
	//
	//// unicode box drawing chars around the edit box
	//termbox.SetCell(midx-1, midy, '│', coldef, coldef)
	//termbox.SetCell(midx+edit_box_width, midy, '│', coldef, coldef)
	//termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
	//termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
	//termbox.SetCell(midx+edit_box_width, midy-1, '┐', coldef, coldef)
	//termbox.SetCell(midx+edit_box_width, midy+1, '┘', coldef, coldef)
	//fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '─'})
	fill(midx, midy+1, 30, 1, termbox.Cell{Ch: '─'})
	fill(midx, midy+1, 1, 30, termbox.Cell{Ch: '|'})
	println("=======midy============", midy)
	//fill(midx, midy+1, 30, 1, termbox.Cell{Ch: '─'})
	//
	//edit_box.Draw(midx, midy, edit_box_width, 1)
	//termbox.SetCursor(midx+edit_box.CursorX(), midy)

	//tbprint(midx+6, midy+3, coldef, coldef, "Press ESC to quit")
	termbox.Flush()
}

func drawMessage(msg string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, msg)

	termbox.Flush()
}

func pollEvent() {
	initdraw()
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
				initdraw()
			}
		default:
			initdraw()
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	pollEvent()
}
