package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (l ByLength) Len() int {
	return len(l)
}

func (l ByLength) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByLength) Less(i, j int) bool {
	return len(l[i]) < len(l[j])
}

func main() {
	strings := []string{"peach", "banana", "kiwi"}
	byLength := ByLength(strings)
	sort.Sort(byLength)
	fmt.Println("Sort by function: ", byLength)
}
