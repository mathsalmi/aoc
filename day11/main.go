package main

import (
	"fmt"
)

func main() {
	s := "vzbxkghb"

	for {
		s = increment(s)
		if isOk(s) {
			break
		}
	}

	fmt.Println("New password: ", s)
}

// Passwords:
// - 8 char length
// - DONT contain i, o, and l
// - CONTAIN TWO different non-overlapping pairs of letters
// - Passwords must include one increasing straight of at least three letters,
//   like abc, bcd, cde, and so on, up to xyz.

func isOk(s string) bool {
	v := []rune(s)
	len := len(v)

	seq := false
	pair := false
	tpair := 0

	for i := 0; i < len; i++ {
		r := v[i]

		// forb chars
		switch r {
		case 'i', 'o', 'l':
			return false
		}

		// sequence like 'abc'
		if i+2 < len && !seq && v[i+1] == r+1 && v[i+2] == r+2 {
			seq = true
		}

		if i+1 < len && !pair && r == v[i+1] {
			tpair++
			pair = true
		} else {
			pair = false
		}
	}

	return seq && tpair >= 2
}

func increment(s string) string { //(err error) { - let the world burn
	r := []rune(s)
	len := len(r)

	incrementAt(r, len-1)

	return string(r)
}

func incrementAt(r []rune, pos int) {
	if r[pos] == 'z' {
		r[pos] = 'a'
		incrementAt(r, pos-1)
		return
	}

	r[pos]++

	switch r[pos] {
	case 'i', 'o', 'l':
		r[pos]++
	}
}
