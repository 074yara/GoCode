package main

import "fmt"

type User struct {
	Id   int64
	Name string
}

func main() {

	Users := []User{
		{
			Id:   1,
			Name: "Petya",
		},

		{
			Id:   2,
			Name: "Ivan",
		},

		{
			Id:   3,
			Name: "Ilya",
		},

		{
			Id:   4,
			Name: "Zheka",
		},

		{
			Id:   1,
			Name: "Petya",
		},
		{
			Id:   6,
			Name: "Dima",
		},
		{
			Id:   7,
			Name: "Sergey",
		},
		{
			Id:   8,
			Name: "Filip",
		},
		{
			Id:   9,
			Name: "Oleg",
		},
	}

	newMap := make(map[int64]User)

	for _, user := range Users {
		if _, ok := newMap[user.Id]; !ok {
			newMap[user.Id] = user
		}
	}
	myPrint(newMap)
	myPrint(findInMap(3, newMap))

}

func myPrint(a interface{}) {
	fmt.Printf("Type %T Value %v\n", a, a)
}

func findInMap(id int64, mapa map[int64]User) *User {
	if user, ok := mapa[id]; ok {
		return &user
	}
	return nil
}
