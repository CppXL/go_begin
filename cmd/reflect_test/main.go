package main

import (
	"fmt"
	"project1/internel/reflect_test"
	"reflect"
)

type (
	user struct {
		name string
		age  int
	}
	ci interface {
		ChangeName(s string)
	}
	ti interface {
		ShowAge()
	}
)

func (u *user) ChangeName(s string) {
	u.name = s
}

func (u user) ShowAge() {
	fmt.Println("age:", u.age)
}

func interfaceT(c ci, s string) {
	c.ChangeName(s)
}

func main() {
	var a string = reflect_test.TReflect(5)
	fmt.Printf("%v\n", a)
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
	var u user = user{name: "cqy", age: 1}
	fmt.Printf("user:%+v\n", u)
	u.ChangeName("sss")
	fmt.Printf("user:%v\n", u)
	fmt.Println("typeof u:", reflect.TypeOf(u))
	fmt.Println("Value Of u:", reflect.ValueOf(u))

	// fmt.Printf("Method of struct user:%v\n", reflect.Value.Interface(reflect.ValueOf(u)))
	v := reflect.ValueOf(u)
	fmt.Println(v.Interface().(user))

	// 通过valueof 获取变量指针
	// 通过s.Elem获取指针指向的变量
	//调用reflect.Value.SetInt设置新的值
	// 使用指针是因为调用函数时是值传递 使用指针来使得传入的是同一个变量
	i := 1
	s := reflect.ValueOf(&i)
	fmt.Println("i address=>", s)
	s.Elem().SetInt(10)
	fmt.Println("i=>", i)
	interfaceT(&u, "saaassss")
	fmt.Printf("user:%v\n", u)
	fmt.Println("-------------------------------------")
	typeofiT := reflect.TypeOf((*ci)(nil)).Elem()
	typeofcT := reflect.TypeOf((*ti)(nil)).Elem()
	userErrorPtr := reflect.TypeOf(&user{})
	userError := reflect.TypeOf(user{})
	// 指针接受者 func (u *user) ChangeName(s string)
	fmt.Println(userErrorPtr.Implements(typeofiT))
	fmt.Println(userError.Implements(typeofiT))
	// 值接受者 都实现了方法 func (u user) ShowAge()
	fmt.Println(userErrorPtr.Implements(typeofcT))
	fmt.Println(userError.Implements(typeofcT))
	fmt.Println("----------动态方法调用----------------")
	reflect_test.CallMethod()
}
