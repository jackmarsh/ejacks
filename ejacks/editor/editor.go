package editor

import (
	"fmt"
	
	"github.com/gdamore/tcell/v2"

	"ejacks/display"
)

// Editor is an emacs-esque text editor.
type Editor struct {
	screen tcell.Screen
	termEventChan chan tcell.Event
}

func NewEditor(screen tcell.Screen) *Editor {
	
	return &Editor{
		screen: screen,
		termEventChan: make(chan tcell.Event, 1),
	}
}

func (e *Editor) Run() {
	e.redraw()
	go e.termPoll()
	go e.run()
}

func (e *Editor) redraw() {
	display.DrawEditor(e.screen)
	e.screen.Sync()
}

func (e *Editor) termPoll() {
	for {
		event := e.screen.PollEvent()
		e.termEventChan <- event
	}
}

func (e *Editor) run() {
	for {
		select {
		case event := <-e.termEventChan:
			fmt.Println(fmt.Sprintf("%+v", event))
		}
	}
			

}
