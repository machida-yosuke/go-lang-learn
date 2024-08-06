package main

func PrintSliceInts(slice []int) {
	for _, v := range slice {
		println(v)
	}
}

func PrintSliceStrings(slice []string) {
	for _, v := range slice {
		println(v)
	}
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		println(v)
	}

}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5})
	// PrintSliceStrings([]string{"a", "b", "c", "d", "e"})
	PrintSlice[int]([]int{1, 2, 3, 4, 5})
	PrintSlice[string]([]string{"a", "b", "c", "d", "e"})
}
