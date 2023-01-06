package main

import (
	"fmt"
	"time"
	
	"github.com/gdamore/tcell/v2"

	"ejacks/editor"
)

func main() {
	fmt.Println("welcome to ejacks")

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()
	screen.EnablePaste()
	
	editor := editor.NewEditor(screen)
	editor.Run()
	time.Sleep(10 * time.Second)
}
