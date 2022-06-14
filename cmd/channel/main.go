package main

import (
	"fmt"
	"project1/internel/goroutines_test/channel"
)

func main() {
	fmt.Println("开始测试通道")
	channel.TChannel()
	channel.TChannel1()
	channel.TBBufferedChannel()
	channel.TChannel2()
	channel.TSelectChannel()
}
