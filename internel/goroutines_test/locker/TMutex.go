package locker

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func TMutex() {
	fmt.Println("call TMutex")
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("final counter:", counter)
	fmt.Println("=======================")
}

func incCounter(id int) {
	defer func() {
		fmt.Println("----------------------")
		wg.Done()
	}()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
