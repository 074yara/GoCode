package main

import (
	"fmt"
)

func main() {

	var first, second, fibo, counter, input int
	fmt.Scan(&input)
	//mrFibo := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155}
	//input = 514229
	first = 1
	second = 1
	fibo = 0
	counter = 2

	for {
		if fibo == input {
			fmt.Println(counter)
			break
		}
		counter++
		fibo = first + second
		first = second
		second = fibo

		if fibo > input {
			fmt.Println(-1)
			break
		}
		if fibo >= 9223372036854775807 {
			fmt.Println(-1)
			break

		}

	}

}
