package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Print(value)
	case value := <-stringChan:
		panic(value)
	}
}
