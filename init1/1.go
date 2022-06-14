package init1

import "fmt"

func init() {
	fmt.Println("call init on init1 package")
}

type st struct{}

var S st = newst()

func newst() st {
	fmt.Println("call newst on init1 package")
	return st{}
}
