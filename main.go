package main

import (
	"log"
	"time"
)

func main() {
	// Create a new game with a grid of 25x25 and a refresh duration of 1 second
	game, err := NewGame(25, 25, time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Run the game with a custom seed for a glider
	// The glider is a pattern that moves diagonally across the grid
	// By default, the grid is seeded randomly
	err = game.Run(WithSeed(func(u Grid) {
		height, width := u.Height(), u.Width()
		if height < 3 || width < 3 {
			panic("grid is too small for a glider")
		}

		y := height / 2
		x := width / 2

		u[y-1][x] = true
		u[y][x+1] = true
		u[y+1][x-1] = true
		u[y+1][x] = true
		u[y+1][x+1] = true
	}))
	if err != nil {
		log.Fatal(err)
	}
}
