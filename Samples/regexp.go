package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{5,}$", a)
	if matched {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
