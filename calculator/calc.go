package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	str     string
	isRoman bool
)

func main() {

	pIsRoman := &isRoman
	fmt.Println("Enter data:")

	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Input Error")
	}

	str = strings.ReplaceAll(str, "\r\n", "")

	_, err = strconv.Atoi(string(str[0]))
	if err == nil {
		*pIsRoman = false
		calculate(arabDataWorker(str))
	} else {
		*pIsRoman = true
		calculate(romanDataWorker(str))
	}

}

func arabDataWorker(s string) (int, int, string) {
	var leftSideIntSlice, rightSideIntSlice []int
	var act, leftNumString, rightNumString string
	trimedSting := strings.ReplaceAll(s, " ", "")
	allStringSlice := strings.Split(trimedSting, "")
	sliceLabel := 0
	pSliceLabel := &sliceLabel

	for _, v := range allStringSlice {

		i, err := strconv.Atoi(v)
		if err == nil && *pSliceLabel == 0 {
			leftSideIntSlice = append(leftSideIntSlice, i)
		} else if err != nil {
			switch v {
			case "*", "/", "+", "-":
				pAct := &act
				*pAct = v
				*pSliceLabel = 1
			case ".", ",":
				fmt.Println("Only integers please!")
				os.Exit(2)
			default:
				continue
			}
		}

		i, err = strconv.Atoi(v)
		if err == nil && *pSliceLabel == 1 {
			rightSideIntSlice = append(rightSideIntSlice, i)
		}

	}

	for _, v := range leftSideIntSlice {
		pLeft := &leftNumString
		*pLeft = *pLeft + strconv.Itoa(v)
	}

	for _, v := range rightSideIntSlice {
		pRight := &rightNumString
		*pRight = *pRight + strconv.Itoa(v)
	}

	leftInt, err := strconv.Atoi(leftNumString)
	if err != nil {
		fmt.Println("Incorrect data type")
		os.Exit(2)
	}

	rightInt, err := strconv.Atoi(rightNumString)
	if err != nil {
		fmt.Println("Incorrect data type")
		os.Exit(2)
	}

	return leftInt, rightInt, act

}

func romanDataWorker(s string) (int, int, string) {

	var leftNumString, rightNumString, act string

	trimedString := strings.ReplaceAll(s, " ", "")
	trimedString = strings.ReplaceAll(s, "\n", "")
	trimedString = strings.ToUpper(trimedString)

	allStringSlice := strings.Split(trimedString, "")
	sliceLabel := 0

	var leftStringSlice, rightStringSlice []string

	for _, v := range allStringSlice {
		switch v {
		case "*", "/", "+", "-":
			pAct := &act
			*pAct = v
			sliceLabel = 1
			continue
		case ".", ",":
			fmt.Println("Only integers please!")
			os.Exit(2)

		}
		if sliceLabel == 0 {
			leftStringSlice = append(leftStringSlice, v)
		}

		if sliceLabel == 1 {
			rightStringSlice = append(rightStringSlice, v)
		}

	}

	for _, v := range leftStringSlice {
		leftNumString = leftNumString + v
	}

	for _, v := range rightStringSlice {
		rightNumString = rightNumString + v
	}

	leftInt, rightInt := romeSwitch(leftNumString, rightNumString)

	return leftInt, rightInt, act
}

func calculate(leftInt, rightInt int, act string) {
	var result int
	if leftInt > 10 || rightInt > 10 {
		fmt.Printf("Sorry! The condition is from -10 to 10. You entered %v and %v\n", leftInt, rightInt)
		os.Exit(2)
	}

	switch act {
	case "+":
		pResult := &result
		*pResult = leftInt + rightInt
	case "-":
		pResult := &result
		*pResult = leftInt - rightInt
	case "*":
		pResult := &result
		*pResult = leftInt * rightInt
	case "/":
		if rightInt == 0 {
			fmt.Println("Devider should not be a zero value")
			os.Exit(2)
		}
		pResult := &result
		*pResult = leftInt / rightInt

	default:
		fmt.Println("Somthing's wrong")

	}
	if isRoman {
		if result <= 0 {
			fmt.Printf("The result is %v and it's < 1, but there're no negative or zero numbers in roman numeral system!\n", result)
			os.Exit(2)
		}
		fmt.Printf("================\n%v %v %v = %v\n================\n", intToRome(leftInt), act, intToRome(rightInt), intToRome(result))
	} else {
		fmt.Printf("================\n%v %v %v = %v\n================\n", leftInt, act, rightInt, result)
	}
}

func romeSwitch(l, r string) (int, int) {
	var x, y int

	switch l {
	case "I":
		{
			x = 1
		}
	case "II":
		{
			x = 2
		}
	case "III":
		{
			x = 3
		}
	case "IV":
		{
			x = 4
		}
	case "V":
		{
			x = 5
		}
	case "VI":
		{
			x = 6
		}
	case "VII":
		{
			x = 7
		}
	case "VIII":
		{
			x = 8
		}
	case "IX":
		{
			x = 9
		}
	case "X":
		{
			x = 10
		}
	default:
		{
			fmt.Printf("%v is not a Roman digit or too many roman digits in a string\n", l)
			os.Exit(2)
		}
	}

	switch r {
	case "I":
		{
			y = 1
		}
	case "II":
		{
			y = 2
		}
	case "III":
		{
			y = 3
		}
	case "IV":
		{
			y = 4
		}
	case "V":
		{
			y = 5
		}
	case "VI":
		{
			y = 6
		}
	case "VII":
		{
			y = 7
		}
	case "VIII":
		{
			y = 8
		}
	case "IX":
		{
			y = 9
		}
	case "X":
		{
			y = 10
		}
	default:
		{
			fmt.Printf("%v is not a Roman digit or too many roman digits in a string\n", r)
			os.Exit(2)
		}
	}
	return x, y
}

func intToRome(number int) string {
	var romanString string
	if number > 3999 {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, conversion := range conversions {
		for number >= conversion.value {
			romanString += conversion.digit
			number -= conversion.value
		}
	}

	return romanString
}
