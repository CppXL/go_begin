package locker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg       sync.WaitGroup
	shutdown int64
)

func TLocker1() {
	fmt.Println("call TLocker1")
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
	fmt.Println("--------------------------")
}
func doWork(name string) {
	defer func() {
		wg.Done()
	}()
	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s down\n", name)
			break
		}
	}
}
