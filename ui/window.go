package ui

import "github.com/nsf/termbox-go"

type Window interface {
	Draw(x, y int)
	Event(e termbox.Event)
}
