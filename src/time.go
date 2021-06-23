package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())

	then := time.Date(2021, 6, 24, 20, 30, 0, 0, time.UTC)
	p(then)
	p(then.Location())
	p(then.Local())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := then.Sub(now)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}
