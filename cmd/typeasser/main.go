package main

import "fmt"

func main() {
	var a interface{} = 10
	t, ok := a.(int)
	if ok {
		fmt.Printf("%T\n", t)
	} else {
		fmt.Print("not int")
	}

	fmt.Println("--------------------------")
	s, ok := a.(string)
	if ok {
		fmt.Printf("%T\n", s)
	} else {
		fmt.Printf("%T\tvalue:%v\n", s, s)
		if s == "" {
			fmt.Print("yes")
		} else {
			fmt.Println("no")
		}
		fmt.Print("not string")
	}
	switch x := a.(type) {
	case int:
		fmt.Printf("x is %T\n", x)
	case string:
		fmt.Printf("x is %T\n", x)
	default:
		fmt.Printf("x is %T\n", x)
	}
}
