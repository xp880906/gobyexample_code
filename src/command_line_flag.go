package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numPtr)
	fmt.Println(*boolPtr)
	fmt.Println(svar)
	fmt.Println(flag.Args())
}
