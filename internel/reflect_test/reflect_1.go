package reflect_test

import (
	"fmt"
	"reflect"
	"strconv"
)

func TReflect(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

func CallMethod() {
	// 通过 reflect.ValueOf 获取函数 Add 对应的反射对象；
	// 调用 reflect.rtype.NumIn 获取函数的入参个数；
	// 多次调用 reflect.ValueOf 函数逐一设置 argv 数组中的各个参数；
	// 调用反射对象 Add 的 reflect.Value.Call 方法并传入参数列表；
	// 获取返回值数组、验证数组的长度以及类型并打印其中的数据；
	s := add
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	argv[0] = reflect.ValueOf(11)
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int())
}
func add(a, b int) int { return a + b }
