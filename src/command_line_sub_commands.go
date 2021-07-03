package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("Expected foo and bar sub commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(*fooEnable)
		fmt.Println(*fooName)
		fmt.Println(fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(*barLevel)
		fmt.Println(barCmd.Args())
	default:
		fmt.Println("Expected foo and bar sub commands")
		os.Exit(1)
	}
}
