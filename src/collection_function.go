package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, s := range vs {
		if s == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if f(s) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if !f(s) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	if len(vs) == 0 {
		return make([]string, 0)
	}

	ret := make([]string, 0)
	for _, s := range vs {
		if f(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func Map(vs []string, f func(string) string) []string {
	ret := make([]string, len(vs))
	for i, s := range vs {
		ret[i] = f(s)
	}
	return ret
}

func main() {
	strs := []string{"apple", "peach", "pear", "plum"}
	fmt.Println("Index: ", Index(strs, "peach"))
	fmt.Println("Exist: ", Include(strs, "grape"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))
	fmt.Println(Map(strs, func(s string) string {
		return strings.ToUpper(s)
	}))
}
