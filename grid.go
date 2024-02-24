package main

import (
	"errors"
	"math/rand"

	"github.com/gdamore/tcell"
)

var (
	cellStyle       = tcell.StyleDefault.Background(tcell.ColorBlack)
	backgroundStyle = tcell.StyleDefault.Background(tcell.ColorWhite)
)

type Grid [][]bool

func NewGrid(height, width int) (Grid, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("height and width must be positive")
	}

	u := make(Grid, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u, nil
}

func (g Grid) SeedRandom() {
	height, width := g.Height(), g.Width()
	for i := 0; i < (width * height / 4); i++ {
		// Selecting random cells to seed
		g[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func (g Grid) Draw(screen tcell.Screen) {
	screen.Clear()
	for i, row := range g {
		for j, cell := range row {
			if cell {
				screen.SetCell(j*2, i, cellStyle, ' ')
				screen.SetCell(j*2+1, i, cellStyle, ' ')
			} else {
				screen.SetCell(j*2, i, backgroundStyle, ' ')
				screen.SetCell(j*2+1, i, backgroundStyle, ' ')
			}
		}
	}
	screen.Show()
}

func (g Grid) Neighbors(x, y int) int {
	height, width := g.Height(), g.Width()
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			nx = (nx + width) % width
			ny = (ny + height) % height
			if g[ny][nx] {
				alive++
			}
		}
	}
	return alive
}

func (g Grid) Next(x, y int) bool {
	aliveNeighbours := g.Neighbors(x, y)
	return aliveNeighbours == 3 || aliveNeighbours == 2 && g[y][x]
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	return len(g[0])
}
