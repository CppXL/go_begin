package channel

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

func TBBufferedChannel() {
	fmt.Println("call TBBufferedChannel")
	// 新建带缓冲通道
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post < taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	close(tasks)
	wg.Wait()
	fmt.Println("-----------------------------")
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker: %d : shutting down\n", worker)
			return
		}
		fmt.Printf("worker: %d: started %s\n", worker, task)
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		fmt.Printf("worker: %d : completed %s\n", worker, task)
	}
}
