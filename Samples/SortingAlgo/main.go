package main

import (
	"fmt"
	"github.com/psilva261/timsort"
	"math/rand"
)

/*
arrLen	QuickSort 	SelectionSort	MergeSort	InsertionSort №2  BubbleSort №1	TimSort №3
50			5971		1542		2477			59.63			35.44		231.5	ns/op
100			11925		5696		5467			109.3			76.96		383.9	ns/op
500			46617		130053		36526			513.0			354.5		1533	ns/op
1500		115018		1148403		141663			1532			1060		3717	ns/op
3000		216638		4583800		280097			3048			2104		7111	ns/op
6000		434250		18364097	621912			6132			4284		13119	ns/op
12000		779053		73669319	1172176			12439			9484		25693	ns/op
24000		1738782		297597418	2438093			26008			43361		51308	ns/op
48000		3967196		1172812056	5279879			68373		2758801416		102150	ns/op
96000		7590493		4667097186	11548153	1224879583		11462669266		202684	ns/op
192000		13681219	18725195051	24839583	4948730349		46203898869		407457	ns/op
384000		25914397	74606210466	48937873	19957275201		185657256526	824739	ns/op

goos: linux
goarch: amd64
pkg: studentgit.kata.academy/074yara/go-kata/course2/4.algo_datastruct/2.algo_sort/task2.4.2.2
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz


*/

//на основании тестов с числами в массиве 0-100

func GeneralSort(data []int) {
	switch {
	case len(data) <= 12000:
		data = bubbleSort(data)
	case len(data) > 12000 && len(data) <= 48000:
		data = insertionSort(data)
	default:
		err := timsort.Ints(data, func(a, b int) bool {
			return a < b
		})
		if err != nil {
			fmt.Println(err)
		}
	}

}

func quickSort(data []int) []int {
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
	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)
	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart

}

func selectionSort(data []int) []int {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}

	return data
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := mergeSort(data[mid:])
	right := mergeSort(data[:mid])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result

}

func insertionSort(data []int) []int {
	n := len(data)
	for i := 1; i < n; i++ {
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

func bubbleSort(data []int) []int {
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

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)

	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quickSort(data)
	fmt.Println("Sorted by Quicksort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", data)
}
