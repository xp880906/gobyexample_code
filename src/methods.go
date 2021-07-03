package main

import "fmt"

type Rect struct {
	Height, Width int
}

func (r *Rect) Area() int {
	return r.Width * r.Height
}

func (r *Rect) Perim() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	r := Rect{10, 5}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perim: ", r.Perim())

	rp := &r
	fmt.Println("Area: ", rp.Area())
	fmt.Println("Perim: ", rp.Perim())
}
