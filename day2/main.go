package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	boxes := readinput()

	tsurface := 0
	tribbon := 0

	for _, box := range boxes {
		o := []int{box.l, box.w, box.h}
		sort.Sort(sort.IntSlice(o))

		a, b, c := box.l*box.w, box.w*box.h, box.h*box.l
		tsurface += (2*a + 2*b + 2*c) + o[0]
		tribbon += (2*o[0] + 2*o[1]) + (box.l * box.w * box.h)
	}

	fmt.Println("total surface is: ", tsurface, "ft^2")
	fmt.Println("total ribbon is: ", tribbon, "ft")
}

func readinput() []*box {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("Error reading the input file", err)
	}

	var boxes []*box

	r := bufio.NewReader(file)
	line, err := "", nil

	for err != io.EOF {
		if err != nil {
			log.Fatalln("Error parsing the input file", err)
		}

		line, err = r.ReadString('\n')
		if line == "" {
			continue
		}

		box, err := parse(line)
		if err != nil {
			log.Fatalln("Error parsing the input file", err)
		}

		boxes = append(boxes, box)
	}

	return boxes
}

func parse(line string) (*box, error) {
	line = strings.Trim(line, "\n")
	vals := strings.Split(line, "x")
	if len(vals) != 3 {
		return nil, fmt.Errorf("error parsing line: %s", line)
	}

	length, err := strconv.Atoi(vals[0])
	width, err := strconv.Atoi(vals[1])
	height, err := strconv.Atoi(vals[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing line: %s", line)
	}

	return &box{l: length, w: width, h: height}, nil
}

type box struct {
	l, w, h int
}
