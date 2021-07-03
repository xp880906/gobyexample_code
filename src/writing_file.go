package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Write %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)
	f.Sync()

	writer := bufio.NewWriter(f)
	n4, err := writer.WriteString("buffered\n")
	check(err)
	fmt.Printf("Write %d bytes\n", n4)

	writer.Flush()
}
