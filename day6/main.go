package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var lights [1000][1000]int

	items := readinput()
	doAction(&items, &lights)
	n := numLightsLit(&lights)
	fmt.Println("# lights lit ", n)
}

func doAction(items *[]*inst, lights *[1000][1000]int) {
	for _, item := range *items {
		for i := item.fromX; i <= item.toX; i++ {
			for j := item.fromY; j <= item.toY; j++ {
				switch item.comm {
				case "turn on":
					lights[i][j]++
				case "turn off":
					lights[i][j]--
					if lights[i][j] < 0 {
						lights[i][j] = 0
					}
				case "toggle":
					lights[i][j] = lights[i][j] + 2
				}
			}
		}
	}
}

func numLightsLit(lights *[1000][1000]int) (out int) {
	for _, line := range lights {
		for _, light := range line {
			out += light
		}
	}

	return out
}

func readinput() (out []*inst) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("Error opening the input file")
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := "", nil

	regex := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

	for err != io.EOF {
		if err != nil {
			log.Fatalln("Error parsing the input file", err)
		}

		line, err = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := regex.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}

		comm := matches[0][1]
		fromX, _ := strconv.Atoi(matches[0][2])
		fromY, _ := strconv.Atoi(matches[0][3])
		toX, _ := strconv.Atoi(matches[0][4])
		toY, _ := strconv.Atoi(matches[0][5])

		o := &inst{comm: comm, fromX: fromX, fromY: fromY, toX: toX, toY: toY}
		out = append(out, o)
	}

	return out
}

type inst struct {
	comm                   string
	fromX, fromY, toX, toY int
}
