package slice_test

import (
	"fmt"
)

func TSlice() {
	fmt.Println("call slice_test.TSlice")
	// 通过make创造一个切片 slice
	// 前一个是长度，后一个是容量
	slice1 := make([]int, 4, 6)

	//通过切片字面量声明切片，长度容量都是4
	slice2 := []int{1, 2, 3, 4}
	slice1 = slice2

	// 在[]里面指定值就是数组，不指定值就是切片
	slice3 := []string{99: ""}

	// 声明 nil int型切片
	var slice4nil []int

	// 声明空切片
	slice4empty := make([]int, 0)
	//slice :=[]int{}
	fmt.Printf("slice1 len %d\n", len(slice1))
	fmt.Printf("slice3 len %d\n", len(slice3))
	// slice4长度为2-1=1 容量为4-1=3
	slice4 := slice2[3:]
	//slice2 {1, 2, 3, 4}
	fmt.Printf("slice2 len %d slice4 len %d\n", len(slice2), len(slice4))
	fmt.Printf("slice4=>%v\nslice2=>%v\n", slice4, slice2)
	fmt.Println(len(slice4empty), len(slice4nil))
	slice4empty = slice4
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice4empty = a[1:4]
	s := struct {
		a int
		b int
	}{1, 2}
	fmt.Println(s)
	// 对slice4赋值
	slice4[0] = 6
	fmt.Printf("slice2[1]=>%d\n", slice2[1])
	fmt.Printf("修改前：slice2[2]=>%d\n", slice2[2])
	slice4 = append(slice4, 7)
	// 此时修改了slice2[2]
	fmt.Printf("修改后：slice2[2]=>%d\n", slice2[2])

	fmt.Printf("修改前：slice2[3]=>%d\n", slice2[3])
	slice4 = append(slice4, 8)
	fmt.Printf("修改后：slice2[3]=>%d\n", slice2[3])

	// 此时append数据时 slice4已达容量上限，自增底层数组，和原来的数组不一样了
	slice4 = append(slice4, 9)
	fmt.Printf("修改前：slice2[1]=>%d\n", slice2[1])
	slice4[0] = 20
	fmt.Printf("修改后：slice2[1]=>%d\n", slice2[1])

	// 以下为实验第三个参数
	sourceColor := []string{"red", "blue", "green", "yellow", "black"}
	// 长度和容量都是2
	newColor := sourceColor[1:3:3]
	fmt.Printf("new color len %d\n", len(newColor))
	fmt.Printf("sourceColor[1]=>%v newColor[0]=>%v\n", sourceColor[1], newColor[0])
	fmt.Printf("sourceColor[2]=>%v newColor[1]=>%v\n", sourceColor[2], newColor[1])
	newColor = append(newColor, "write")
	// 由于newColor长度容量一样，当append时会创建新的切片给newColor，sourceColor不受影响
	fmt.Printf("sourceColor[3]=>%v newColor[2]=>%v\n", sourceColor[3], newColor[2])
	secColor := []string{"color"}
	sourceColor = append(sourceColor, secColor...)
	fmt.Println(sourceColor, "\n--------------------------------")
}

func TSlice2() {
	fmt.Println("call slice_test.TSlice2")
	sourceColor := []string{"red", "blue", "green", "yellow", "black"}
	for index, value := range sourceColor {
		fmt.Printf("Index %d value %s\n", index, value)
	}
	fmt.Println("\n----------------------------")
}

func TMultiSlice() {
	fmt.Println("call slice_test.TMultiSlice")
	slice1 := [][]int{{10}, {100, 200}}
	fmt.Printf("slice1 len %d, %d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, make([]int, 2, 2))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d] len %d cap %d\t", i, len(slice1[i]), cap(slice1[i]))
		for j := 0; j < len(slice1[i]); j++ {
			fmt.Printf("slice[%d][%d] value %d\t", i, j, slice1[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Println("\n----------------------------")
	slice2 := tFuncSlice2Func(slice1[2])
	fmt.Printf("tFuncSlice2Func return slice len %d cap %d\n", len(slice2), cap(slice2))
}

func tFuncSlice2Func(slice []int) []int {
	fmt.Println("\ncall slice_test.tFuncSlice2Func")

	fmt.Println("---------------------------")

	return slice
}
