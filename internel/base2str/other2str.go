package base2str

import "fmt"

func TTran() {
	var num1 int = 5
	var num2 float32 = 1.11
	var num3 int64 = int64(num1)
	var b bool = true
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T value is %v\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T value is %v\n", str, str)
	str = fmt.Sprintf("%d", num3)
	fmt.Printf("str type %T value is %v\n", str, str)
	str = fmt.Sprintf("%t", b)

	fmt.Printf("str type %T value is %v\n", str, str)
	str = fmt.Sprintf("%t", b)
	// 加双引号
	fmt.Printf("str type %T value is %q\n", str, str)
	str = "asd"
	fmt.Printf("str type %T value is %q\n", str, str)
	var str2 string = "sss"
	fmt.Println(str2)
}
