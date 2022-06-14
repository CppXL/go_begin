package main

import "fmt"

type s struct{}

type h struct{ i }

type i interface {
	a()
	b()
	c()
}

func (H h) a() {
	fmt.Println("a")
}

func (H h) b() {
	fmt.Println("b")
}

func (s *s) a() {
	fmt.Println("call struct s's a func")
}

func (s *s) b() {
	fmt.Println("call struct s's b func")
}

func (s *s) c() {
	fmt.Println("call struct s's c func")
}

func call(I i) {
	I.a()
	I.b()
}

func main() {
	SS := new(s)
	call(SS)
	SS.c()
	HH := h{}
	call(&HH)
}
