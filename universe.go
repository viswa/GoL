package main

// Universe represents a 2D grid of cells
type Universe struct {
	grid          [][]bool
	width, height int
}

// NewUniverse returns an empty universe of the spcified width and height
func NewUniverse(w, h int) *Universe {
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}
	return &Universe{grid: grid, width: w, height: h}
}

// Set sets the state of the cell specified
func (u *Universe) Set(x, y int, newState bool) {
	u.grid[y][x] = newState
}

// Alive returns whether the specified cell is alive
func (u *Universe) Alive(x, y int) bool {
	x = (x + u.width) % u.width   // wrapping around x near edges
	y = (y + u.height) % u.height // wrapping around y near edges
	return u.grid[y][x]
}

// Next returns state of a cell for the next generation
func (u *Universe) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// only count neighbours, not the cell specified
			if (j != 0 || i != 0) && u.Alive(x+i, y+j) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && u.Alive(x, y)
}
