// A graphical implementation of Conway's Game of Life
package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Window Dimensions
const (
	WindowWidth  = 640
	WindowHeight = 480
)

// run is the implicit entry point to the program
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application launch failed")
		os.Exit(1)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		imd.Clear()

		win.Clear(colornames.White)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
