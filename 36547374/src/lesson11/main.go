package main

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title      string
	Estimation int
}

func main() {
	task1 := Task{
		Title:      "Learn Go",
		Estimation: 25,
	}

	task1.Title = "Learn Go in 1 hour"
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	var task2 = task1
	task2.Title = "new"
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)
	fmt.Printf("%[1]T %+[1]v %v\n", task2, task2.Title)

	task1p := &Task{
		Title:      "Learn",
		Estimation: 2,
	}
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))

	task1p.Title = "Changed"
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))

	var task2p *Task = task1p
	task2p.Title = "Changed by task2p"
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	fmt.Printf("task2p: %T %+v %v\n", task2p, *task2p, unsafe.Sizeof(task2p))

	task1.extendEstimation()
	fmt.Printf("task1: %T %+v %v\n", task1, task1, unsafe.Sizeof(task1))

	(&task1).extendEstimatePointer()
	fmt.Printf("task1: %T %+v %v\n", task1, task1, unsafe.Sizeof(task1))
}

func (task Task) extendEstimation() {
	task.Estimation += 10
}

func (taskp *Task) extendEstimatePointer() {
	taskp.Estimation += 10
}
