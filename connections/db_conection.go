package connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

const DNS = "root:password@tcp(127.0.0.1:3306)/go1?charset=utf8mb4&parseTime=True&loc=Local" //<user>:<password>@<protocol>(ip:port)/<databaseName also called schemaName>?charset=utf8mb4&parseTime=True&loc=Local
/*
init() function is just like the main function, does not take any argument nor return anything.
This function is present in every package and this function is called when the package is initialized.
This function is declared implicitly, so you cannot reference it from anywhere and you are allowed to
create multiple init() function in the same program and they execute in the order they are created.
You are allowed to create init() function anywhere in the program and they are called in lexical file name order (Alphabetical Order).
And allowed to put statements if the init() function,
***
but always remember to init() function is executed before the main() function call, so it does not depend to main() function.
***
The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.
*/

func init() {
	DB, Err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	fmt.Printf("DB conected DNS: %s\n", DNS)
	if Err != nil {
		fmt.Printf("error in connecting DB: %s\n", Err.Error())
	}
}
