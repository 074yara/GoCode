package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		sum += x

	}
	os.Stdout.WriteString(strconv.Itoa(sum))

}
