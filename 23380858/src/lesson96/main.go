package main

import (
	"fmt"
	"sort"
)

// sort

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	i := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	s := []string{"Alice", "Bob", "Dave", "Eve", "Charlie"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)

	el := []Entry{
		{"Alice", 22},
		{"Bob", 25},
		{"Dave", 22},
		{"Eve", 28},
		{"Charlie", 23},
	}

	fmt.Println(el)

	sort.Slice(el, func(i, j int) bool { // 数値の場合はSliceを使う
		return el[i].Value < el[j].Value
	})
	fmt.Println(el)

	sort.SliceStable(el, func(i, j int) bool { // 文字列の場合はStableを使う
		return el[i].Name < el[j].Name
	})
	fmt.Println(el)

	m := map[string]int{"S": 1, "A": 2, "B": 3, "D": 3, "C": 5}
	fmt.Println(m)

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}
	sort.Sort(lt)
	fmt.Println(lt)

	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
