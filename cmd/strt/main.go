package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " a  dadsf dafadf \\  "

	retStr := strings.TrimSpace(str)
	for i, v := range retStr {
		fmt.Printf("%d:%c\n", i, v)
	}
}
