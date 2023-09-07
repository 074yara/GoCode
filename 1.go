package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	simpleConc()
}

func simpleConc() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mx.Lock()
			counter++
			mx.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())

}
