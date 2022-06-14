package array

import "fmt"

func SingleArray() {
	// 声明包含五个元素的int型数组
	var array = [5]int{1, 2, 3, 4, 5}
	array[1] = 100
	// 声明包含5个int型指针数组
	arrayP := [5]*int{}
	for i := 0; i < 5; i++ {
		arrayP[i] = &array[i]

	}
	// 声明指针数组 并为0 1分配int变量
	arrayP1 := [5]*int{0: new(int), 1: new(int)}
	*arrayP1[0] = 10
	for i := 0; i < 5; i++ {
		fmt.Println("arrayP:", *arrayP[i])
	}
	fmt.Println("arrayP1=>", *arrayP1[0])
	var array1 [5]int
	array1 = array
	for i := 0; i < 5; i++ {
		fmt.Println(array1[i])
	}

	var stringArray = [...]string{
		"你好",
		"世界",
		"你好",
		"你好",
		"你好",
	}
	var stringArray1 [5]string
	if len(stringArray) == len(stringArray1) {
		stringArray1 = stringArray
		for i := 0; i < 5; i++ {
			fmt.Println(stringArray1[i])
		}
	}

}
