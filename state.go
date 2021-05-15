package main

import (
	"math/rand"
)

// State represents the present and future state of Universe
type State struct {
	Present, future *Universe
	width, height   int
}

// NewState returns a new State with random initial values
func NewState(w, h int) *State {
	Present := NewUniverse(w, h)
	for i := 0; i < (w * h / 4); i++ {
		Present.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &State{
		Present: Present, future: NewUniverse(w, h),
		width: w, height: h,
	}
}

// Step advances the game by one generation
func (s *State) Step() {
	// Update future from the current state
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			s.future.Set(x, y, s.Present.Next(x, y))
		}
	}
	// Swap future with present
	s.Present, s.future = s.future, s.Present
}
