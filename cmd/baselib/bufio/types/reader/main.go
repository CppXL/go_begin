package main

import (
	"bufio"
	"bytes"
	"fmt"
)

var call = "---------------call %s---------------"

func main() {
	data := []byte("Hello World\nhey there\nwho are you\n")
	// 实现了 io.Reader接口
	//br := bytes.NewReader(data)
	//
	//// 实例化bufio.Reader
	//r := bufio.NewReader(br)
	buffered(data)
	discard(data)
	// fmt.Printf("len(data): %v\n", len(data))

	// // 一开始w=r=0
	// fmt.Printf("r.Buffered(): %v\n", r.Buffered())

	// // readbytes
	// fmt.Println(r.ReadByte())
	// fmt.Printf("r.Buffered(): %v\n", r.Buffered())

	// fmt.Println(r.ReadByte())
	// fmt.Printf("r.Buffered(): %v\n", r.Buffered())

	// fmt.Println(r.ReadByte())
	// fmt.Printf("r.Buffered(): %v\n", r.Buffered())

	// fmt.Println(r.Discard(2))
	// l, _, _ := r.ReadLine()
	// s, _ := r.Peek(5)
	// fmt.Println(string(s))
	// fmt.Println(string(l), l)

	//b := make([]byte, r.Buffered())
	//// r.Read(b)
	//s, _ := r.ReadBytes('t')
	//// s, _,_ := r.ReadLine()
	//// r.Reset()
	//// r.ReadSlice()
	//fmt.Println("s:", string(s))
	//fmt.Printf("r.Buffered(): %v\n", r.Buffered())
	//fmt.Printf("b: %s\t%d\n", string(b), len(b))
	//b = []byte("Imaaa\naa")
	//ba := bytes.NewReader(b)
	//r.Reset(ba)
	//fmt.Printf("r.Buffered(): %v\n", r.Buffered())
	//fmt.Println(r.ReadString('\n'))
	//fmt.Printf("r.Buffered(): %v\n", r.Buffered())
	//fmt.Printf("r.Size(): %v\n", r.Size())
	//err := r.UnreadByte()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//writeto()
}

func writeto() {
	data := []byte("Hello World\nhey there\nwho are you\n")
	br := bytes.NewReader(data)
	r := bufio.NewReader(br)
	b := make([]byte, len(data))
	bn := bytes.NewBuffer(b)
	r.WriteTo(bn)
	fmt.Println(bn.Bytes())
	fmt.Printf("b: %v\n", b)
}
