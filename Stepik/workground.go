package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("task.data")
	reader := bufio.NewReader(file)
	defer file.Close()
	counter := 1
	for {
		data, _ := reader.ReadString(';')
		if data == "0;" {
			fmt.Println(counter)
		} else {
			counter++
		}
	}
}
