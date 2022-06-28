package main

import (
	"fmt"
	"gobegin/internel/factory/factory"
	_ "gobegin/internel/factory/use"
)

func main() {
	fmt.Printf("%v\n", factory.U)
}
