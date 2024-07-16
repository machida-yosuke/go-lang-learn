package main

// 関数を引数にする関数

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		println("I'm a function")
	})
}
