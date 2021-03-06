package memo

import (
  "github.com/nsf/termbox-go"
)

func Inputloop() {

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				break loop
			case termbox.KeyCtrlE:
				// TODO: 編集
			case termbox.KeyArrowUp:
				if lineno > 1 {
					lineno--
				}
			case termbox.KeyArrowDown:
				lineno++
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
				}
			}

			ReDraw_all()

		case termbox.EventError:
			panic(ev.Err)
		}
	}
}


