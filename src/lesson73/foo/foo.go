package foo

const (
	Max = 100 // public
	min = 1   // private
)

func ReturnMin() int {
	return min
}

func returnMax() int {
	return Max
}
