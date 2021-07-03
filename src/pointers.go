package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtr(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println("init: ", i)

	zeroVal(i)
	fmt.Println("zero val: ", i)

	zeroPtr(&i)
	fmt.Println("zero pinter: ", i)

	fmt.Println("addr: ", &i)
}
