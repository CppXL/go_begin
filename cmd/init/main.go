package main

import (
	"fmt"
	_ "project1/init"
)

func init() {
	fmt.Println("1 call init")
}

type st struct {
	Name string
}

var S *st = news()

func news() *st {
	fmt.Println("call news")
	return &st{}
}

func main() {
	fmt.Println("call main")

}

func init() {
	fmt.Println("2 call init")
}
