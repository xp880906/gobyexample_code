package main

import "os"

func main() {
	panic("something wrong")

	_, err := os.Create("tmp/file")
	if err != nil {
		panic(err)
	}
}
