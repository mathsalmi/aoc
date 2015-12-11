package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln("Error reading the input file")
	}

	// houses: map[X,Y position] number of gifts
	h := make(map[string]int)
	s, rs := new(santa), new(santa) // santa and robot-santa (rb)

	curr := s // either santa or rb

	coords := string(bytes)
	for _, coord := range coords {
		if coord == ' ' || coord == '\n' {
			continue
		}

		// first position
		if _, key := h["0,0"]; curr.x == 0 && curr.y == 0 && !key {
			h["0,0"] = 1
		}

		// move direction
		switch coord {
		case '^':
			curr.y++
		case 'v':
			curr.y--
		case '>':
			curr.x++
		case '<':
			curr.x--
		default:
			log.Fatalln("Position unknown '" + string(coord) + "'")
		}

		h[curr.pos()]++

		if curr == s {
			curr = rs
		} else {
			curr = s
		}
	}

	total := 0
	for _, value := range h {
		if value >= 1 {
			total++
		}
	}

	fmt.Println("# of houses with >= 1 presents: ", total)
}

type santa struct {
	x, y int
}

func (s santa) pos() string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}
