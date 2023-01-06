package display

import (
	"github.com/gdamore/tcell/v2"
)

// DrawSearchQuery draws the search query (if any) on the last line of the screen.
// This overwrites the status bar.
func DrawSearchQuery(screen tcell.Screen, query string) {
	
	screenWidth, screenHeight := screen.Size()
	if screenHeight == 0 {
		return
	}

	row := screenHeight - 1
	sr := NewScreenRegion(screen, 0, row, screenWidth, 1)
	sr.Fill(' ', tcell.StyleDefault)
	sr.SetContent(0, 0, rune('/'), nil, tcell.StyleDefault)
	//col := drawStringNoWrap(sr, query, 1, 0, tcell.StyleDefault)
	sr.ShowCursor(0, 0)
}
