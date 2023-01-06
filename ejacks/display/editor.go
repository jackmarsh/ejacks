package display

import (
	"github.com/gdamore/tcell/v2"
)

func DrawEditor(screen tcell.Screen) {
	screen.Fill(' ', tcell.StyleDefault)
	DrawSearchQuery(screen, "hello")
}
