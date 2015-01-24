package ui

import "github.com/nsf/termbox-go"

func DrawBorder(x, y, width, height int, label string) {
	// Top line
	for i := x + 1; i < x+width-1; i++ {
		termbox.SetCell(i, y, HLine, termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(i, y+height-1, HLine, termbox.ColorWhite, termbox.ColorDefault)
	}

	// Left line
	for i := y + 1; i < y+height-1; i++ {
		termbox.SetCell(x, i, VLine, termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(x+width-1, i, VLine, termbox.ColorWhite, termbox.ColorDefault)
	}

	// Draw corners
	termbox.SetCell(x, y, TopLeftChar, termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y, TopRightChar, termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(x, y+height-1, BotLeftChar, termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y+height-1, BotRightChar, termbox.ColorWhite, termbox.ColorDefault)

	if label != "" {
		DrawString(x+2, y, " "+label+" ")
	}
}

func DrawString(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
}
