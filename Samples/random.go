package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func randVal(max int) int {
	return rand.Intn(max) + 1
}

func main() {

	fmt.Println("Now will generate a random number. Enter the upper limit: ")

	var pMax int

	fmt.Scan(&pMax)
	fmt.Printf("Random number is %v\n", randVal(pMax))

	fmt.Println("Repeat it 'n' times or exit program?")

	var input string
	fmt.Scan(&input)

	if input == "exit" {
	} else {
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("An error '%v' occured\n", err)
		} else {
			for i := 0; i <= n; i++ {
				fmt.Printf("Random number is %v\n", randVal(pMax))
			}
		}

	}

}
