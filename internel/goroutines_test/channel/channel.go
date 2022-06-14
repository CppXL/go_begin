package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 声明wg
var (
	wg sync.WaitGroup
)

/*
//只能向chan里写数据
func send(c chan<- int) {
    for i := 0; i < 10; i++ {
        c <- i
    }
}
//只能取channel中的数据
func recv(c <-chan int) {
    for i := range c {
        fmt.Println(i)
    }
}
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func TChannel() {
	fmt.Println("call TChannel")
	// 无缓冲区
	//unbuffered := make(chan int)
	// 有缓冲区通道
	//buffered := make(chan string, 10)
	//buffered <- "use it"
	court := make(chan int)
	wg.Add(2)
	go player("first", court)
	go player("second", court)
	// 发球
	court <- 1
	wg.Wait()
	fmt.Println("--------------------------")

}
func player(name string, court chan int) {
	defer wg.Done()
	fmt.Printf("type of court %T\n", court)
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

func TChannel2() {
	jobs := make(chan int, 5)
	wg.Add(1)
	go readonlyChan(jobs)
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// 声明只读通道
func readonlyChan(jobs <-chan int) {
	defer wg.Done()
	fmt.Printf("type of jobs %T\n", jobs)

	for job := range jobs {
		fmt.Println("type of results ", job)
	}

}

func TSelectChannel() {
	fmt.Println("call TSelectChannel")
	job := make(chan string)
	st := make(chan int)
	wg.Add(1)
	go twork(job, st)

	job <- "1"
	job <- "2"
	st <- 1
	// st <- 2
	wg.Wait()
	fmt.Println("-------------------")
}

func twork(job <-chan string, st <-chan int) {
	defer wg.Done()
	for {
		select {
		case s, ok := <-job:
			if ok {
				fmt.Println("job:", s)
			} else {
				return
			}

		case sti, ok := <-st:
			if ok {
				if sti == 2 {
					fmt.Println("return channel")
					return
				} else if sti == 1 {
					fmt.Println("type 1")
					continue
				}
			} else {
				return
			}
		default:
			continue
		}
	}
}
