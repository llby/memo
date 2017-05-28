package main

import (
	"./memo"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	memo.OpenFile()
	memo.ReDraw_all()
	termbox.Flush()

	memo.Inputloop()
}
