package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	//Create sciter window object
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_TITLEBAR, nil)

	//set window title
	win.SetTitle("Simple Sciter Program")

	win.show()

	win.run()
}
