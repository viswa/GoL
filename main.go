// A graphical implementation of Conway's Game of Life
package main

import (
	"fmt"
	"os"
	"time"

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

// Width of a living cell
const LifeCellSize = 5

// drawUniverse draws the present state of the game to window
func drawUniverse(imd *imdraw.IMDraw, state *State) {
	for i := 0; i < WindowWidth/LifeCellSize; i++ {
		for j := 0; j < WindowHeight/LifeCellSize; j++ {
			if state.Present.Alive(i, j) {
				x1, y1 := float64(i*LifeCellSize), float64(j*LifeCellSize)
				x2, y2 := x1+LifeCellSize, y1+LifeCellSize
				imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
				imd.Rectangle(0)
			}
		}
	}
}

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

	universe := NewState(WindowWidth/LifeCellSize, WindowHeight/LifeCellSize)
	imd := imdraw.New(nil)

	for !win.Closed() {
		imd.Clear()
		win.Clear(colornames.White)
		imd.Color = colornames.Black
		drawUniverse(imd, universe)
		imd.Draw(win)
		win.Update()

		time.Sleep(40 * time.Millisecond)
		universe.Step()
	}
}

func main() {
	pixelgl.Run(run)
}
