package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func main() {
	n := 50

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

//это не экспериментальное решение, а применение паттерна Future.
//Функция запускает свое вычисление в фоне и возвращает канал, в который отправит результат по завершении.
//Чтобы было понятнее, это можно реализовать так:

type Worker struct {
	sync.Mutex
	sync.WaitGroup
}

func (w *Worker) Do(work func(int) int, x int, res *int) {
	w.Add(1)
	go func() {
		y := work(x)
		w.Lock()
		*res += y
		w.Unlock()
		w.Done()
	}()
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		result := make([]int, n)
		var w Worker
		for i := range result {
			w.Do(fn, <-in1, &result[i])
			w.Do(fn, <-in2, &result[i])
		}
		w.Wait()
		for i := range result {
			out <- result[i]
		}
	}()
}
