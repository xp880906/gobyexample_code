package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"b", "c", "a"}
	sort.Strings(strs)
	fmt.Println("string: ", strs)

	ints := []int{1, 3, 5, 4}
	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", sorted)
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	sorted = sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", sorted)
}
