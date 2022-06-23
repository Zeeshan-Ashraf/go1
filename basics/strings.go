package basics

import (
	"fmt"
	"strconv"
	"strings"
)

func stringOperations() {
	s := "ABC"
	fmt.Printf("%s %c\n", s, s[1]) //OUTPUT: ABC B

	for idx, char := range s { // SYNTAX:    for index,elem := collection{}
		fmt.Printf("idx=%d character=%c\n", idx, char)
	} /*OUTPUT
	idx=0 character=A
	idx=1 character=B
	idx=2 character=C
	*/

	fmt.Printf("lenght of string = %d", len(s)) //OUTPUT: 3 for ABC

	//STRING TO INT & INT to STRING
	s = "1"
	numbr, err := strconv.Atoi(s)
	fmt.Printf("%d %s", numbr, err.Error()) //OUTPUT: ABC B

	t := strconv.Itoa(123)
	fmt.Printf("%s", t)

	//STRING split
	s = "AB#CD##EF"
	r := strings.Split(s, "#")      //it'll return array []string with splitted words using # into 4 words viz: AB CD emptyString EF
	fmt.Printf("%+v %d", r, len(r)) //OUTPUT: [AB CD  EF]

	//STRING replace
	fmt.Println(strings.Replace("zee zee zee", "z", "Z", 2))  //OUTPUT: Zee Zee zee  i.e replace z with Z 2 times
	fmt.Println(strings.Replace("zee zee zee", "z", "Z", -1)) //n<0 means replace all occurances
}
