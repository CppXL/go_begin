package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func MySplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) < 5 {
		return len(data), data, nil
	}
	return 5, data[:5], nil
}

func main() {
	data := "Hello World\nhey there\nwho are you\n"
	bdata := []byte(data)

	// ScanBytes
	// 返回一个字符
	i, scanner, err := bufio.ScanBytes(bdata, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	fmt.Println(scanner)

	// ScanLines
	i, scanner, err = bufio.ScanLines(bdata, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print `12` `hello world`
	fmt.Println(i)
	fmt.Println(string(scanner))
	i, scanner, err = bufio.ScanLines(bdata[i:], true)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print `10` `hey there`
	fmt.Println(i)
	fmt.Println(string(scanner))

	// ScanRunes
	data = "你好 世界\n这是哪里\n你是谁\n"
	bdata = []byte(data)
	i, scanner, err = bufio.ScanRunes(bdata, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	fmt.Println(scanner)
	fmt.Println(string(scanner))

	// ScanWords
	i, scanner, err = bufio.ScanWords(bdata, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print `7` `你好`
	fmt.Println(i)
	fmt.Println(string(scanner))
	fmt.Println("begin test MySplit")
	TSplitFunc()
}

func TSplitFunc() {
	scanner := bufio.NewScanner(bytes.NewReader([]byte("Hello World\nhey there\nwho are you\n")))
	// scanner.Split(MySplit)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
}
