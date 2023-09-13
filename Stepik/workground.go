package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rawString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic("error")
	}

	rawString = strings.ReplaceAll(rawString, ",", ".")
	rawString = strings.ReplaceAll(rawString, " ", "")
	rawString = strings.ReplaceAll(rawString, "\n", "")
	rawString = strings.ReplaceAll(rawString, "\r", "")
	cleanSlice := strings.Split(rawString, ";")
	cleanSlice[0] = strings.ReplaceAll(cleanSlice[0], ";", "")

	x, _ := strconv.ParseFloat(cleanSlice[0], 64)
	y, _ := strconv.ParseFloat(cleanSlice[1], 64)

	fmt.Printf("%.4f", x/y)

}
