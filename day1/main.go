package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Error reading the input file")
	}

	pos := 0
	icharbasement := 0

	input := string(bytes)
	for index, char := range input {
		switch char {
		case '(':
			pos++
		case ')':
			pos--
		}

		if pos == -1 && icharbasement == 0 {
			icharbasement = index
		}
	}

	fmt.Println("Santa is on: ", pos)
	fmt.Println("Basement on char #: ", (icharbasement + 1))
}
