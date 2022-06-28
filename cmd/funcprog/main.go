package main

import (
	"fmt"
	"gobegin/internel/funcprog/closures"
)

func main() {
	closures.TClosure()
	type Person struct {
		firstName string
		lastName  string
		fullName  string
		age       int
	}
	var getFullName = func(in Person) string {
		in.firstName = "adf"
		in.fullName = in.firstName + in.lastName
		return in.fullName
	}

	john := Person{
		"john", "doe", "", 30,
	}

	fmt.Println(getFullName(john))
	fmt.Println(john)
	af := adder()
	fmt.Println(af(1))
	fmt.Println(af(2))
	fmt.Println(af(9))
}

func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		fmt.Println("\t", sum)
		sum += v
		return sum
	}
}
