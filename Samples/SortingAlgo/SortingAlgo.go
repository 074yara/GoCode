package main

import (
	"fmt"
	"math/rand"
)

var UnsortedSlice = []int{3, 8, 7, 0, 3, 2, 5, 67, 54, 98, 78, 94, 45}

//10.29 ns/op

func BubbleSort(data []int) []int {
	i := 1
	swapped := true
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

//1132 ns/op !!! quick???

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[rand.Intn(len(data))]
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
	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)
	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart
}

func main() {
	fmt.Println(QuickSort(UnsortedSlice))
}
