package controllers

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) { //first letter upper case means method is public
	c.String(200, "Hello There")
}
