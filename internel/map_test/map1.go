package map_test

import "fmt"

func TMap() map[string]string {
	fmt.Println("call map_test.TMap")
	// make创建
	dict := make(map[string]int)
	dict2 := make(map[int]string)

	// 字面量创建
	dict1 := map[string]string{"red": "#ddd", "blue": "#eee", "orange": "#aaa"}

	// 声明nil型映射
	var nilmap map[string]string

	// 给映射赋值
	dict1["black"] = "#uuu"
	dict2[1] = "d"
	dict["red"] = 1

	// 查找映射中的键值
	_, exist := dict1["test"]
	if exist {
		fmt.Println("map dict1[\"test\"] exist")
	} else {
		fmt.Println("map dict1[\"test\"] doesn't exist")
	}
	// 两种方法都可以
	value := dict1["test"]
	if value != "" {
		fmt.Println("map dict1[\"test\"] exist")
	} else {
		fmt.Println("map dict1[\"test\"] doesn't exist")
	}
	// 使用映射
	fmt.Printf("before delete dict1[\"black\"]:%s\n", dict1["black"])
	delete(dict1, "black")
	fmt.Printf("after delete dict1[\"black\"]:%s\n", dict1["black"])

	fmt.Println("----------------------------------------")
	fmt.Printf("before delete dict1[\"red\"]:%s\n", dict1["red"])
	removeColor(dict1, "red")
	fmt.Printf("after delete dict1[\"red\"]:%s\n", dict1["red"])
	return nilmap
}

func removeColor(colors map[string]string, key string) {
	_, exist := colors[key]
	if exist {
		delete(colors, key)
		fmt.Printf("删除%s\n", key)
	} else {
		fmt.Printf("未找到%s\n", key)
	}
}
