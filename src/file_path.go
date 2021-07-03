package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

const filename = "config.json"

func main() {
	p := filepath.Join("foo", "bar", filename)
	fmt.Println(p)

	fmt.Println(filepath.Join("dir//", filename))
	fmt.Println(filepath.Join("dir1/.../dir1", filename))

	fmt.Println("Dir p: ", filepath.Dir(p))
	fmt.Println("Base p: ", filepath.Base(p))

	fmt.Println("is abs: ", filepath.IsAbs("foo/bar"))
	fmt.Println("is abs: ", filepath.IsAbs("/foo/bar"))

	ext := filepath.Ext(filename)
	fmt.Println("Ext: ", ext)
	fmt.Println("Trim Suffix: ", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("foo/bar", "foo/bar/buzz/filename")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("foo/bar", "foo/buzz/bar/filename")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
