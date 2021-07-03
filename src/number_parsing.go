package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	println(u)

	k, _ := strconv.Atoi("135")
	println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
