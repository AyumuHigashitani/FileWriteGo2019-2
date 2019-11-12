package main

import (
	"flag"
	"fmt"
	"ie.u-ryukyu.ac.jp/hg/students/y18/e185747/os/FileWriteGo2019-2/fileIO"
)

func main() {
	u := fileIO.FileIO{} //変数の初期化の省略型

	u = getOpts(u)
	u.FileWrite()
	fmt.Println("Hello, World")
}

func getOpts(u fileIO.FileIO) fileIO.FileIO {
	fmt.Println("Hello, getOpts")
	flag.BoolVar(&u.Buffering, "b", false, "buffering")
	flag.StringVar(&u.FileName, "filename", "testData.txt", "file name to write")
	flag.IntVar(&u.BufferSize, "buffersize", 1024, "buffer size")
	flag.IntVar(&u.WriteSize, "writesize", 102400, "write size")
	flag.Parse()
	return u
}
