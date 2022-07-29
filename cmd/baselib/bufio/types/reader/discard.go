package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func discard(b []byte) {
	fmt.Println(fmt.Sprintf(call, "discard"))
	br := bytes.NewReader(b)
	r := bufio.NewReader(br)
	rb, _ := r.ReadByte()

	fmt.Println("read:", rb)
	fmt.Println("buffered:", r.Buffered())
}
