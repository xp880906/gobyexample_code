package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutext = &sync.Mutex{}
	var readOps, writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutext.Lock()
				total += state[key]
				mutext.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutext.Lock()
				state[key] = val
				mutext.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("Read Ops: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("Write Ops: ", writeOpsFinal)

	mutext.Lock()
	fmt.Println("state: ", state)
	mutext.Unlock()
}
