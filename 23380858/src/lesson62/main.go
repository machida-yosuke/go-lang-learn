package main

import (
	"fmt"
)

// 構造体
// クラスみたいなやつ

type User struct {
	Name string
	Age  int
	// x, y int
}

func UpdateUser(u User) {
	u.Name = "a"
	u.Age = 1000
}

func UpdateUser2(u *User) {
	u.Name = "a"
	u.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 20
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 30
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 20}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5 := User{30, "user5"}
	// fmt.Println(user5) // エラー

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7) // ポインタを返す

	user8 := &User{}
	fmt.Println(user8) // ポインタを返す

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
