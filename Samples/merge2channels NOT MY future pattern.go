package main

import (
	"fmt"
	"math/rand"
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

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	futureIn1, futureIn2 := mapFuture(fn, in1, n), mapFuture(fn, in2, n)
	go func() {
		for i := 0; i < n; i++ {
			out <- <-<-futureIn1 + <-<-futureIn2
		}
	}()
}

// See http://www.golangpatterns.info/concurrency/futures
func future(fn func(int) int, arg int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		out <- fn(arg)
	}()
	return out
}

func mapFuture(fn func(int) int, in <-chan int, size int) <-chan (<-chan int) {
	out := make(chan (<-chan int), size)
	go func() {
		defer close(out)
		for x := range in {
			out <- future(fn, x)
		}
	}()
	return out
}
