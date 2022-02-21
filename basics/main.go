package main

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello World")
	basicDataType()
	userInput()
	arrayZ()
	sliceZ()

	//router using gin
	rt := gin.Default() //create gin router engine variable
	rt.GET("/zee", controllers.Hello)
	rt.Run(":8585")
}
