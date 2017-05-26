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
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
				}
			}

			//draw_keyboard()
			//dispatch_press(&ev)
			//pretty_print_press(&ev)
			ReDraw_all()


		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
