package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func main() {
	n := 5

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 100
	}
	in1 := make(chan int, n)
	in2 := make(chan int, n)
	out := make(chan int, n)

	merge2Channels(fn, in1, in2, out, n)

	for i := 0; i < n+1; i++ {
		in1 <- i * 5
		in2 <- i * 10
	}

	for i := 0; i < n; i++ {
		fmt.Println("out", <-out)

	}

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		wg := sync.WaitGroup{}
		arr1 := make([]int, n)
		arr2 := make([]int, n)

		for i := 0; i < n; i++ {
			wg.Add(2)
			x := <-in1
			y := <-in2
			go func(i int) {
				arr1[i] = fn(x)
				defer wg.Done()
			}(i)
			go func(i int) {
				arr2[i] = fn(y)
				defer wg.Done()
			}(i)
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- arr1[i] + arr2[i]
		}

	}()

}
