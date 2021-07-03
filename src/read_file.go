package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	fp, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := fp.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, b1[:n1])

	o2, err := fp.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := fp.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %d\n", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := fp.Seek(6, 0)
	check(err)
	b3 := make([]byte, 5)
	n3, err := io.ReadAtLeast(fp, b3, 2)
	check(err)
	fmt.Printf("%d bytes @%d: %s\n", n3, o3, string(b3))

	_, err = fp.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(fp)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	fp.Close()
}
