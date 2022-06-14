package init

import (
	"fmt"
	_ "project1/init1"
	_ "project1/init2"
)

type st struct{}

var S st = newst()

func newst() st {
	fmt.Println("call newst on init package")
	return st{}
}

func init() {
	fmt.Println("call init on init package")
}

func T() {

}
