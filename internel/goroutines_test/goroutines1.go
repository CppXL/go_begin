package goroutines_test

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func TGoroutine() {
	fmt.Println("call TGoroutine")
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			fmt.Printf("%c ", 'a'+i)
		}
		println("\n")
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			fmt.Printf("%c ", 'A'+i)
		}
		println("\n")
	}()

	wg.Wait()
	fmt.Println("-------------------------")
}

func TGoroutine1() {
	fmt.Println("call TGoroutine1")
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("final counter:", counter)
	fmt.Println("-------------------------")
}
func incCounter(id int) {
	defer func() {
		wg.Done()
		fmt.Println("-------------------------")
	}()
	fmt.Println("call incCounter")
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}
