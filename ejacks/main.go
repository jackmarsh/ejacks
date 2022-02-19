package main

import(
	"fmt"
	
	"github.com/gdamore/tcell/termbox"
)

func main() {
     fmt.Println("welcome to ejacks")
     err := termbox.Init()
     if err != nil {
         panic(err)
     }
     defer termbox.Close()
     termbox.SetInputMode(termbox.InputAlt)
}