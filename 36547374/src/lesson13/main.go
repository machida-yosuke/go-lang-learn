package main

import "fmt"

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (v vehicle) String() string {
	return fmt.Sprintf("speed: %d, enginePower: %d", v.speed, v.enginePower)
}

func (v *bycycle) speedUp() int {
	v.speed += 3 * v.humanPower
	return v.speed
}

func (v *bycycle) speedDown() int {
	v.speed -= 1 * v.humanPower
	return v.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %d\n", c.speedUp())
	fmt.Printf("current speed: %d\n", c.speedDown())
}

func main() {
	v := &vehicle{0, 5}
	speedUpAndDown(v)

	b := &bycycle{0, 1}
	speedUpAndDown(b)
	fmt.Println(v)

	var i1 interface{}
	var i2 any
	checkType(i1)
	i2 = 1
	i2 = "hello"
	checkType(i2)
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
