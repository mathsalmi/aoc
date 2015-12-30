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

			// skip on corners
			if isCorner(iline, icol) {
				continue
			}

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

func isCorner(iline, icol int) bool {
	len := len(grid)
	return (iline == 0 && icol == 0) || (iline+1 == len && icol == 0) || (iline == 0 && icol+1 == len) || (iline+1 == len && icol+1 == len)
}

// because math.Min and math.Max use float64 and fuck type casting…
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func shouldTurnOn(line, col int, isOff bool) bool {
	numneighbors := 0

	for i := max(0, line-1); i <= min(len(grid)-1, line+1); i++ {
		for j := max(0, col-1); j <= min(len(grid)-1, col+1); j++ {
			if (i != line || j != col) && grid[i][j] == on {
				numneighbors++
			}
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
