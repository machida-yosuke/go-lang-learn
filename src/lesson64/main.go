package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) setName() {
	u.Name = "A"
}

func main() {
	t := T{User{"user1", 20}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	t.User.setName()
	fmt.Println(t)
}
