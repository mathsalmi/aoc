package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readinput()

	tok := 0
	for _, line := range lines {
		if isStringOk(line) {
			tok++
		}
	}

	fmt.Println("# strings ok: ", tok)
}

func isStringOk(s string) bool {
	// a pair of runes that repeat at least twice
	if !hasPairRepeatTwice(s) {
		return false
	}

	// repeated runes with exactly one rune between them
	if !hasPausedRepetition(s) {
		return false
	}

	return true
}

func readinput() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("Error opening the input file")
	}

	r := bufio.NewReader(file)
	line, err := "", nil
	var out []string

	for err != io.EOF {
		if err != nil {
			log.Fatalln("Error parsing the input file", err)
		}

		line, err = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		out = append(out, line)
	}

	return out
}

// hasPausedRepetition tells whether or not the string
// has runes that repeat with exactly one rune between them.
//
// ex.: aba or cdc or even aaa
func hasPausedRepetition(s string) bool {
	r := []rune(s)
	for i := 0; i+2 < len(s); i++ {
		if r[i] == r[i+2] {
			return true
		}
	}
	return false
}

func hasPairRepeatTwice(s string) bool {
	r := []rune(s)
	len := len(s)
	pair := ""

	for i := 0; i < len; i++ {
		pair = string(r[i : i+2])

		if strings.Count(s, pair) >= 2 {
			return true
		}
	}

	return false
}
