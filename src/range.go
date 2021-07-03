package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("SUM: ", sum)

	for i, n := range nums {
		fmt.Printf("index: %d n: %d\r\n", i, n)
	}

	fruits := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range fruits {
		fmt.Printf("key: %s -> value: %s\r\n", k, v)
	}

	for k := range fruits {
		fmt.Println("k: ", k)
	}

	for i, c := range "go" {
		fmt.Printf("i: %d -> char: %c\r\n", i, c)
	}
}
