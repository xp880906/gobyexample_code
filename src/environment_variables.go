package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("foo", "1")
	fmt.Println("foo: ", os.Getenv("foo"))
	fmt.Println("bar: ", os.Getenv("bar"))
	println()

	for _, e := range os.Environ() {
		//pair := strings.SplitN(e, "=", 2)
		fmt.Println(e)
	}
}
