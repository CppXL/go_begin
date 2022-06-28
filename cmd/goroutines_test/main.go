package main

import (
	"fmt"
	"gobegin/internel/goroutines_test"
	"gobegin/internel/goroutines_test/locker"
	"gobegin/internel/goroutines_test/runner_test"
	"runtime"
)

func init() {

}
func main() {
	// 设置逻辑处理器数量
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	fmt.Printf("use %d cpu\n", runtime.NumCPU())
	goroutines_test.TGoroutine()
	goroutines_test.TGoroutine1()
	locker.TLocker1()
	locker.TMutex()

	runner_test.TRunner()

}
