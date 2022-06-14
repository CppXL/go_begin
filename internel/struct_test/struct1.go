package struct_test

import (
	"fmt"
)

type user struct {
	username string
	isAdmin  bool
	ext      int
	age      int
}

type admin struct {
	user
	level int
}

func (u user) print() {
	fmt.Printf("user struct value:{username:%v isAdmin:%v ext:%v age:%v}\n", u.username, u.isAdmin, u.ext, u.age)
}
func (u user) notify() {
	fmt.Println("call notify function")
}
func (u *user) changeName(newName string) {
	u.username = newName
}

func TStruct() {
	fmt.Println("call struct_test.TStruct")
	var bill user
	bill.username = "cqy"
	bill.isAdmin = true
	bill.age = 22
	bill.ext = 0
	lisa := user{
		username: "lisa",
		isAdmin:  false,
		age:      18,
		ext:      11,
	}

	fred := admin{
		user: user{
			username: "fred",
			isAdmin:  true,
			age:      10,
			ext:      10,
		},
		level: 1,
	}

	//fears := admin{
	//	person: lisa,
	//	level:  2
	//}
	fmt.Printf("lisa value %+v \nfred value %+v\n", lisa, fred)
	bill.print()
	// 使用指针调用方法 语法糖 直接使用pBill.print也一样
	bill.changeName("ppp")
	pBill := &bill
	(*pBill).print()

	admins := Administrator{
		admin: &fred,
	}
	admins.user.print()
	admins.user.isAdmin = false
	admins.user.print()
	fred.user.notify()
	// 因为user类型直接嵌入了admin类型，所以user的值接收者方法被提升到了外部类型admin中，
	fred.print()
	fmt.Printf("type of admins: %T\n", admins)
	fmt.Println("------------------------------")

}

func GetAdministrator(username string, isAdmin bool, ext int, age int, level int) Administrator {
	person := user{
		username: username,
		isAdmin:  isAdmin,
		ext:      ext,
		age:      age,
	}
	administrator := Administrator{
		admin: &admin{
			user:  person,
			level: level,
		},
	}
	return administrator
}

func PrintSec() {
	base := Sec{
		Base: &Base{Name: "a"},
	}
	fmt.Printf("%v\n", base)
	th := Third{
		Base{Name: "dddd"},
	}
	fmt.Printf("%+v\n", base.Base)
	fmt.Printf("%+v\n", th)
}
