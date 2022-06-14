package main

import (
	"fmt"
)

func main() {

	var a int = 10
	var b int = a &^ 5
	fmt.Println(b)
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)

	s := make([]int, 10, 10)
	s[1] = 2
	s[4] = 5
	s[5] = 6
	fmt.Printf("%d %d\n", len(s), cap(s))
	ss := s[1:5]
	fmt.Printf("%d %d\n", len(ss), cap(ss))
	fmt.Printf("%d \n", ss)
	sa := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sas := sa[:]
	fmt.Printf("%v %[1]q %[1]T\n", sas)

	p := "bcdefg"
	ps := p[:3]
	fmt.Printf("%v %v\n", &p, &ps)

}
