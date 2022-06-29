package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	// create file privilege was 0666
	f, err := os.Create("out/out.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// crgeate tar.Writer
	nw := tar.NewWriter(f)
	defer nw.Close()

	fileInfo, err := os.Stat("input/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// write file info
	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	hdr.Name = "test.txt"
	err = nw.WriteHeader(hdr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fl, err := os.Open("input/test.txt")
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

	// 下面是解压过程

}
