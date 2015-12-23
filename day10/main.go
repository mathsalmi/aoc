package main

import (
	"bytes"
	"fmt"
)

func main() {
	out := "1321131112"

	for i := 1; i <= 50; i++ {
		out = c(out)
	}

	fmt.Println("Final length: ", len(out))
}

func c(s string) string {
	out := bytes.Buffer{}
	input := []rune(s)
	len := len(input)

	for i := 0; i < len; i++ {
		c, count := input[i], 1

		j := i + 1
		for ; j < len && c == input[j]; j++ {
			count++
		}
		i = j - 1

		out.WriteString(fmt.Sprintf("%d%c", count, c))
	}

	return out.String()
}
