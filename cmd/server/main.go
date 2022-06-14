package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strconv"
	"sync"
)

var ConnLi []net.Conn
var ConnStat []int
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("Begian Listen\n")
	ln, err := net.Listen("tcp", ":9898")
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go AcCliConn(ln)
	wg.Wait()
}

func StoreConn(conn net.Conn, connLi []net.Conn) []net.Conn {
	connLi = append(connLi, conn)
	return connLi
}
func AcCliConn(ln net.Listener) {
	defer wg.Done()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("Accept Conn from Cline %v\n", conn.RemoteAddr())
		// 存储连接
		ConnLi = StoreConn(conn, ConnLi)
		wg.Add(1)
		// 处理和clint的连接

		for {
			var code int
			fmt.Print("控制Client,输入操作码\n1. 展示所有连接\t2. 控制连接\t3. 退出\n输入:")
			fmt.Scan(&code)
			switch code {
			case 1:
				ShowAllCli(ConnLi)
				continue
			case 2:
				ShowAllCli(ConnLi)
				var s int
				fmt.Print("输入控制那个client:")
				fmt.Scan(&s)
				Snd2Cli(ConnLi[s])
				continue
			}
		}

	}
}
func Snd2Cli(conn net.Conn) {
	var code int
	fmt.Print("输入命令:")
	fmt.Scan(&code)
	conn.Write([]byte(strconv.Itoa(code)))
	cmdrt := make([]byte, 0)
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadByte()
		if err == io.EOF {
			break
		}

		cmdrt = append(cmdrt, data)
		log.Printf("Client执行结果:%v", cmdrt)
	}

}
func ShowAllCli(Connlist []net.Conn) {
	for idx, conn := range Connlist {
		log.Printf("No %d Conn %v\n", idx+1, conn.RemoteAddr())
	}
}
