package main

import (
	"fmt"
)

func fib(len int) chan int {
	c := make(chan int)

	go func() {
		counter := 0
		for i, j := 0, 1; counter != 100; i, j = i+j, i {
			c <- i
			counter++
		}
		close(c)

	}()
	return c
}

func main() {
	arr := []int{4, 7, 3, 3, 5, 8, 9, 0, 5, 3, 2, 5, 7}
	fmt.Println(quickSort(arr))

}

func quickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[len(data)/2]
	lowPart := make([]int, 0, len(data))
	middlePart := make([]int, 0, len(data))
	highPart := make([]int, 0, len(data))
	for _, value := range data {
		switch {
		case value < pivot:
			lowPart = append(lowPart, value)
		case value == pivot:
			middlePart = append(middlePart, value)
		case value > pivot:
			highPart = append(highPart, value)

		}
	}
	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)
	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart
}

func bubbleSort(data []int) []int {
	swapped := true
	i := 1
	for swapped {
		swapped = false
		for j := 0; j < len(data)-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true
			}
		}
		i++
	}
	return data
}

func insertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	return data
}
