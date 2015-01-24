package main

import (
	"fmt"

	"./ui"
	"github.com/nsf/termbox-go"
)

type GameWindow struct {
	Events []string
}

func (w *GameWindow) Draw(width, height int) {
	// Main pane
	ui.DrawBorder(0, 0, width-40, height-10, "")

	// Actions - 40 characters on the right
	ui.DrawBorder(width-40, 0, 40, height, "")

	// Stats - 10 characters on the bottom
	ui.DrawBorder(0, height-10, width-40, 10, "")

	// Draw the events we know about
	for k, v := range w.Events {
		ui.DrawString(1, height-10+k+1, v)
	}
}

func (w *GameWindow) Event(e termbox.Event) {
	w.Events = append(w.Events, fmt.Sprintf("%v", e))
	if len(w.Events) > 8 {
		w.Events = w.Events[1:]
	}
}
