package main

import (
	"fmt"
	"project1/internel/factory/factory"
	_ "project1/internel/factory/use"
)

func main() {
	fmt.Printf("%v\n", factory.U)
}
