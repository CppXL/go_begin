package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := GetMsg4Ser(1)
	fmt.Println(a, err)
}

func GetMsg4Ser(port int) (int, error) {
	return 1, errors.New("sa")
}
