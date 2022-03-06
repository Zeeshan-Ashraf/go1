package main

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/dao"
	"github.com/Zeeshan-Ashraf/go1/router"
	"gorm.io/gorm"
)

var db *gorm.DB = connections.DB
var err error = connections.Err

func main() {
	fmt.Print("Starting Application\n")
	/*Basic syntax of Go*/
	//basics.BasicDataType()
	//basics.UserInput()
	//basics.ArrayZ()
	//basics.SliceZ()
	//basics.Structs()
	//basics.MapsZ()

	/*calling get & post request from web*/
	//others.GetSimpleHtmlPageFromWeb()
	//controllers.SendPostReqToWebWithData()
	//controllers.ValidateJsonFromWeb(controllers.Json_data_from_web_correct)
	//controllers.ValidateJsonFromWeb(controllers.Json_data_from_web_wrong)
	//controllers.UnmarshalJson()

	/*build http router to serve GET POST PUL DELETE request*/
	//architecture of microservices: https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=go1#R5V1dd5s4EP01Pqf7kBwQn36043bb3e4mp2m3zSPGiq0tRgRwYufXrwQCgyRjEoNRti8tkgWIOzNXM9JIGRlX6%2B3vsRet%2FsILGIyAttiOjNkIAF13bfIfrdnlNY5t5BXLGC1Yo33FLXqGrFJjtRu0gEmtYYpxkKKoXunjMIR%2BWqvz4hg%2F1Zvd46D%2B1shbQqHi1vcCsfY7WqSrvNa1tH39R4iWq%2BLNusZ%2BWXtFY1aRrLwFfqpUGe9HxlWMcZpfrbdXMKDgFbjk93048GvZsRiGaZsbfOsZLX1wO%2F%2B6ch6APVksH58v2FMevWDDPniJ4zXrcLorUHhaoRTeRp5Py09E0iNjukrXASnp5NJLohz7e7SF5HVT9lQYp3B7sLt6CQLRHojXMI13pElxg8twY4qjg7GbVzzt5WCMWaNVRQZlpcdkvywfvoeHXDCE5GjFyfTLP5%2B%2BrfzpvzD45LsP48%2F%2FSNCKfi5PBuseBcEVDnBM6kIcwhI%2FASwJpIfxMzn8Couqwgdk8IG%2B4AMCfOTToX65xKpiaHMYGiKGZV0Vw7KycwwNKYZAXQwNTTkMTSmGhroY8noI3MExtEQMN6GvKwqgoXNKaIoAWhL8rL7gs6XwAVXhcxSDz5HCZyoKn64afK4A3wjYQUoBwOQrqxDaDxtc%2FHCRZO75hDRwo%2B3%2BN3K1zNrkDKBlN2ohpk%2FahOiBFrw1pL0KKZbUZ8pfR3qfvzF%2FwBuRH9DsYeU3Pll%2BuikTIJEM6E005GEsvHN7k5TJBw1ndXq%2F%2FPz7eYL0j%2B6dFm3mt6bz%2FGleMH0FObggESYr4jhd4SUOveD9vnYa4024oKjNiC1N920%2BYxwxfP%2BFabpjeHqbFNfRJwjGux%2Fs%2FqxwRwuXVlGcbas%2FznasdFAECd7EPmz4TDZ0pV68hGlDO%2BY3UAwaBRrDwEvRYz0K71w4YkSySWAsSIyE7hG93KyDiZ9SnZ1SnUQ%2BEYk3h8ENTlCKcEiazHGa4nWlwSRAS%2FpDijkLwZs0QCG8KicuGgXwAhuom4A1lpiAxALsvgxAF53C%2F6UFWC0toKAkRUxAdOQJ9ik1gnwcmMf7IYCv8SKUgUX0ivKn7a2pjmeNLsoiu5fO0KUxDgKJfb1ueum0Md3gzMSWhFSykULvbaTQjV%2FDUAoDOG4pQClL0d1B5LNF6Y%2B9REjprvLLXji0sKtK6kddjHdVeQ8nU7XG%2F6LbFfar0tRxBiy48ij7EbfiEfk0BiI9X%2BFFogIJGoalHAlaQxjZCYrf1vHVTbU0fzwomVWoTLt0rHZspo9qI5TT4xDVWqpqOXNFvyt8VhLPcTKrUt9RQlt4OGv1gfw7m5aXRrzIMU%2FpV09uPqnAc5aCzp64CDETJ1mKwNPfkVBxAWOD4HEEv3luo5%2FnZYXn%2F1xmlnudR5ysPp8MmulWNxg7fNwpWaOwJRC7vSHsvLGRxH6jbrG4luF7%2FgoK6Kcx8sJlBvMRHe7C5LklM8c6nzqCdOo%2BEsu9%2ByuabcI%2F0Q1%2BmEimmkT1DBcTmrJCDT7wkgT5dVjqynkQpKN60XLaurX6sDfcYJRNPm%2Fr6DIBjHkqzQ2C3bQHV3iOA448KDcY4UGZlMqvfr3gxBX3tZekklkMins2HcgNaGwC0CdygtWpw2JmcI0Wi5x8IOFkb549j0o4op%2BUfaQ1HVkzqcwbdU3Mv2FpUuwlo2omksyKLrRL2zWMGv7madpRNMH39wnsRWASQvp1JKZdupZekxdQXV7iDOT%2FnBn59cTuqNFpR40ES29XacbUtn2HnVqqIbnIn9itGYuL%2FHP5UoIaJlyocRcmrLnjGuBF5qG6RixxAoFEWkGAoqSFB9hHFAh4azljGoZcw4eZg%2Blzclj6newzq4GNfBDsfDJFTmdlgnCRJaDxXn9HvGmYFvcix24kdEtrvqEfpi1NqZIFjmivSBQXx7uLCPk%2FFSbewow6cXc1zXBrIjjR3S03QtTv6I%2BIxdwgX%2BZQDUnEhkDE1sBELAbnvixtYUjQdLeul64kr%2BO8oImBsQ8kS9hDgmaavKaBgUH75YIdU%2B8p2NE1Prm9o1Gb73H5on7DHTEpdy7RDXVG3Q7jHbvu9agf7RSqqPIoK%2BTCamdMW95FX79YyfW3D%2FH38ewafr99uPrjYpj8JrbkXFxXspta5M%2BMus%2BJkkJTbG2thEjSdqAlt3e%2B9NPUa0nuICdnRfL4TaduEbpkjdKVWEQXi0JSAAdZonyNRbxes0FLzR4sPaap1wfywpTUbuCopt2D5kuqzfdmS6sYLF2yqdey7CIlTcLUVDOJQbPuzmESWekGxojgRSORE%2B1EsqlC2s5Wyk7EIH82uVbURiygmo2IS0kCdOfKi%2BsHcttuAXlfyUnSY2AGSbrulGL22yTGZo3mDHCE6EhBzlccmRwlMCmyrkhg0nbmmfhK%2BvJBAvMOQZbEGPLTjk5FWT5hCMT9%2Fb2lqDV9v8wv0wJvl%2BVzv5tvEkJySULr8BL5v508IB3gvZec8sQfsGMWyUtnGHykUA6zXVu2LeIga1mg9N16JC5JhHKY4E41qZdO0dvGAYvraoq%2BCRIVHJMuTI%2FH0DhjirScnrVh3ZDX5MN0YmvFUup5jE1MU7F4a9I5ER9Yo%2BvKsIbdi3ZpVHfWEgfSGbcMjS9Nza37rppmvtZ7VcAJzfWwf%2F0STmjs2T8Sl1Zn7EQCuk5JD%2B4Q9K%2Ftyl0HRCysN5fHbR1xgvhs3%2B6IWHQo1%2Bxk16xEkIo3PpUwuqQ%2B5pw0i%2BJCeAO7lIA%2FcNCU5HmWjap4At5t7w7QYU8B4jbOgnbz2ucd2rqJzV6cockvkhfbFg9loQg3mB3nwjeiyM8o8jt432W8hiMYe5TWaKyHsjMHCuutb%2BpdeFjc2DubKhEZ8qlbugVEM%2B5v065cDOqsaLXzULKEVs5FMUzw%2Bjn8bjihbWzZjT%2FyUk4AfHq4aR%2FhhJNvsJzmG0xj3HhDT6wj2aP%2B9ZpcoDBJvTCbXsL3e3oZnjSAxuFkSmJaVxc5w21QopMoAwwT0559EbwbYpDsFDmceDTUbLlk2ueNnQOlDStSOarDylQ8O3gwmTaP7YPI5tRx%2BDSvS9zSsvaQGLN3tp5%2B2t%2Bt4HZwnXM5vcGwDs5%2BvBEYZX%2F947wwSra74PAeLSl6MN1EbwRIreXk0iuAJMX9n7PJndD9HwUy3v8H
	//router without gin
	//others.Server_without_gin() //if we run this server gin server won't run coz it'll never leave this line and keep running the http server to listen to web request

	/*initialize DB*/
	//Automatically initialize DB via init() as soon as we import connection pkg coz init() is defined in that pkg which connects to DB
	//so no need to call: db, err = connections.init() also init() is executed even before main() here, call is independent of main func
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		fmt.Printf("error in connecting DB, connections.init() error: %s\n", err.Error())
		fmt.Printf("DB conection Error: %v\n", err.Error())
	}
	dao.Init()

	//initialize router using gin
	router.InitializeRouter()

	//init() method //https://youtu.be/GszGvj6eBZY?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&t=230
	//gorm dao  //DB.Raw("SELECT id, name, price FROM courses WHERE name = ?", "ECE").Scan(&result) //DB.Where("name = ?", name).First(&course); == SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1; here &course defines which table
	//pagination //db.Limit(10).Offset(5).Find(&users) // SELECT * FROM users OFFSET 5 LIMIT 10; offset is how many rows to skip befr sending results
	//unit testing
	//middle ware
	//router group  //rt_group1 := rt.Group("/api"); rt_group1.GET("/getjsoncase4", controllers.funcz)
	//url params api/get/1   // id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	//url query  api/get?name=zee   //name := c.Query("name")
	//DTO vs DAO //DTO is struct which holds the data retrieved from DB, DAO is data access object same thing
	//cache layer
	//kafka
	//method (property of struct / custom type) vs func
	//goroutine & mutex & chanels

	//what happens if nested struct -> DAO -> DB
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
