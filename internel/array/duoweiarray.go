package array

import "fmt"

// 以下是多维数组

func MuteArray() {
	// 外层数组有4个元素 每个元素为长度为2的int型数组
	var array [4][2]int
	array = [4][2]int{1: {23, 45}, 3: {1, 1}}
	var array2 [4][2]int
	array2 = array
	fmt.Printf("多维数组array2:%d %d\n", len(array2), len(array2[0]))
	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array2[0]); j++ {
			fmt.Print(array2[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
