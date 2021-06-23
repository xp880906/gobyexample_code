package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)
	fmt.Println("Temporary File Name: ", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dName, err := ioutil.TempDir("", "sample_dir")
	check(err)
	fmt.Println("Temporary Dir Name: ", dName)

	defer os.Remove(dName)

	fName := filepath.Join(dName, "file1")
	err = ioutil.WriteFile(fName, []byte{1, 2}, 0666)
	check(err)
}
