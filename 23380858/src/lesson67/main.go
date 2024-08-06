package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{"user1", 10}: "user1",
		{"user2", 20}: "user2",
	}

	fmt.Println(m2)

	m3 := make(map[int]User)
	m3[1] = User{"user1", 10}
	m3[2] = User{"user2", 20}
	fmt.Println(m3)

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(v.Age)
	}
}
