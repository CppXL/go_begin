package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := "Hello World\nhey there\nwho are you\n"
	bdata := []byte(data)
	// Scanner.Scan
	scanner := bufio.NewScanner(bytes.NewReader(bdata))
	scanner.Scan()
	// print `Hello world`
	fmt.Println(scanner.Text())
	// print `hey there`
	scanner.Scan()
	fmt.Println(scanner.Text())

}
