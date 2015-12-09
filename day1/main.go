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

	input := string(bytes)
	for _, char := range input {
		switch char {
		case '(':
			pos++
		case ')':
			pos--
		}
	}

	fmt.Println("Santa is on: ", pos)
}
