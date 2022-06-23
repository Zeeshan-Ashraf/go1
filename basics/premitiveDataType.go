package basics

import (
	"bufio"
	"fmt"
	"os"
)

//global variable
var global_var = 44
var global_var2 int

//global_var2 := 66 //cant this way in global var

//Note: go is like C i.e it passes copy of variable in fn calls hence one need to pass pointer in case required

func BasicDataType() {
	fmt.Print("\ninside basicDataType\n")
	var username string = "Zeeshan"
	var username2 = "Zeeshan"
	//var adhaar string //declare a var with nil value
	roll := 101
	balance := 55.506789
	bool_val := true
	fmt.Printf("%T %T %T %T %T %T\n", username, username2, roll, balance, bool_val, global_var)
	fmt.Printf("%s %s %d %.3f %v %d\n", username, username2, roll, balance, bool_val, global_var)
	fmt.Printf("%v %v %v %v %v %+v\n", username, username2, roll, balance, bool_val, global_var)
}

func UserInput() string {
	fmt.Print("enter any string and end with delimiter '\\n':")
	reader := bufio.NewReader(os.Stdin)      //create a reader to listen to OS standard input
	inputZee, err := reader.ReadString('\n') //reads until new line is provided
	if err != nil {
		return err.Error()
	}
	fmt.Printf("user input value: %s", inputZee)
	return "success"
}
