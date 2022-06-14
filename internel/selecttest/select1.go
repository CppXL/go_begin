package selecttest

import "fmt"

func TSelect() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		chan1 <- 1
		chan2 <- 2
		close(chan1)
		close(chan2)
	}()
	for {
		select {
		case _, ok := <-chan1:
			fmt.Printf("do something")
			if ok {
				fmt.Printf("ok")
			}
		case <-chan2:
			fmt.Printf("do not do")

		}
	}

}
