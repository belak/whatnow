package main

import (
	"./ui"
	"github.com/nsf/termbox-go"
)

func main() {
	// Hello World!
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	width, height := termbox.Size()
	var win ui.Window = &GameWindow{}

	// Do our initial draw
	win.Draw(width, height)
	termbox.Flush()

	for {
		e := termbox.PollEvent()
		switch e.Type {
		case termbox.EventResize:
			width = e.Width
			height = e.Height
		case termbox.EventKey:
			if e.Key == termbox.KeyEsc {
				// If esc was pressed, kill everything
				return
			} else if e.Key == termbox.KeyCtrlL {
				termbox.Sync()
			}

			win.Event(e)
		}

		win.Draw(width, height)
		termbox.Flush()
	}
}
