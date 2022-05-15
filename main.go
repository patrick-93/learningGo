package main

import (
	"learning/linuxclient"
	"learning/winclient"
	"runtime"
)

func main() {
	// Check OS and run specific stuff
	if runtime.GOOS == "linux" {
		linuxclient.Start()
	}
	if runtime.GOOS == "windows" {
		winclient.Start()
	}

	// Slice_test()
	// Struct_test()
	Map_test()
}
