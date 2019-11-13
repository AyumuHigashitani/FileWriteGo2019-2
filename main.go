package main

import (
	"flag"
	"fmt"
	"ie.u-ryukyu.ac.jp/hg/students/y18/e185747/os/FileWriteGo2019-2/fileIO"
)

func main() {
	u := fileIO.FileIO{} //変数の初期化の省略型

	u = getOpts("testDate", false, 1024, 10024)
	u.FileWrite()
	fmt.Println("Hello, World")
}

func getOpts(filename string, buffering bool, buffer_size int, write_size int) fileIO.FileIO {
	u1 := fileIO.FileIO{}
	fmt.Println("Hello, getOpts")
	flag.StringVar(&u1.FileName, "filename", filename, "file name to write")
	flag.BoolVar(&u1.Buffering, "b", buffering, "buffering")
	flag.IntVar(&u1.BufferSize, "buffersize", buffer_size, "buffer size")
	flag.IntVar(&u1.WriteSize, "writesize", write_size, "write size")
	flag.Parse()
	return u1
}
