package iorwtest

import (
	"bytes"
	"fmt"
	"os"
)

func TIotest1() {
	fmt.Println("call TIotest1")
	var b bytes.Buffer
	// a := []int{1, 2, 3, 4}
	b.Write([]byte("hello "))
	fmt.Fprintf(&b, "world")
	b.WriteTo(os.Stdout)
	fmt.Println("-----------------------------")
}
