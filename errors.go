package main

import "fmt"

type User struct {
	name string
	age  int
}

func getUser() (User, error) {
	//do things here
}

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
