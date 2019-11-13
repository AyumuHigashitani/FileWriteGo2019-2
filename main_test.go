package main

import (
	"ie.u-ryukyu.ac.jp/hg/students/y18/e185747/os/FileWriteGo2019-2/fileIO"
	"testing"
)

func Test_getOpts(t *testing.T) {
	type args struct {
		u fileIO.FileIO
	}

	// u1 := fileIO.FileIO{}
	var tests = []struct {
		name string
		args args
	}{
		{"test1", args{fileIO.FileIO{"test.txt", false, 1024, 1024}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.u = getOpts(tt.args.u.FileName, tt.args.u.Buffering, tt.args.u.BufferSize, tt.args.u.WriteSize)
			if tt.args.u.FileName == "" {
				t.Fail()
			}
		})
	}
}
