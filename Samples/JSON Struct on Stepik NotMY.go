package main

import (
	"encoding/json"
	"os"
)

type Group1 struct {
	Students []struct {
		Rating []int
	}
}

func main() {
	var group Group1
	if err := json.NewDecoder(os.Stdin).Decode(&group); err != nil {
		panic(err)
	}

	var avg float64
	for _, st := range group.Students {
		avg += float64(len(st.Rating))
	}
	avg /= float64(len(group.Students))

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(struct{ Average float64 }{avg})
}
