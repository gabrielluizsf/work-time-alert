package main

import (
	"fmt"
	"runtime"

	"github.com/i9si-sistemas/command"
)

func main() {
	binaryName := "work-time-alert"
	if runtime.GOOS == "windows" {
		binaryName = "work-time-alert.exe"
	}
	output, err := command.New().Execute("./"+binaryName).CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}