package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	println("creating...")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	println("writing")
	_, err := fmt.Fprintln(f, "data")
	if err != nil {
		return
	}
}

func closeFile(f *os.File) {
	print("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
