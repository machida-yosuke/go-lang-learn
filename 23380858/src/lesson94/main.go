package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json
type A struct{}

//	type User struct {
//		ID      int       `json:"id,string"` // jsonにしたときidになり、文字列になる
//		Name    string    `json:"name"`
//		Email   string    `json:"email"`
//		Created time.Time `json:"-"` // jsonにしたとき無視される
//		A       A         `json:"A"`
//	}

// type User struct {
// 	ID      int       `json:"id,omitempty"` // 初期値の場合は無視される
// 	Name    string    `json:"name"`
// 	Email   string    `json:"email"`
// 	Created time.Time `json:"created"`
// 	A       *A        `json:"A, omitempty"`
// }

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + u.Name,
	})
	return v, err
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "aaa"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// ========

	fmt.Printf("%T\n", bs)
	u2 := new(User)
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
