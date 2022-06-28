package main

import (
	"fmt"
	"gobegin/internel/struct_test"
)

func main() {
	struct_test.TStruct()
	administrator := struct_test.GetAdministrator("aaa", false, 2, 2, 1)
	fmt.Printf("%v", administrator)
	base := struct_test.Sec{
		Base: &struct_test.Base{Name: "a"},
	}
	fmt.Printf("%+v\n", base)
	struct_test.PrintSec()
}
