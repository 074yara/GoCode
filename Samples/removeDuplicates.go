package main

import "fmt"

func main() {

	str := "1113313444555"
	inputPipe := make(chan string)
	outputPipe := make(chan string)

	go removeDuplicates(inputPipe, outputPipe)

	go func() {
		defer close(inputPipe)
		for _, v := range str {
			inputPipe <- string(v)
		}
	}()

	for x := range outputPipe {
		fmt.Print(x)
	}

}
func removeDuplicates(inputStream, outputStream chan string) {
	var oldStr, newStr string
	for value := range inputStream {
		newStr = value
		if newStr != oldStr {
			oldStr = newStr
			outputStream <- newStr
		}
	}
	close(outputStream)

}
