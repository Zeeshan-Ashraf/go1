package basics

import "fmt"

//func definintion position deosn't matter where it is written in page

func SimpleFnWithoutInput() {
	fmt.Printf("hello")
}

func SimpleFnWithInpit(i int, s, t string) {
	fmt.Printf("%d %s %s", i, s, t)
}

func SimpleFnWithInputOutput(i int) string {
	fmt.Printf("%d", i)
	return "bye"
}

func SimpleFnWithMultipleOutput(i int) (int, string) {
	fmt.Printf("%d", i)
	return 100, "bye"
}

func SimpleFnWithUnknowInputCount(slicez ...int) string { //call this fn as SimpleFnWithUnknowInputCount(1,2,5,6)
	fmt.Printf("%v", slicez)
	return "bye"
}
