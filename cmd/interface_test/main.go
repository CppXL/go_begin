package main

import (
	"fmt"
	"gobegin/internel/interface_test"
)

type test struct {
	user string
	age  int
	t_interface
}

type t_interface interface {
	changeUser(u string)
	getUser() string
	getAge() int
	changeAge(age int)
}

type t2 struct {
	user string
	age  int
}

func (t *t2) changeUser(u string) {
	fmt.Println("changeUser")
	t.user = u
}

func (t *t2) changeAge(age int) {
	fmt.Println("changeAge")
	t.age = age
}

func (t *t2) getUser() string {
	return t.user
}

func (t *t2) getAge() int {
	return t.age
}

func (t *test) changeUser(u string) {
	t.user = u
}

func (t *test) changeAge(age int) {
	t.age = age
}

func (t *test) getUser() string {
	return t.user
}

// func (t *test) getAge() int {
// 	return t.age
// }

func intTest(t t_interface) {
	fmt.Println(t.getUser())
	// fmt.Println(t.getAge())
	t.changeUser("test")
	fmt.Println(t.getUser())
	t.changeAge(10)
	// fmt.Println(t.getAge())
}

func main() {
	interface_test.TInterface()
	interface_test.TInterface2()
	s := t2{"test", 10}
	var t test = test{user: "testaa", age: 10, t_interface: &s}
	intTest(&t)
	t.t_interface.changeUser("testaaaas")
	fmt.Println(t.t_interface.getUser())
}
