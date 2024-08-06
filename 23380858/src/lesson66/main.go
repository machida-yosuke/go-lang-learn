package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{"user1", 10}
	user2 := User{"user2", 20}
	users := Users{&user1, &user2}
	fmt.Println(users)

	user3 := User{"user3", 30}
	users = append(users, &user3)
	fmt.Println(users)

	for _, user := range users {
		fmt.Println(*user)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user2)
	fmt.Println(users2)
}
