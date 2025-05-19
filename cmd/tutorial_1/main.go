package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
