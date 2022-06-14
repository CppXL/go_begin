package closures

import "fmt"

func TClosure() {
	base10 := add(10)
	base20 := add(20)
	base30 := add(30)
	fmt.Println("10+10=", base10(10))
	fmt.Println("10+10=", base20(10))
	fmt.Println("10+10=", base30(10))
}

func add(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}

func muti(add bool, onAdd, onMuti int) int {
	if add {
		return onAdd
	}
	return onMuti
}
