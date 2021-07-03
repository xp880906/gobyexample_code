package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println("nums: ", nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum: ", sum)
	return sum
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
