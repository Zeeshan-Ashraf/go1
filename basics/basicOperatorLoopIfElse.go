package basics

import "fmt"

//same as C except no bracket
func Loop() {
	for i := 0; i < 5; i++ {
		fmt.Printf("value i = %d\n", i)
	}
	i := 0
	for ; i < 2; i++ {
		fmt.Printf("v i = %d\n", i)
	}
	fmt.Printf("end i = %d\n", i)

	//STRING loop in strings.go

	//break and continue == same as c
	for i := 0; i < 6; i++ {
		if i < 2 {
			continue
		}
		if i > 4 {
			break
		}
		fmt.Printf("value i = %d\n", i)
	}

}

/*
OUTPUT:
value i = 0
value i = 1
value i = 2
value i = 3
value i = 4
v i = 0
v i = 1
end i = 2
*/

//same as C except must have parenthesis if{} and new line matters, so else if & else should be as shown below not in next line
func IfElse() {
	i := 11

	if (i != 0 && i > 1) || i <= 10 {
		fmt.Printf("%d", i)
	} else if i > 10 {
		fmt.Printf("%d", i)
	} else {
		fmt.Printf("%d", i)
	}

	//initialization + condition both in if (just like for loop)
	if i := 10; i < 20 {
		fmt.Printf("in if %d", i) //OUTPUT: in if 10
	}

	mz := map[int]string{}
	if v, ok := mz[100]; ok {
		fmt.Printf("%s", v)
	}
}

func MathOperators() {
	fmt.Printf("%d", 7/3) //output: 2
	fmt.Printf("%d", 7%3) //output: 1
	fmt.Printf("%d", 7+3) //output: 10
	fmt.Printf("%d", 7-3) //output: 2
}

func MemoryAdd() {
	i := 0
	j := i
	fmt.Printf("add i = %d j=%d j=%d\n", &i, &j, *(&j))
}

/*
*OUTPUT:
add i = 824634933264 j=824634933272 j=0
*/
