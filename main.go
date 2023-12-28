package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
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

var filepath = "./_index.md"

func main() {
	var timeField, counterField []byte
	var counter int
	data, err := os.ReadFile(filepath)
	checkError(err)
	newData := make([]byte, len(data))
	fields := bytes.Split(data, []byte("\n"))
	for _, field := range fields {
		if bytes.Contains(field, []byte("Текущее время:")) {
			timeField = field
		}
		if bytes.Contains(field, []byte("Счетчик:")) {
			counterField = field
		}
	}
	for {
		time.Sleep(time.Second * 5)
		counter++
		currTime := time.Now().Format("2006-01-02 15-04-05")
		newTimeField := []byte(fmt.Sprintf("Текущее время: %v", currTime))
		newCounterField := []byte(fmt.Sprintf("Счетчик: %v", counter))
		newData = bytes.Replace(data, timeField, newTimeField, 1)
		newData = bytes.Replace(newData, counterField, newCounterField, 1)
		err = os.WriteFile(filepath, newData, 0644)
		checkError(err)
	}

	/*
		for {
			time.Sleep(time.Second * 5)
			for _, field := range fields {
				if bytes.Contains(field, []byte("Текущее время:")) {
					currTime := time.Now().Format("2006-01-02 15-04-05")
					newField := []byte(fmt.Sprintf("Текущее время: %v", currTime))
					data = bytes.Replace(data, field, newField, 1)
					err = os.WriteFile(filepath, data, 0644)
				}
				if bytes.Contains(field, []byte("Счетчик:")) {
					newField := []byte(fmt.Sprintf("Счетчик: %v", counter))
					counter++
					data = bytes.Replace(data, field, newField, 1)
					err = os.WriteFile(filepath, data, 0644)
				}
			}
		}

	*/

}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
