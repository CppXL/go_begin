package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建文件
	f, err := os.Create("./out.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 创建 tar.Writer
	nw := tar.NewWriter(f)
	defer nw.Close()

	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 写入文件信息
	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	hdr.Name = "test.txt"
	err = nw.WriteHeader(hdr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fl, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()
	_, err = io.Copy(nw, fl)
	if err != nil {
		fmt.Println(err)
		return
	}

}
