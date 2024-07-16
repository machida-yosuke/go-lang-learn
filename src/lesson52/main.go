package main

// for

func main() {
	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}

	for k, v := range m {
		println(k, v)
	}

}
