package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello "))
	fmt.Fprintf(&b, "world")
	b.WriteTo(os.Stdout)
	b.Write([]byte("aaa"))
	fmt.Printf("b.Bytes(): %v\n", b.Bytes())
	var a bytes.Buffer = *bytes.NewBufferString("a")
	fmt.Printf("a.Bytes(): %v\n", a.Bytes())
	fmt.Printf("a.Cap(): %v\n", a.Cap())
	a.Grow(10)
	fmt.Printf("a.Cap(): %v\n", a.Cap())
	fmt.Printf("a.Len(): %v\n", a.Len())
	var c []byte = make([]byte, 10)
	b.Grow(64)
	// b.WriteTo(os.Stdout)

	// 从buffer中读取数据到指定缓冲区中，直到buffer耗尽
	_, err := b.Read(c)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("c: %v\n", c)
	b.WriteTo(os.Stdout)
	b.Write([]byte("bbb"))
	d, _ := b.ReadByte()
	fmt.Printf("d: %v\n", d)
	b.Write([]byte("ccc"))

	// bbc 直到读到了 ‘c’返回包含c的切片
	c, _ = b.ReadBytes('c')
	fmt.Printf("c: %v\n", string(c))

	// b.ReadFrom(os.Stdin)
	fmt.Printf("b.Bytes(): %v\n", string(b.Bytes()))
	s, _ := b.ReadString('\n')
	fmt.Printf("s: %v\n", s)

	b.Reset()
}
