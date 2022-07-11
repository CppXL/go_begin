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
	// set entry name
	hdr.Name = "test.txt"
	err = nw.WriteHeader(hdr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// open file for reading
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
	fmt.Printf("%+v\n\n", hdr)

	// write second file
	fileInfo, err = os.Stat("input/test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// write file info
	hdr, err = tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	// set entry name
	// hdr.Name = "test2.txt"
	err = nw.WriteHeader(hdr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// open file for reading
	fl, err = os.Open("input/test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(nw, fl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n\n", hdr)
	// below was extra code
	// open tar file
	f, err = os.Open("input/out.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	// create tar.Reader
	r := tar.NewReader(f)
	// loop for reading tar file
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		// open file for writing
		fl, err := os.Create("out/extra/" + hdr.Name)
		if err != nil {
			fmt.Println("create error " + err.Error())
			return
		}
		defer fl.Close()

		// copy file from tar to out folder
		_, err = io.Copy(fl, r)
		if err != nil {
			fmt.Println("copy error " + err.Error())
			return
		}
		fmt.Printf("%+v\n", hdr.FileInfo().Name())

	}
}
