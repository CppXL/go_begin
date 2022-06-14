package factory

import "fmt"

type User struct {
	Name string
	Ut   []Utype
}

type Utype uint8

var U User = *NewUser("aaa")

func NewUser(name string) *User {
	fmt.Println("init U on factory package")
	return &User{Name: name}
}

func Register(t Utype) {
	fmt.Println("register on register package")
	U.Ut = append(U.Ut, t)
}

func init() {
	fmt.Println("call init func on factory package")
}
