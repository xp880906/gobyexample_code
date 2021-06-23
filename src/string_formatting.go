package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}
	fmt.Printf("%v\n", p)
	// With Struct Members' Names
	fmt.Printf("%+v\n", p)
	// Include Struct Namespace
	fmt.Printf("%#v\n", p)
	// Print Type Of Value
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", false)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	// Hex Format
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
	// Align
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 12.0, 345.0)
	// Left Padding
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.0, 345.0)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
