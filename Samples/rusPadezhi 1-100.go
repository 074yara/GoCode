package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Scan(&a)

	//a = 10
	marker := a % 100 % 10
	word := ""
	pWord := &word

	//fmt.Println(marker)

	if a >= 11 && a <= 14 {
		*pWord = "korov"
	} else {

		if marker == 0 {
			*pWord = "korov"
		}
		if marker == 1 {
			*pWord = "korova"
		}
		if marker >= 2 && marker <= 4 {
			*pWord = "korovy"
		}
		if marker >= 5 && marker <= 9 {
			*pWord = "korov"
		}
	}

	fmt.Printf("%v %v\n", a, word)

}
