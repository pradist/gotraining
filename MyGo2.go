package main

import "fmt"

func filterEven(nums []int) []int {
	var results []int
	for _, v := range nums {
		if v%2 == 0 {
			results = append(results, v)
		}
	}
	return results
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(input)
	fmt.Println(filterEven(input))
}
