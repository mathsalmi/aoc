package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		panic("tell the start number and the gifts number")
	}

	startnum, _ := strconv.Atoi(os.Args[1])
	maxgifts, _ := strconv.Atoi(os.Args[2])

	for house := startnum; ; house++ {
		tgifts := 0

		for elf := 1; elf <= house; elf++ {
			if house%elf == 0 {
				tgifts += elf * 10
			}
		}

		if tgifts >= maxgifts {
			fmt.Println("House # ", house)
			return
		}
	}

}
