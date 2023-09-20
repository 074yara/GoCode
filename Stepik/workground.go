package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func main() {
	z := 5

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 100
	}
	in1 := make(chan int, z)
	in2 := make(chan int, z)
	out := make(chan int, z)

	go merge2Channels(fn, in1, in2, out, z)

	for i := 0; i < z; i++ {
		in1 <- i
		in2 <- i + 10
	}

	for i := 0; i < z; i++ {
		fmt.Println(<-out)

	}

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	wg := sync.WaitGroup{}
	pipe1 := make(chan int)
	pipe2 := make(chan int)

	go func() {
		for x := range in1 {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				pipe1 <- fn(x)
			}(x)
		}
	}()
	go func() {
		for y := range in2 {
			wg.Add(1)
			go func(y int) {
				defer wg.Done()
				pipe2 <- fn(y)
			}(y)
		}
	}()
	wg.Wait()
	go func() {
		for x := range pipe1 {
			for y := range pipe2 {
				out <- x + y
			}
		}
	}()
}
