package router

import (
	"github.com/Zeeshan-Ashraf/go1/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	//router using gin
	rt := gin.Default() //create gin router engine variable with middleware & logger
	//rt := gin.New() //create empty gin router engine variable without middleware & logger i.e you wont see any get logs in term whenever url is hit
	rt.GET("/zee", controllers.Hello)                         //note this[ hello() ] requires return type i.e this rt.GET("/zee", controllers.Hello()) & it not recommended in router, so in order to call function you need to call them from handler (no params except default *gin.Context)
	rt.GET("/weather/:loc", controllers.GetWeatherByLocation) //localhost:8585/weather/Kolkata, returns HTML to user received from web
	rt.GET("/getjsoncase1", controllers.SendJsonUsingGinH)    //returns JSON data by converting a map to JSON
	rt.GET("/getjsoncase2", controllers.SendSimpleJSON)       //returns JSON data by converting a struct to JSON
	rt.GET("/getjsoncase3", controllers.SendJsonUsingMarshal) //returns JSON data by converting a struct to JSON using json.Marshal lib
	rt_group1 := rt.Group("/api")
	rt_group1.GET("/getjsoncase4", controllers.SendJsonWithCustomkeyName) //hit localhost:8585/api/getjsoncase4 returns JSON data by converting a struct to JSON having custom json fields, hide password etc
	rt.GET("/getwithtimeout/:tim", controllers.GetReqWithTimeoutFromWeb)  //req 50MB JSON data to Web with timeout
	rt_group2 := rt_group1.Group("/v1")
	{
		rt_group2.GET("/getenvvar", controllers.GetPathEnvVar) //read OS env var, user must hit localhost:8585/api/v1/getenvvar
	}
	//send local html files to user
	rt.LoadHTMLFiles("templates/index.html")                              // can also use rt.LoadHTMLGLob for dir path if we have multiple html files inside a dir
	rt.GET("/index", controllers.SendLocalHtmlFile)                       // renders static HTML page from local
	rt.POST("/returnadharjson", controllers.UnmarshalJsonUsingShouldBind) //send json data in body e.g curl --request POST '127.0.0.1:8585/returnjson' -d '{"Name":"Zeeshan","Id":[101,102],"Addresses":[{"City":"Kolkata","Pincode":700046},{"City":"Bangalore","Pincode":800002}]}'
	rt.POST("/returnanyjson", controllers.UnmarshalAnyJsonToMap)          //send json data in body e.g curl --request POST '127.0.0.1:8585/returnanyjson' -d '{"Name":"Zeeshan","Id":[101,102],"Addresses":[{"City":"Kolkata","Pincode":700046},{"City":"Bangalore","Pincode":800002}]}'

	//handle DB select & insert operarion
	rt.POST("/course", controllers.InsertCourse)

	//run gin server to server http requests
	rt.Run(":8585") //rt.Run() by default runs on port 8080 p.s port 8080 & 80 are not same
}
