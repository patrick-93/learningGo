package main

import (
	"fmt"
	"os"
)

var LOG_FILE string = "/tmp/app.log"

func LoggingSetup() {
	file, err := os.Open(LOG_FILE)
	if err != nil {
		fmt.Println("Error opening log file, does it exist?")
	}
	fmt.Println(file)
}
