package main

import "fmt"

func main() {
	c := make(chan int, 5)
	c <- 1
	fmt.Print(len(c))
}
