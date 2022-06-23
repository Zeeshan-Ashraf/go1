package basics

import (
	"fmt"
	"sort"
)

func SortingArray() {
	//we cant sort Array[struct{}] but slice from go1.18
	s := []string{"c", "a", "b"}
	sort.Strings(s)
	fmt.Println("Strings:", s)

	a := []int{7, 2, 4}
	sort.Ints(a)
	fmt.Println("Ints:   ", a)

	isSorted := sort.IntsAreSorted(a)
	fmt.Println("Sorted: ", isSorted)
}

//sort slice of struct{} using custom sort func
//https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
