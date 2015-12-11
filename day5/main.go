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
	tv := countVowels(s)
	if tv < 3 {
		return false
	}

	if !hasRepeatedRunes(s) {
		return false
	}

	if hasString(s) {
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

func countVowels(s string) int {
	out, vs := 0, []string{"a", "e", "i", "o", "u"}
	for _, v := range vs {
		out += strings.Count(s, v)
	}

	return out
}

func hasString(s string) bool {
	vs := []string{"ab", "cd", "pq", "xy"}
	for _, v := range vs {
		if strings.Contains(s, v) {
			return true
		}
	}

	return false
}

func hasRepeatedRunes(s string) bool {
	runes := []rune(s)
	for i, r := range runes {
		if i+1 < len(runes) && r == runes[i+1] {
			return true
		}
	}

	return false
}
