package main

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/dao"
	"github.com/Zeeshan-Ashraf/go1/router"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	fmt.Print("Starting Application\n")
	/*Basic syntax of Go*/
	//basicDataType()
	//userInput()
	//arrayZ()
	//sliceZ()

	/*calling get & post request from web*/
	//others.GetSimpleHtmlPageFromWeb()
	//controllers.SendPostReqToWebWithData()
	//controllers.ValidateJsonFromWeb(controllers.Json_data_from_web_correct)
	//controllers.ValidateJsonFromWeb(controllers.Json_data_from_web_wrong)
	//controllers.UnmarshalJson()

	/*build http router to serve GET POST PUL DELETE request*/
	//router without gin
	//others.Server_without_gin() //if we run this server gin server won't run coz it'll never leave this line and keep running the http server to listen to web request

	//initialize DB
	db, err = connections.ConnectToDB()
	if err != nil {
		fmt.Printf("error in main.connections.ConnectToDB() error: %s\n", err.Error())
		panic(err)
	}
	dao.Init()

	//initialize router using gin
	router.InitializeRouter()

	//init() method //https://youtu.be/GszGvj6eBZY?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&t=230
	//gorm dao
	//unit testing
	//middle ware
	//router group
	//url query
	//DTO vs DAO
	//cache layer
	//kafka
	//method (property of struct / custom type) vs func
	//goroutine & mutex & chanels
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
