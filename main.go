package main

import (
	"./memo"

	"github.com/nsf/termbox-go"
)

// [*]入力欄を描画する
// [*]入力する
// ファイルの内容を描画する
// フィルタリングする
// 編集する
// 追加する
// 検索する

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	memo.ReDraw_all()
	termbox.Flush()
	//inputmode := 0

	memo.Inputloop()
}
