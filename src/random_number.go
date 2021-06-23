package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	println()

	println(rand.Float64())
	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print(rand.Float64() * 5 + 5)
	println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	print(r1.Intn(100), ", ")
	print(r1.Intn(100))
	println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	print(r2.Intn(100), ", ")
	print(r2.Intn(100))
	println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	print(r3.Intn(100), ", ")
	print(r3.Intn(100))
	println()
}
