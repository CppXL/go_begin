package interface_test

import (
	"fmt"
	"net"
)

type (
	// 接口
	notifier interface {
		notify2()
		notify()
	}
	nt interface {
		nts()
	}
	users struct {
		name string
		addr net.IP
	}
	admin struct {
		users
		level int
	}
	test struct {
		i int
	}
)

// 指针接受者实现接口
func (u users) notify() {
	fmt.Printf("notify %s ip address %v\n", u.name, u.addr)
	for i := 0; i < len(u.addr); i++ {
		fmt.Printf("%v ", u.addr[i])
	}
}

func (u users) notify2() {
	fmt.Println("call notify2")
	fmt.Println(u)
}

func (t test) nts() {
	fmt.Println("t struct call interface nts")
}

func (u *users) nts() {
	fmt.Println("\nusers struct call interface nts")
}

func TInterface() {
	fmt.Println("call TInterface")
	u := users{"bill", net.ParseIP("192.168.50.1")}
	a := admin{
		users: u,
		level: 1,
	}
	u.notify()
	fmt.Printf("a type %T\n", a.users)
	var i notifier
	i = u
	i.notify()
}

func aa(n notifier) {
	n.notify()
}

func TInterface2() {
	fmt.Println("call TInterface2")
	var t test = test{1}
	var i nt = &t
	i = t
	i.nts()
	var u users = users{}
	i = &u
	i.nts()
	fmt.Println(t)
	fmt.Println("-----------------------")
}
