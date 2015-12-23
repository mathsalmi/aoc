package main

import "fmt"

func main() {
	out := "1321131112"

	for i := 1; i <= 40; i++ {
		out = c(out)
	}

	fmt.Println("Final length: ", len(out))
}

func c(s string) (out string) {
	input := []rune(s)
	len := len(input)

	for i := 0; i < len; i++ {
		c, count := input[i], 1

		j := i + 1
		for ; j < len && c == input[j]; j++ {
			count++
		}
		i = j - 1

		out += fmt.Sprintf("%d%c", count, c)
	}

	return out
}
