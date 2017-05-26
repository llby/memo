package memo

import (
	//"io/ioutil"
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
)

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func open_file() []string {
	var fp *os.File
	var err error
	var res []string

	if len(os.Args) < 2 {
		fp  = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return res
}

func PrintFile() {

	rows := open_file()
	y := 1
	for _, v := range rows {
		print_tb(0, y, termbox.ColorWhite, termbox.ColorBlack, v)
		y++
	}


}

