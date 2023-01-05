package main

import(
	"fmt"
	"time"
	
	"github.com/gdamore/tcell/v2"
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
	time.Sleep(2 * time.Second)

	screen.EnablePaste()
	editor := editor.NewEditor(screen)
	
}
