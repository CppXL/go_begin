package main

import (
	"bytes"
	"fmt"
)

func main() {

	var r bytes.Reader = *bytes.NewReader([]byte("hello "))
	fmt.Printf("r.Len(): %v\n", r.Len())
	r.Read([]byte("world"))
	fmt.Println(r.ReadByte())

}
