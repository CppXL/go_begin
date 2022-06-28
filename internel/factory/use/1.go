package use

import (
	"fmt"
	"gobegin/internel/factory/factory"
)

func init() {
	fmt.Println("call init func on use package")
	factory.Register(1)
}
