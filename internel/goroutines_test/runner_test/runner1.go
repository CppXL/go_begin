package runner_test

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type runner struct {
	tasks []func(int) string
}

func TRunner() {
	fmt.Println("call TRunner")
	a := runner{tasks: []func(int) string{test}}
	b := a.tasks[0](1)
	fmt.Printf("%v\n", b)
	fmt.Println("-----------------------------")
}
func test(a int) string {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("aaaaaaaaaa ", a)
	}()
	wg.Wait()
	return fmt.Sprintf("hello world %d", a)
}
