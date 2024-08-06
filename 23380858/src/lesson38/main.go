package main

// switch文
func main() {
	// n := 3
	// switch n {
	// case 1:
	// 	println("one")
	// case 2, 3, 4:
	// 	println("two")
	// default:
	// 	println("default")
	// }

	// switch n := 2; n {
	// case 1:
	// 	println("one")
	// case 2, 3, 4:
	// 	println("two")
	// default:
	// 	println("default")
	// }

	// n := 6
	// switch {
	// case n > 0 && n < 4:
	// 	println("0 < n < 4")
	// case n > 3 && n < 7:
	// 	println("3 < n < 7")
	// default:
	// 	println("default")
	// }

	switch n := 2; n {
	case 1:
		println("one")
	case 2, 3, 4:
		println("two")
	// case n > 3 && n < 6: // この条件は評価されないえらーになる
	// println("3 < n < 6")
	default:
		println("default")
	}

}
