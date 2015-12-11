package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	Init()
	Clear()
	Flush()
}

func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Flush() {
	termbox.Flush()
}
