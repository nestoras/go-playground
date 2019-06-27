package main

import "fmt"

type User interface {
	GetUsername()
	GetAge()
}

//struct private
type user struct {
	username string
	age      uint
}

func (u user) GetUsername() {
	fmt.Println(u.username)
}

func (u user) GetAge() {
	fmt.Println(u.age)
}

//expose only interface
func NewUser(username string, age uint) User {
	return &user{
		username: username,
		age:      age,
	}
}

func main() {
	user := NewUser("nestoras", 31)
	user.GetAge()
	user.GetUsername()
}
