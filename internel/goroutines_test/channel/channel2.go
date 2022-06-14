package channel

import (
	"fmt"
	"time"
)

func TChannel1() {
	fmt.Println("call TChannel1")
	baton := make(chan int)
	wg.Add(1)
	go runnerFunc(baton)
	baton <- 1
	wg.Wait()
	fmt.Println("-------------------------")
}

func runnerFunc(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d running with baton\n", runner)
	if runner != 4 {
		// 如果routine没有到达4个则新起u一个
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go runnerFunc(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d finished, race over\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("runner %d exchange with runner %d\n",
		runner, newRunner)
	baton <- newRunner
}
