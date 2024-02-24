package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type GameOption func(*Game)

func WithSeed(seed func(Grid)) GameOption {
	return func(g *Game) {
		seed(g.grid)
	}
}

type Game struct {
	screen          tcell.Screen
	grid            Grid
	refreshDuration time.Duration
}

func (g *Game) Run(options ...GameOption) error {
	if len(options) == 0 {
		options = append(options, WithSeed(func(u Grid) {
			u.SeedRandom()
		}))
	}
	for _, opt := range options {
		opt(g)
	}
	g.grid.Draw(g.screen)

	// Create a ticker to update the grid every specified duration
	ticker := time.NewTicker(g.refreshDuration)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// Update the grid every second
				UpdateGridState(g.grid)
				g.grid.Draw(g.screen)
			}
		}
	}()

	for {
		// Poll event for key press to exit the game
		ev := g.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				g.screen.Fini()
				done <- true
				return nil
			}
		}
	}
}

func NewGame(height, width int, refreshDuration time.Duration) (*Game, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	grid, err := NewGrid(height, width)
	if err != nil {
		return nil, err
	}

	return &Game{
		grid:            grid,
		screen:          screen,
		refreshDuration: refreshDuration,
	}, nil
}

func UpdateGridState(grid Grid) {
	tempGrid := make(Grid, len(grid))
	for i := range grid {
		tempGrid[i] = make([]bool, len(grid[i]))
		copy(tempGrid[i], grid[i])
	}

	for y := range grid {
		for x := range grid[y] {
			tempGrid[y][x] = grid.Next(x, y)
		}
	}

	for i := range grid {
		copy(grid[i], tempGrid[i])
	}
}
