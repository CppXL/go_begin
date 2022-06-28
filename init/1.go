package init

import (
	"fmt"
	_ "gobegin/init1"
	_ "gobegin/init2"
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
