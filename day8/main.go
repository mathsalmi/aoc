package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	strlen, vallen := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, b, _ := parse(scanner.Text())
		strlen += a
		vallen += b
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Total ", vallen-strlen)
}

func parse(s string) (strlen, vallen int, ss string) {
	strlen = utf8.RuneCountInString(s)
	ss = fmt.Sprintf("%q", s)
	vallen = utf8.RuneCountInString(ss)

	return
}
