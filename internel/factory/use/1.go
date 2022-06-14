package use

import (
	"fmt"
	"project1/internel/factory/factory"
)

func init() {
	fmt.Println("call init func on use package")
	factory.Register(1)
}
