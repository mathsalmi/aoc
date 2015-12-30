package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	off = '.'
	on  = '#'
)

var grid [][]rune

type action struct {
	line, col int
	action    rune
}

var stack []action

func main() {
	for i := 0; i < 100; i++ {
		step()
	}

	t := 0
	for _, line := range grid {
		for _, col := range line {
			if col == on {
				t++
			}
		}
	}

	fmt.Println("Total on: ", t)
}

func init() {
	file, _ := os.Open("input")
	defer file.Close()

	bf := bufio.NewScanner(file)
	for bf.Scan() {
		line := bf.Text()
		grid = append(grid, []rune(line))
	}
}

func step() {
	for iline, line := range grid {
		for icol, col := range line {
			a := action{}
			a.line = iline
			a.col = icol

			if shouldTurnOn(iline, icol, col == off) {
				a.action = on
			} else {
				a.action = off
			}

			stack = append(stack, a)
		}
	}

	for _, a := range stack {
		grid[a.line][a.col] = a.action
	}
}

func shouldTurnOn(line, col int, isOff bool) bool {
	numneighbors := 0
	inc := func(i, j int) {
		if grid[i][j] == on {
			numneighbors++
		}
	}

	// line above
	if line > 0 {
		if col > 0 {
			inc(line-1, col-1)
		}

		inc(line-1, col)

		if (col + 1) < len(grid[line-1]) {
			inc(line-1, col+1)
		}
	}

	// same line
	if col > 0 {
		inc(line, col-1)
	}

	if (col + 1) < len(grid[line]) {
		inc(line, col+1)
	}

	// next line
	if line+1 < len(grid) {
		if col > 0 {
			inc(line+1, col-1)
		}

		inc(line+1, col)

		if col+1 < len(grid[line+1]) {
			inc(line+1, col+1)
		}
	}

	// if off
	if isOff {
		if numneighbors == 3 {
			return true
		}
		return false
	}

	// if on
	if numneighbors == 2 || numneighbors == 3 {
		return true
	}
	return false
}
