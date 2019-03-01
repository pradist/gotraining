package main

import "fmt"

func sender(ch chan int) {
	ch <- 10
}

func main() {

	//ch := make(chan int)
	//go sender(ch)
	//n := <- ch
	//fmt.Println(n)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	ch <- 4
	fmt.Println(<-ch)
	ch <- 5
	ch <- 6
}
