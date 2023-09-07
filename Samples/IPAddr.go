package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"1st": {127, 0, 0, 1},
		"2nd": {8, 8, 8, 8},
	}
	a := hosts["1st"]
	b := hosts["2nd"]
	fmt.Println(a)
	fmt.Println(b)
}
