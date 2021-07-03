package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Printf("Write %d as ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Wednesday:
		fmt.Println("It is Wednesday")
	case time.Saturday, time.Sunday:
		fmt.Println("Today is workday")
	default:
		fmt.Println("It is weekday")
	}

	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("It's before noon")
	case now.Hour() >= 12:
		fmt.Println("It's after noon")
	}

	// Type Switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Unknown type %T \n", t)
		}
	}
	whoAmI(true)
	whoAmI(123)
	whoAmI("foo")
}
