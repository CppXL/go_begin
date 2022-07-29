package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func buffered(b []byte) {
	fmt.Println(fmt.Sprintf(call, "buffered"))
	br := bytes.NewReader(b)
	r := bufio.NewReader(br)
	fmt.Println("buffered:", r.Buffered())
	rb, _ := r.ReadByte()
	fmt.Println("read:", rb)
	fmt.Println("buffered:", r.Buffered())
	rb, _ = r.ReadByte()
	fmt.Println("read:", rb)
	fmt.Println("buffered:", r.Buffered())

}
