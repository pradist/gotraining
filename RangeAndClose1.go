package main

import "fmt"

func sender(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sender(ch) // start sender goroutine

	for v := range ch {
		fmt.Println(v)
	}

	//for {         // receiver
	//	v, ok := <-ch
	//	if !ok {
	//		fmt.Println("DONE")
	//		return // done
	//	}
	//	fmt.Println(v)
	//}
}
