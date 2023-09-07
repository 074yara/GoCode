package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Calculate struct {
	first  int
	second int
	action string
}

var (
	input  string
	temp   []string
	action string
	result int
)

func main() {
	fmt.Scan(&input)
	fmt.Println(input)
	temp = strings.Split(input, "")
	first, err := strconv.Atoi(temp[0])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(temp[2])
	if err != nil {
		log.Fatal(err)
	}
	action = temp[1]
	switch action {
	case "*":
		result = first * second
	case "/":
		result = first / second
	case "+":
		result = first + second
	case "-":
		result = first - second
	default:
		fmt.Println("Action error")
	}
	fmt.Printf("=%v\n", result)

}
