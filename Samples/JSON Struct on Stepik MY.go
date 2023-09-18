package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Students
}

type Students struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

func main() {

	var allGroup Group

	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("reader error:", err)
	}

	err = json.Unmarshal(inputData, &allGroup)
	if err != nil {
		fmt.Println("Unmarshal ", err)

	}

	ratingCount := 0

	for _, v := range allGroup.Students {
		ratingCount += len(v.Rating)
	}

	newMap := struct {
		Average float64
	}{
		Average: Round(float64(ratingCount)/float64(len(allGroup.Students)), 1),
	}

	newJS, err := json.MarshalIndent(newMap, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newJS))

}
