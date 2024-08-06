package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now()
	// fmt.Println(t)

	// t2 := time.Date(2020, 6, 21, 10, 10, 10, 10, time.Local)
	// fmt.Println(t2)
	// fmt.Println(t2.Year())
	// fmt.Println(t2.Month())
	// fmt.Println(t2.Day())
	// fmt.Println(t2.Weekday())
	// fmt.Println(t2.Hour())
	// fmt.Println(t2.Minute())
	// fmt.Println(t2.Second())
	// fmt.Println(t2.Nanosecond())
	// fmt.Println(t2.Location())
	// fmt.Println(t2.Zone())

	// fmt.Println(time.Hour)
	// fmt.Println(time.Minute)
	// fmt.Println(time.Second)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Microsecond)
	// fmt.Println(time.Nanosecond)

	// d, _ := time.ParseDuration("1h30m")
	// fmt.Println(d)

	// t3 := time.Now()
	// t3 = t3.Add(2*time.Hour + 15*time.Second)
	// fmt.Println(t3)

	// t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	// t6 := time.Now()

	// d2 := t5.Sub(t6)
	// fmt.Println(d2)

	// fmt.Println(t6.Before(t5))
	// fmt.Println(t6.After(t5))
	// fmt.Println(t6.Equal(t5))

	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止")
}
