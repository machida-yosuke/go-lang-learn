package main

// 関数を返す関数

func returnFunc() func() {
	return func() {
		println("I'm a function")
	}
}

func main() {
	f := returnFunc()
	f()
}
