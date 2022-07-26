package main

import (
	"bufio"
	"bytes"
	"fmt"
	"unsafe"
)

func main() {
	data := "Hello World\nhey there\nwho are you\n"
	bdata := []byte(data)
	// Scanner.Scan
	scanner := bufio.NewScanner(bytes.NewReader(bdata))

	adata := make([]byte, 256)
	// scanner.Buffer(adata, len(adata))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Printf("adata: %v\n", string(adata))
	fmt.Printf("addr of adata:%v\naddr of bdata:%v\n", (unsafe.Pointer)(&adata), (unsafe.Pointer)(&bdata))
}
