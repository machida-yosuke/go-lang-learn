package main

import (
	f "fmt"
	. "foo/foo" // .省略できる(非推奨)
)

func appName() string {
	const appName = "GoApp"
	var Version = "1.0"
	return appName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBB"
		f.Println(b)
	}
	f.Println(b)
	return b
}

func main() {
	f.Println(Max)
	// f.Println(foo.min) // Error: cannot refer to unexported name foo.min

	f.Println(ReturnMin())
	// f.Println(foo.returnMax()) // Error: cannot refer to unexported name foo.returnMax

	f.Println(appName())
	// f.Println(appName, Version) // Error: undefined: Version

	f.Println(Do("AAA"))
}
