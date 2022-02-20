package main

import (
	"fmt"
	"sort"
)

func arrayZ() {
	var arrZee [3]int
	arrZee[0] = 101
	arrZee[1] = 102
	//array by default initializes to 0 instead of garbage value
	//%v prints entire array (space delimited)
	fmt.Printf("\n%d %d %d %v %+v\n", arrZee[0], arrZee[2], len(arrZee), arrZee, arrZee)
}

func sliceZ() { //slice is like LL
	var slice1 []int //by default slice1 is nil
	var slice2 = []int{1001, 1002}
	fmt.Printf("\nslice values: slice1=%v slice2=%v", slice1, slice2)
	slice1 = append(slice2) //slice1
	fmt.Printf("\nafter append slice1=%v", slice1)
	slice1 = append(slice2)
	fmt.Printf("\nafter 2nd append slice1=%v", slice1)
	slice1 = append(slice1, slice1...) //cant do slice1 = append(slice1, slice1) as multiple values
	fmt.Printf("\nafter 3rd append slice1=%v", slice1)
	slice1 = append(slice1, 2001, 2002)
	fmt.Printf("\nafter 4th append slice1=%v", slice1)
	slice1 = append(slice1[1:])
	fmt.Printf("\nafter 5th append slice1=%v", slice1)
	slice1 = append(slice1[:2])
	fmt.Printf("\nafter 6th append slice1=%v\n", slice1)
	sort.Ints(slice1)
	fmt.Printf("after sorting: %v", slice1)
}
