package main

import "fmt"

func main() {
	var x, y int
out:
	for x = 0; x < 10; x++ {
		for y = 0; y <= x; y++ {
			fmt.Printf("%d * %d = %d ", x, y, x*y)
			if y == 6 {
				continue out
			}
		}
		fmt.Println()
	}
}
