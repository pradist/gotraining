package main

import "fmt"

func pyramid(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func main() {
	pyramid(5)
}
