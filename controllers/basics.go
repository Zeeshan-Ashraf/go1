package controllers

import (
	"github.com/gin-gonic/gin"
)

/*
simplest get request, returns string (not json) = "Hello There" to users
*/
func Hello(c *gin.Context) { //first letter upper case means method is public, c contains request details like url host body cookies headers etc, c is Context which can also validates and serializes JSON
	c.String(200, "Hello There")
}
