package main

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/controllers"
	"github.com/Zeeshan-Ashraf/go1/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello World\n")
	basicDataType()
	userInput()
	arrayZ()
	sliceZ()

	utils.GetSimpleMsgFromWeb()

	//router using gin
	rt := gin.Default()               //create gin router engine variable
	rt.GET("/zee", controllers.Hello) //note this[ hello() ] requires return type i.e this rt.GET("/zee", controllers.Hello()) & it not recommended in router, so in order to call function you need to call them from handler (no params except default *gin.Context)
	rt.GET("/weather/:loc", controllers.Get_weather_by_location)
	rt.Run(":8585")
}

/*
o/p:
SW-LP08933:go1 md.ashraf$ pwd
/Users/md.ashraf/z/go1
SW-LP08933:go1 md.ashraf$ cd basics/
SW-LP08933:basics md.ashraf$ go run *.go
Hello World
inside basicDataType
string string int float64 bool int
Zeeshan Zeeshan 101 55.507 true 44
Zeeshan Zeeshan 101 55.506789 true 44
enter any string and end with delimiter '\n':abcd
user input value: abcd

101 0 3 [101 102 0] [101 102 0]

slice values: slice1=[] slice2=[1001 1002]
after append slice1=[1001 1002]
after 2nd append slice1=[1001 1002]
after 3rd append slice1=[1001 1002 1001 1002]
after 4th append slice1=[1001 1002 1001 1002 2001 2002]
after 5th append slice1=[1002 1001 1002 2001 2002]
after 6th append slice1=[1002 1001]
after sorting: [1001 1002][GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /zee                      --> github.com/Zeeshan-Ashraf/go1/controllers.Hello (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8585

*/
