package fileIO

import (
	"bufio"
	"fmt"
	"os"
)

type FileIO struct {
	FileName   string
	Buffering  bool
	BufferSize int
	WriteSize  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (u *FileIO) FileWrite() {
	fmt.Println("Hello, FileWrite " + u.FileName)
	var f, err = os.Create(u.FileName)
	check(err)
	d1 := []byte("b")

	if !u.Buffering {
		u.BufferSize = 1
	}
	w := bufio.NewWriterSize(f, u.BufferSize)
	for i := 0; i < u.WriteSize; i++ {
		w.Write(d1)
	}
	w.Flush()
	defer f.Close()

}
