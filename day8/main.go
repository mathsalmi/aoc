package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
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

	fmt.Println("Total ", strlen-vallen)
}

func parse(s string) (strlen, vallen int, ss string) {
	strlen = utf8.RuneCountInString(s)

	bf := bytes.Buffer{}
	reader := bufio.NewReader(strings.NewReader(s))
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}

		if c != '\\' {
			bf.WriteByte(c)
		} else {
			peek, _ := reader.Peek(3)
			if len(peek) > 0 {
				nc := peek[0]
				switch nc {
				case '\\', '"':
					bf.WriteByte(nc)
					reader.Discard(1)
				case 'x':
					dst := make([]byte, 1)
					hex.Decode(dst, peek[1:3])

					bf.Write(dst)
					reader.Discard(3)
				}
			}
		}
	}

	vallen = utf8.RuneCountInString(bf.String()) - 2 // remove the literals
	ss = bf.String()

	return
}
