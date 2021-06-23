package main

import "strings"

func main() {
	println("Contains: ", strings.Contains("test", "es"))
	println("Count: ", strings.Count("test", "t"))
	println("HasPrefix: ", strings.HasPrefix("test", "te"))
	println("HasSuffix: ", strings.HasSuffix("test", "st"))
	println("Index: ", strings.Index("test", "e"))
	println("Join: ", strings.Join([]string{"a", "b"}, "-"))
	println("Repeat: ", strings.Repeat("a", 5))
	println("Replace: ", strings.Replace("foo", "o", "0", -1))
	println("Replace: ", strings.Replace("foo", "o", "0", 1))
	println("Split: ", strings.Split("a-b-c-d-e", "-"))
	println("ToUpper: ", strings.ToUpper("test"))
	println("ToLower: ", strings.ToLower("TEST"))
	println()
	println("Len: ", len("hello"))
	println("Char: ", "hello"[1])
}
