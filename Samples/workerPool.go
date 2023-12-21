package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	//simpleConc()

	//funcBufChan()
	//funcUnBufChan()

	//selectFunc()
	//workerPool()

	workerPool()
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

func funcBufChan() {
	start := time.Now()
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 5, 6}
	unBufChan := make(chan int, cap(numbers))

	go func() {
		for _, v := range numbers {
			unBufChan <- v
		}
		close(unBufChan)
	}()

	for v := range unBufChan {
		fmt.Print(v, " ")
	}
	fmt.Println("funcBufChan ", time.Now().Sub(start).Nanoseconds())
}

func funcUnBufChan() {
	start := time.Now()
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 5, 6}
	unBufChan := make(chan int)

	go func() {
		for _, v := range numbers {
			unBufChan <- v
		}
		close(unBufChan)
	}()

	for v := range unBufChan {
		fmt.Print(v, " ")
	}
	fmt.Print("funcUnBufChan ", time.Now().Sub(start).Nanoseconds())

}

func selectFunc() {
	timer := time.After(time.Second)
	resChan := make(chan int)

	go func() {
		defer close(resChan)
		for i := 1; i <= 1000; i++ {
			select {
			case <-timer:
				fmt.Println("Time's up")
				return
			default:
				time.Sleep(time.Millisecond)
				resChan <- i
			}

		}
	}()
	for v := range resChan {
		fmt.Println(v)
	}
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- v * v

		}
	}
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	toWorkerChan, fromWorkerChan := make(chan int), make(chan int)
	wg := &sync.WaitGroup{}
	counter := 0

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, toWorkerChan, fromWorkerChan)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			toWorkerChan <- i
		}
		close(toWorkerChan)
	}()

	go func() {
		wg.Wait()
		close(fromWorkerChan)
	}()

	for returnedValue := range fromWorkerChan {
		fmt.Println(returnedValue)
		counter++
	}

	fmt.Println(counter)
}
