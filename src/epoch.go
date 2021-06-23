package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nano := now.UnixNano()
	millis := nano / 1000000
	fmt.Printf("now: %v, secs: %v, nano: %v, mills: %v\n", now, secs, nano, millis)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nano))
}
