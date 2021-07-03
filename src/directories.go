package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

const BASEDIR = "subdir"

func main() {
	err := os.Mkdir(BASEDIR, 0755)
	check(err)
	defer os.RemoveAll(BASEDIR)

	createEmptyFile := func (filename string) {
		d := []byte("")
		e := ioutil.WriteFile(filename, d, 0644)
		check(e)
	}

	createEmptyFile(filepath.Join(BASEDIR, "file1"))
	err = os.MkdirAll(filepath.Join(BASEDIR, "parent/child"), 0755)
	check(err)

	createEmptyFile(filepath.Join(BASEDIR, "parent/file2"))
	createEmptyFile(filepath.Join(BASEDIR, "parent/file3"))
	createEmptyFile(filepath.Join(BASEDIR, "parent/child/file4"))

	c, err := ioutil.ReadDir(filepath.Join(BASEDIR, "parent"))
	check(err)
	fmt.Println("Listing subdir/parent Dir")
	for _, entry := range c {
		fmt.Println("", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(filepath.Join(BASEDIR, "parent", "child"))
	check(err)

	c, err = ioutil.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child dir: ")
	for _, entry := range c {
		fmt.Println("", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk(BASEDIR, visit)
}

func visit(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println("", path, info.IsDir())
	return nil
}
