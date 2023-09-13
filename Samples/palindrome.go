package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = clearString(text)
	if len([]rune(text))%2 == 0 {
		leftString := text[:len(text)/2]
		rightString := text[len(text)/2:]
		rightStringSlice := []rune(rightString)
		for i, j := 0, len(rightStringSlice)-1; i < j; i, j = i+1, j-1 {
			rightStringSlice[i], rightStringSlice[j] = rightStringSlice[j], rightStringSlice[i]
		}
		if leftString == string(rightStringSlice) {
			fmt.Println("Палиндром")
		} else {
			fmt.Println("Нет")
		}

	}
	if len([]rune(text))%2 != 0 {
		leftString := text[:len(text)/2+1]
		rightString := text[len(text)/2-1:]
		rightStringSlice := []rune(rightString)
		for i, j := 0, len(rightStringSlice)-1; i < j; i, j = i+1, j-1 {
			rightStringSlice[i], rightStringSlice[j] = rightStringSlice[j], rightStringSlice[i]
		}
		if leftString == string(rightStringSlice) {
			fmt.Println("Палиндром")
		} else {
			fmt.Println("Нет")
		}
	}
}
func clearString(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ToLower(s)
	return s
}
