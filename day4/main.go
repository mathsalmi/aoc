package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i, input := 0, "yzbqklnj"

	hash := ""
	for ; !strings.HasPrefix(hash, "000000"); i++ {
		x := md5.Sum([]byte(input + strconv.Itoa(i)))
		hash = fmt.Sprintf("%x", x[:])
	}

	fmt.Println("# is ", (i - 1), hash)
}
