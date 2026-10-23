package main

import "os"

func main() {
	environmentFor(os.Args[1:]).Start()
}
