package basics

import (
	"fmt"
	"sort"
)

func ArrayZ() {
	var arrZee [3]int
	arrZee[0] = 101
	arrZee[1] = 102
	//array by default initializes to 0 instead of garbage value
	//%v prints entire array (space delimited)
	fmt.Printf("\n%d %d %d %v %+v\n", arrZee[0], arrZee[2], len(arrZee), arrZee, arrZee)
}

func SliceZ() { //slice is like LL
	var slice1 []int //by default slice1 is nil
	var slice2 = []int{1001, 1002}
	fmt.Printf("\nslice values: slice1=%v slice2=%v", slice1, slice2)
	slice1 = append(slice2) //slice1
	fmt.Printf("\nafter append slice1=%v", slice1)
	slice1 = append(slice2)
	fmt.Printf("\nafter 2nd append slice1=%v", slice1)
	slice1 = append(slice1, slice1...) //... is required for appending slice so we cant do this-> slice1 = append(slice1, slice1) as multiple values
	fmt.Printf("\nafter 3rd append slice1=%v", slice1)
	slice1 = append(slice1, 2001, 2002, 2003)
	fmt.Printf("\nafter 4th append slice1=%v", slice1)
	slice1 = append(slice1[1:])
	fmt.Printf("\nafter 5th append slice1=%v", slice1)
	slice1 = append(slice1[:2])
	fmt.Printf("\nafter 6th append slice1=%v\n", slice1)

	//remove elem from slice
	//remove index = 2
	slice1 = append(slice1[:2], slice1[3:]...)
	sort.Ints(slice1)
	fmt.Printf("after sorting: %v", slice1)

	//sort slice based on custom fn
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
}

func Structs() {
	//creating struct without any value
	type st1 struct{}
	var st11 st1 = st1{}
	fmt.Printf("value st1{} %+v\n", st11)

	type st2 struct {
		Id int
	}
	fmt.Printf("value st2{Id int} %+v\n", st2{101})

	/* not allowed this way of creation */
	//var st33 = type st3 struct {
	//	Id int
	//}{201}
	//fmt.Printf("value st1{} %+v\n", st11)
	type st3 struct {
		Id int
	}
	var st33 = st3{201}
	fmt.Printf("value st1{} %+v\n", st33)

	//note if you want to save nil in struct not possible, so use struct pointer & pointer can be nil

}

func MapsZ() {
	var map_int_string map[int]string
	if map_int_string == nil {
		fmt.Printf("\nMap is empty i.e map_int_string = %+v", map_int_string)
	}
	map_int_string = nil //we can assign nil to map unlike string struct int etc

	if map_int_string != nil {
		fmt.Printf("\nMap is empty i.e map_int_string = %+v", map_int_string)
	}
	map_int_string2 := map[int]string{1: "sun", 2: "mon"}
	fmt.Printf("\nmap_int_string2: %+v", map_int_string2)
	fmt.Printf("\nmap_int_string2[1]=%+v", map_int_string2[1])
	fmt.Printf("\nmap_int_string2[100]=%+v", map_int_string2[100])

	map_int_string3 := map[int]string{} //create empty map (its not nil map)
	if map_int_string3 == nil {
		fmt.Printf("\nMap is null i.e map_int_string3 = %+v", map_int_string3)
	} else {
		fmt.Printf("\nMap is not null but just empty i.e map_int_string3 = %+v", map_int_string3)
	}

	//var map_int_string map[int]string{1:"sun"} //not allowed creating this way but
	var map_int_string4 map[int]string = map[int]string{1: "sun"}
	fmt.Printf("\nmap_int_string4: %+v", map_int_string4)

	//if key exists in map or not
	if v, ok := map_int_string4[100]; ok {
		fmt.Printf("%s", v)
	}

	var map_int_string5 map[int]interface{} //map[int]anything so use "interface{}" (not "interface") for any data type
	map_int_string5 = map[int]interface{}{
		1: 101,
		2: "mon",
		3: struct {
			id   int
			name string
		}{
			101, "zee",
		}}
	fmt.Printf("map_int_string5: %+v", map_int_string5)
}
