package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou!")
	ch <- 4
	ch <- 5
	ch <- 6
}
