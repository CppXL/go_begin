package init2

import "fmt"

func init() {
	fmt.Println("call init on init2 package")
}

type st struct{}

var S st = newst()

func newst() st {
	fmt.Println("call newst on init2 package")
	return st{}
}
