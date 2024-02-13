package main

import (
	"fmt"
)

func test() {
	myChannel := make(chan int, 1)
	myChannel = make(chan int, 2)
	myChannel <- 10
	fmt.Println(<-myChannel)
}

